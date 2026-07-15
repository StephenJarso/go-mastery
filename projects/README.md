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
