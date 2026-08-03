package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Event represents a message routed through the event bus.
type Event struct {
	ID        string    `json:"id"`
	Topic     string    `json:"topic"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
	RetryCount int      `json:"retry_count"`
}

// Subscriber represents an active subscriber with an isolated buffer channel.
type Subscriber struct {
	ID        string
	Topic     string
	Ch        chan Event
	AckCh     chan string // Event ID ACKs
	cancel    context.CancelFunc
}

// BusStats tracks throughput and active subscriber metrics.
type BusStats struct {
	TotalPublished int64            `json:"total_published"`
	TotalDelivered int64            `json:"total_delivered"`
	ActiveSubscribers int           `json:"active_subscribers"`
	TopicCounts    map[string]int   `json:"topic_counts"`
}

// EventBus coordinates concurrent publish-subscribe with wildcard topic routing and backpressure.
type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string]map[string]*Subscriber // topic -> subID -> Subscriber
	bufferSize  int
	published   int64
	delivered   int64
	ctx         context.Context
	cancel      context.CancelFunc
}

// NewEventBus initializes a new EventBus instance.
func NewEventBus(bufferSize int) *EventBus {
	ctx, cancel := context.WithCancel(context.Background())
	if bufferSize <= 0 {
		bufferSize = 100
	}
	return &EventBus{
		subscribers: make(map[string]map[string]*Subscriber),
		bufferSize:  bufferSize,
		ctx:         ctx,
		cancel:      cancel,
	}
}

// Close gracefully stops the event bus.
func (eb *EventBus) Close() {
	eb.cancel()
}

// MatchTopic checks if an event topic matches a subscription topic pattern (supporting wildcard * and >).
func MatchTopic(pattern, topic string) bool {
	if pattern == topic || pattern == ">" {
		return true
	}

	pParts := strings.Split(pattern, ".")
	tParts := strings.Split(topic, ".")

	for i := 0; i < len(pParts); i++ {
		if pParts[i] == ">" {
			return true
		}
		if i >= len(tParts) {
			return false
		}
		if pParts[i] != "*" && pParts[i] != tParts[i] {
			return false
		}
	}
	return len(pParts) == len(tParts)
}

// Subscribe registers a new subscriber for a given topic pattern.
func (eb *EventBus) Subscribe(subID, topicPattern string) (*Subscriber, error) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if subID == "" || topicPattern == "" {
		return nil, errors.New("subID and topicPattern are required")
	}

	ctx, cancel := context.WithCancel(eb.ctx)
	sub := &Subscriber{
		ID:     subID,
		Topic:  topicPattern,
		Ch:     make(chan Event, eb.bufferSize),
		AckCh:  make(chan string, eb.bufferSize),
		cancel: cancel,
	}

	if _, exists := eb.subscribers[topicPattern]; !exists {
		eb.subscribers[topicPattern] = make(map[string]*Subscriber)
	}
	eb.subscribers[topicPattern][subID] = sub

	go func() {
		<-ctx.Done()
		eb.Unsubscribe(subID, topicPattern)
	}()

	return sub, nil
}

// Unsubscribe removes a subscriber.
func (eb *EventBus) Unsubscribe(subID, topicPattern string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if subs, exists := eb.subscribers[topicPattern]; exists {
		if sub, found := subs[subID]; found {
			close(sub.Ch)
			delete(subs, subID)
		}
		if len(subs) == 0 {
			delete(eb.subscribers, topicPattern)
		}
	}
}

// Publish delivers an event to all matching subscribers. Non-blocking when backpressure allows.
func (eb *EventBus) Publish(topic, payload string) (*Event, error) {
	if topic == "" {
		return nil, errors.New("topic cannot be empty")
	}

	event := Event{
		ID:        fmt.Sprintf("evt-%d", time.Now().UnixNano()),
		Topic:     topic,
		Payload:   payload,
		Timestamp: time.Now(),
	}

	eb.mu.RLock()
	defer eb.mu.RUnlock()

	atomic.AddInt64(&eb.published, 1)

	for pattern, subs := range eb.subscribers {
		if MatchTopic(pattern, topic) {
			for _, sub := range subs {
				select {
				case sub.Ch <- event:
					atomic.AddInt64(&eb.delivered, 1)
				default:
					// Backpressure triggered (subscriber buffer full)
					log.Printf("Backpressure: Subscriber %s buffer full, dropping event %s", sub.ID, event.ID)
				}
			}
		}
	}

	return &event, nil
}

// GetStats gathers monitoring metrics for the event bus.
func (eb *EventBus) GetStats() BusStats {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	totalSubs := 0
	topicCounts := make(map[string]int)

	for pattern, subs := range eb.subscribers {
		count := len(subs)
		totalSubs += count
		topicCounts[pattern] = count
	}

	return BusStats{
		TotalPublished:   atomic.LoadInt64(&eb.published),
		TotalDelivered:   atomic.LoadInt64(&eb.delivered),
		ActiveSubscribers: totalSubs,
		TopicCounts:      topicCounts,
	}
}

// ListTopics returns all active topic subscription patterns.
func (eb *EventBus) ListTopics() []string {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	topics := make([]string, 0, len(eb.subscribers))
	for pattern := range eb.subscribers {
		topics = append(topics, pattern)
	}
	return topics
}

// Router sets up HTTP REST and SSE endpoints for the event bus.
func (eb *EventBus) Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Topic   string `json:"topic"`
			Payload string `json:"payload"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		evt, err := eb.Publish(req.Topic, req.Payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(evt)
	})

	mux.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		topic := r.URL.Query().Get("topic")
		if topic == "" {
			http.Error(w, "query parameter 'topic' is required", http.StatusBadRequest)
			return
		}

		subID := fmt.Sprintf("sub-%d", time.Now().UnixNano())
		sub, err := eb.Subscribe(subID, topic)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		for {
			select {
			case <-r.Context().Done():
				sub.cancel()
				return
			case evt, open := <-sub.Ch:
				if !open {
					return
				}
				data, _ := json.Marshal(evt)
				fmt.Fprintf(w, "data: %s\n\n", string(data))
				flusher.Flush()
			}
		}
	})

	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eb.GetStats())
	})

	mux.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eb.ListTopics())
	})

	return mux
}

func main() {
	bus := NewEventBus(50)
	defer bus.Close()

	port := 9092
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: bus.Router(),
	}

	log.Printf("Starting Event Stream Processor on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
