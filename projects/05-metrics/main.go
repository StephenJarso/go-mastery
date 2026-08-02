package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// MetricType defines supported observability metric types.
type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeGauge     MetricType = "gauge"
	MetricTypeHistogram MetricType = "histogram"
)

// Sample represents a single timestamped metric observation.
type Sample struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

// MetricStats contains aggregated metrics calculations.
type MetricStats struct {
	Name        string     `json:"name"`
	Type        MetricType `json:"type"`
	Count       int        `json:"count"`
	Sum         float64    `json:"sum"`
	Min         float64    `json:"min"`
	Max         float64    `json:"max"`
	Avg         float64    `json:"avg"`
	P95         float64    `json:"p95"`
	P99         float64    `json:"p99"`
	LastUpdated time.Time  `json:"last_updated"`
}

// MetricSeries holds time-windowed values for a specific metric key.
type MetricSeries struct {
	Name      string     `json:"name"`
	Type      MetricType `json:"type"`
	Samples   []Sample   `json:"samples"`
	Buckets   map[string]int `json:"buckets,omitempty"` // For histogram bucketing (e.g. "<=10", "<=50", "<=100", "+Inf")
}

// MetricsCollector manages thread-safe ingestion, retention, and aggregation of metrics.
type MetricsCollector struct {
	mu            sync.RWMutex
	series        map[string]*MetricSeries
	retentionWindow time.Duration
	ctx           context.Context
	cancel        context.CancelFunc
}

// NewMetricsCollector initializes a new collector with a retention time window.
func NewMetricsCollector(retentionWindow time.Duration) *MetricsCollector {
	ctx, cancel := context.WithCancel(context.Background())
	mc := &MetricsCollector{
		series:          make(map[string]*MetricSeries),
		retentionWindow: retentionWindow,
		ctx:             ctx,
		cancel:          cancel,
	}

	go mc.startPruningLoop()
	return mc
}

// Close stops background maintenance tasks.
func (mc *MetricsCollector) Close() {
	mc.cancel()
}

// Observe records a new measurement for a given metric.
func (mc *MetricsCollector) Observe(name string, mType MetricType, val float64) {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	s, exists := mc.series[name]
	if !exists {
		s = &MetricSeries{
			Name:    name,
			Type:    mType,
			Samples: make([]Sample, 0),
			Buckets: make(map[string]int),
		}
		mc.series[name] = s
	}

	sample := Sample{Timestamp: time.Now(), Value: val}
	s.Samples = append(s.Samples, sample)

	if mType == MetricTypeHistogram {
		mc.updateBuckets(s, val)
	}
}

func (mc *MetricsCollector) updateBuckets(s *MetricSeries, val float64) {
	bounds := []float64{10, 50, 100, 500, 1000}
	for _, b := range bounds {
		if val <= b {
			s.Buckets[fmt.Sprintf("<=%.0f", b)]++
		}
	}
	s.Buckets["+Inf"]++
}

// Reset clears all recorded data for a metric.
func (mc *MetricsCollector) Reset(name string) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	delete(mc.series, name)
}

// GetStats computes aggregated stats over the retention window for a named metric.
func (mc *MetricsCollector) GetStats(name string) (*MetricStats, bool) {
	mc.mu.Lock()
	mc.pruneExpiredLocked()
	s, exists := mc.series[name]
	if !exists || len(s.Samples) == 0 {
		mc.mu.Unlock()
		return nil, false
	}
	// Copy samples under lock
	samples := make([]Sample, len(s.Samples))
	copy(samples, s.Samples)
	sType := s.Type
	mc.mu.Unlock()

	values := make([]float64, len(samples))
	var sum float64
	minVal := samples[0].Value
	maxVal := samples[0].Value

	for i, sample := range samples {
		v := sample.Value
		values[i] = v
		sum += v
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	sort.Float64s(values)
	count := len(values)
	avg := sum / float64(count)

	p95 := percentile(values, 0.95)
	p99 := percentile(values, 0.99)

	stats := &MetricStats{
		Name:        name,
		Type:        sType,
		Count:       count,
		Sum:         sum,
		Min:         minVal,
		Max:         maxVal,
		Avg:         avg,
		P95:         p95,
		P99:         p99,
		LastUpdated: samples[count-1].Timestamp,
	}
	return stats, true
}

func percentile(sorted []float64, p float64) float64 {
	if len(sorted) == 0 {
		return 0
	}
	index := p * float64(len(sorted)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))
	if lower == upper {
		return sorted[lower]
	}
	weight := index - float64(lower)
	return sorted[lower]*(1-weight) + sorted[upper]*weight
}

// ListMetricNames returns a slice of all currently tracked metric names.
func (mc *MetricsCollector) ListMetricNames() []string {
	mc.mu.RLock()
	defer mc.mu.RUnlock()

	names := make([]string, 0, len(mc.series))
	for name := range mc.series {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (mc *MetricsCollector) pruneExpiredLocked() {
	cutoff := time.Now().Add(-mc.retentionWindow)
	for name, s := range mc.series {
		var valid []Sample
		for _, sample := range s.Samples {
			if sample.Timestamp.After(cutoff) {
				valid = append(valid, sample)
			}
		}
		if len(valid) == 0 {
			delete(mc.series, name)
		} else {
			s.Samples = valid
		}
	}
}

func (mc *MetricsCollector) startPruningLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-mc.ctx.Done():
			return
		case <-ticker.C:
			mc.mu.Lock()
			mc.pruneExpiredLocked()
			mc.mu.Unlock()
		}
	}
}

// Router builds HTTP endpoints for ingestion, query, and dashboard.
func (mc *MetricsCollector) Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mc.ListMetricNames())

		case http.MethodPost:
			var req struct {
				Name  string     `json:"name"`
				Type  MetricType `json:"type"`
				Value float64    `json:"value"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if req.Name == "" || req.Type == "" {
				http.Error(w, "name and type are required", http.StatusBadRequest)
				return
			}

			mc.Observe(req.Name, req.Type, req.Value)
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintln(w, `{"status":"accepted"}`)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/metrics/", func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(r.URL.Path, "/metrics/")
		if name == "" {
			http.Error(w, "metric name required", http.StatusBadRequest)
			return
		}

		stats, ok := mc.GetStats(name)
		if !ok {
			http.Error(w, "metric not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stats)
	})

	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-r.Context().Done():
				return
			case <-ticker.C:
				names := mc.ListMetricNames()
				var sb strings.Builder
				sb.WriteString("=== METRICS DASHBOARD ===\n")
				for _, name := range names {
					if stats, ok := mc.GetStats(name); ok {
						sb.WriteString(fmt.Sprintf("[%s] %s | Count: %d | Avg: %.2f | Min: %.2f | Max: %.2f | P95: %.2f | P99: %.2f\n",
							stats.Type, stats.Name, stats.Count, stats.Avg, stats.Min, stats.Max, stats.P95, stats.P99))
					}
				}
				fmt.Fprintf(w, "data: %s\n\n", strings.ReplaceAll(sb.String(), "\n", "\\n"))
				flusher.Flush()
			}
		}
	})

	return mux
}

func main() {
	collector := NewMetricsCollector(5 * time.Minute)
	defer collector.Close()

	// Seed sample data
	collector.Observe("http_requests_total", MetricTypeCounter, 1)
	collector.Observe("http_requests_total", MetricTypeCounter, 1)
	collector.Observe("memory_usage_bytes", MetricTypeGauge, 1024*1024*45)
	collector.Observe("request_latency_ms", MetricTypeHistogram, 12.5)
	collector.Observe("request_latency_ms", MetricTypeHistogram, 45.0)

	port := 9091
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: collector.Router(),
	}

	log.Printf("Starting Metrics Collector & Dashboard on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
