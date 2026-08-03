package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMatchTopic(t *testing.T) {
	tests := []struct {
		pattern  string
		topic    string
		expected bool
	}{
		{"orders.created", "orders.created", true},
		{"orders.*", "orders.created", true},
		{"orders.*", "orders.updated", true},
		{"orders.*", "orders.items.created", false},
		{"orders.>", "orders.items.created", true},
		{">", "anything.at.all", true},
		{"users.created", "orders.created", false},
	}

	for _, tt := range tests {
		got := MatchTopic(tt.pattern, tt.topic)
		if got != tt.expected {
			t.Errorf("MatchTopic(%q, %q) = %v; expected %v", tt.pattern, tt.topic, got, tt.expected)
		}
	}
}

func TestEventBusPublishSubscribe(t *testing.T) {
	bus := NewEventBus(10)
	defer bus.Close()

	sub, err := bus.Subscribe("sub1", "orders.*")
	if err != nil {
		t.Fatalf("failed to subscribe: %v", err)
	}

	evt, err := bus.Publish("orders.created", "order_123")
	if err != nil || evt == nil {
		t.Fatalf("failed to publish: %v", err)
	}

	select {
	case received := <-sub.Ch:
		if received.Payload != "order_123" || received.Topic != "orders.created" {
			t.Errorf("received incorrect event: %+v", received)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for event delivery")
	}

	stats := bus.GetStats()
	if stats.TotalPublished != 1 || stats.TotalDelivered != 1 || stats.ActiveSubscribers != 1 {
		t.Errorf("unexpected stats: %+v", stats)
	}
}

func TestEventBusHTTPRouter(t *testing.T) {
	bus := NewEventBus(10)
	defer bus.Close()

	router := bus.Router()

	// Register subscriber via HTTP is tested separately, let's subscribe directly
	bus.Subscribe("http-sub", "logs.*")

	// POST /publish
	payload := `{"topic":"logs.error","payload":"DB connection timeout"}`
	req := httptest.NewRequest("POST", "/publish", bytes.NewBufferString(payload))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Fatalf("expected 202 Accepted, got %d", rr.Code)
	}

	// GET /stats
	req = httptest.NewRequest("GET", "/stats", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}
	var stats BusStats
	if err := json.Unmarshal(rr.Body.Bytes(), &stats); err != nil || stats.TotalPublished != 1 {
		t.Errorf("unexpected stats response: %v, stats=%+v", err, stats)
	}

	// GET /topics
	req = httptest.NewRequest("GET", "/topics", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}
	var topics []string
	if err := json.Unmarshal(rr.Body.Bytes(), &topics); err != nil || len(topics) != 1 || topics[0] != "logs.*" {
		t.Errorf("unexpected topics response: %v, topics=%v", err, topics)
	}
}
