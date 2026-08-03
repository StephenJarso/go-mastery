# Phase 7: Unified Capstone Project - Distributed Task Manager 🚀

This folder contains the complete Phase 7 capstone project, which integrates all the concepts learned from Phase 2 through Phase 6 into a single, cohesive, three-part system:

1. **`01-cli-tool`**: A CLI client (`task-cli`) used to submit, list, inspect, and trigger tasks.
2. **`02-rest-api`**: An in-memory REST API HTTP server that exposes endpoints for CRUD operations and manages task state using a mutex.
3. **`03-concurrent-system`**: A concurrent background worker pool that processes task payloads in parallel using channels, goroutines, and timeouts.

---

## 🏗️ Architecture

```
                 +-------------------+
                 |    task-cli       |  (CLI Client)
                 +-------------------+
                           |
                       HTTP Requests
                           |
                           v
                 +-------------------+
                 |   Task Server     |  (REST API / Mutex Safe)
                 +-------------------+
                           |
                     Queue Submission
                           |
                           v
                 +-------------------+
                 |    Worker Pool    |  (Concurrent Goroutines)
                 +-------------------+
```

---

## ⚙️ Running the Project

### 1. Start the REST API Server
Navigate to `02-rest-api/` and run:
```bash
go run server.go
```
The server will start listening on `http://localhost:8080`.

### 2. Use the CLI Client
In another terminal, navigate to `01-cli-tool/`. Build or run the CLI:

#### Add a Task
```bash
go run cli.go add -title "Compress Log Files" -desc "Archive /var/log/*.log" -payload "log-archive-payload"
```

#### List Tasks
```bash
go run cli.go list
```

#### Submit a Task for Processing
Trigger background execution of the task:
```bash
go run cli.go process -id task-1
```

#### Get Task Details
Verify the result after background workers complete processing:
```bash
go run cli.go get -id task-1
```

---

## 🧪 Running Tests

Each module includes robust test coverage:

```bash
# Run tests for CLI tool
cd 01-cli-tool && go test -v

# Run tests for REST API server
cd ../02-rest-api && go test -v

# Run tests for Concurrent worker pool
cd ../03-concurrent-system && go test -v
```

---

## 📦 Additional Projects

Beyond the unified capstone, the following standalone projects explore specific Go patterns and real-world systems:

### Project 04: Task Scheduler

**`04-scheduler/`** — A cron-like task scheduler with retry logic and status API.

- Schedule one-time and recurring tasks
- Retry failed tasks with exponential backoff
- Query scheduled and completed tasks via HTTP API
- Graceful shutdown with context cancellation

See [`04-scheduler/README.md`](04-scheduler/README.md) for the full problem statement and implementation guide.

### Project 05: Metrics Collector & Dashboard

**`05-metrics/`** — A metrics aggregation system with a real-time text dashboard.

- Collect counter, gauge, and histogram metrics
- Compute min, max, avg, p95, p99 aggregations over time windows
- Expose a REST API for metric submission and queries
- Real-time dashboard using Server-Sent Events (SSE)

See [`05-metrics/README.md`](05-metrics/README.md) for the full problem statement and implementation guide.

### Project 06: Event Stream Processor

**`06-event-stream/`** — A publish/subscribe event streaming system with backpressure.

- Topic-based publish/subscribe with wildcard matching
- At-least-once delivery with ACK and retry
- Backpressure via buffered channels
- Monitoring API for throughput and subscriber metrics

See [`06-event-stream/README.md`](06-event-stream/README.md) for the full problem statement and implementation guide.
