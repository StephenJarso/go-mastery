# Project 06: Event Stream Processor

## Problem Statement

Build an event streaming and processing system that accepts events from multiple producers, routes them to subscribers based on topic filters, and processes them concurrently with guaranteed delivery semantics. This mirrors real-world systems like Apache Kafka or NATS.

## Goals

- Implement an `EventBus` that supports publish/subscribe patterns with topic-based routing
- Allow multiple producers to publish events concurrently without data loss
- Support wildcard topic subscriptions (e.g., `orders.*` matches `orders.created`, `orders.updated`)
- Guarantee at-least-once delivery with configurable acknowledgment timeouts
- Implement backpressure using buffered channels to prevent subscriber overload
- Provide a monitoring API that reports throughput, lag, and active subscriber counts
- Use context cancellation for graceful shutdown of producers and subscribers

## Guide

1. **Define the Event type** — Create an `Event` struct with ID, topic, payload, timestamp, and retry count
2. **Build the EventBus** — Implement a central `EventBus` with a map of topic patterns to subscriber channels, protected by a `sync.RWMutex`
3. **Implement topic matching** — Support wildcard matching (`*` matches one segment, `>` matches all remaining segments) using a simple pattern matcher
4. **Add producer support** — Producers publish events via `Publish(topic, payload)` which routes to all matching subscriber channels
5. **Add subscriber support** — Subscribers register with `Subscribe(topic)` and receive events on a channel; they ACK after processing
6. **Implement backpressure** — Use buffered channels with configurable capacity; when full, producers block or drop based on a configurable policy
7. **Build monitoring API** — Expose HTTP endpoints for throughput stats, subscriber counts, and channel depths
8. **Graceful shutdown** — Use `context.WithCancel` to drain in-flight events before stopping

## Key Concepts

- Publish/Subscribe pattern with topic routing
- Wildcard pattern matching for topics
- Buffered channels for backpressure
- `sync.RWMutex` for concurrent map access
- Context cancellation for graceful shutdown
- At-least-once delivery semantics with ACK/retry

## Running

```bash
cd 06-event-stream
go run main.go
```

The event stream processor starts on `http://localhost:9092` with endpoints:
- `POST /publish` — Publish an event (JSON body with topic and payload)
- `GET /subscribe?topic=orders.*` — Subscribe to a topic (SSE stream)
- `GET /stats` — View throughput and subscriber metrics
- `GET /topics` — List all active topics

## Testing

```bash
go test -v
```
