# Project 05: Metrics Collector & Dashboard

## Problem Statement

Build a metrics collection system that aggregates numeric metrics from multiple sources, stores them in memory with time-windowed retention, and exposes both an API for querying aggregated stats and a simple text-based dashboard. This mirrors real-world observability tools like Prometheus.

## Goals

- Implement a `MetricsCollector` that accepts counter, gauge, and histogram metric types
- Store metrics with configurable time-window retention (e.g., last 5 minutes, last 1 hour)
- Provide aggregation functions: min, max, avg, p95, p99, count, sum
- Expose a REST API for submitting metrics and querying aggregated results
- Build a text-based dashboard that refreshes in real-time using Server-Sent Events (SSE)
- Use goroutines and channels for concurrent metric ingestion and aggregation

## Guide

1. **Define metric types** — Create `Counter`, `Gauge`, and `Histogram` types with `Observe()`, `Value()`, and `Reset()` methods
2. **Build the collector** — Implement `MetricsCollector` with a mutex-protected map of metric name to value, and a background goroutine that prunes expired entries
3. **Implement aggregation** — For each metric, compute min, max, avg, percentiles over configurable time windows
4. **Create HTTP API** — Endpoints for `POST /metrics` (submit), `GET /metrics/{name}` (query), and `GET /metrics` (list all)
5. **Build the dashboard** — Use SSE to push metric updates to a browser-based or terminal-based dashboard that refreshes every second
6. **Add histogram bucketing** — For histogram metrics, track bucket counts and compute percentiles from bucket data

## Key Concepts

- Mutex-protected concurrent map access
- Background goroutines with `time.Ticker` for pruning
- Server-Sent Events (SSE) for real-time updates
- Percentile computation from histogram buckets
- HTTP handler patterns for metrics ingestion and query

## Running

```bash
cd 05-metrics
go run main.go
```

The collector starts on `http://localhost:9091` with endpoints:
- `POST /metrics` — Submit a metric (JSON body with name, type, value)
- `GET /metrics/{name}` — Query aggregated stats for a metric
- `GET /dashboard` — View the real-time text dashboard
- `GET /metrics` — List all tracked metric names

## Testing

```bash
go test -v
```
