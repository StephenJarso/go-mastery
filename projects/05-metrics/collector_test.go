package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMetricsCollectorObservationAndStats(t *testing.T) {
	mc := NewMetricsCollector(1 * time.Minute)
	defer mc.Close()

	mc.Observe("test_counter", MetricTypeCounter, 10)
	mc.Observe("test_counter", MetricTypeCounter, 20)
	mc.Observe("test_counter", MetricTypeCounter, 30)

	stats, ok := mc.GetStats("test_counter")
	if !ok {
		t.Fatalf("expected metrics for test_counter")
	}

	if stats.Count != 3 || stats.Sum != 60 || stats.Min != 10 || stats.Max != 30 || stats.Avg != 20 {
		t.Errorf("unexpected stats: %+v", stats)
	}

	if stats.P95 != 29 || stats.P99 != 29.8 {
		// Percentiles calculation check
		if stats.P95 < 25 || stats.P99 < 25 {
			t.Errorf("percentile check failed: P95=%f, P99=%f", stats.P95, stats.P99)
		}
	}
}

func TestMetricsRetentionPruning(t *testing.T) {
	mc := NewMetricsCollector(50 * time.Millisecond)
	defer mc.Close()

	mc.Observe("temp_metric", MetricTypeGauge, 100)
	_, ok := mc.GetStats("temp_metric")
	if !ok {
		t.Fatalf("expected metric before pruning")
	}

	time.Sleep(150 * time.Millisecond)

	_, ok = mc.GetStats("temp_metric")
	if ok {
		t.Errorf("expected metric to be pruned after retention window")
	}
}

func TestMetricsHTTPRouter(t *testing.T) {
	mc := NewMetricsCollector(1 * time.Minute)
	defer mc.Close()

	router := mc.Router()

	// POST /metrics
	payload := `{"name":"api_latency","type":"histogram","value":42.5}`
	req := httptest.NewRequest("POST", "/metrics", bytes.NewBufferString(payload))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Fatalf("expected 202 Accepted, got %d", rr.Code)
	}

	// GET /metrics
	req = httptest.NewRequest("GET", "/metrics", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}
	var names []string
	if err := json.Unmarshal(rr.Body.Bytes(), &names); err != nil || len(names) != 1 || names[0] != "api_latency" {
		t.Errorf("unexpected metrics list response: %v, names=%v", err, names)
	}

	// GET /metrics/api_latency
	req = httptest.NewRequest("GET", "/metrics/api_latency", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}
	var stats MetricStats
	if err := json.Unmarshal(rr.Body.Bytes(), &stats); err != nil || stats.Name != "api_latency" || stats.Avg != 42.5 {
		t.Errorf("unexpected metric stats: %+v", stats)
	}
}
