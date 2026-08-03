# Project 04: Task Scheduler

## Problem Statement

Build a task scheduler that executes tasks at specified intervals or times, similar to a cron system. The scheduler should support one-time and recurring tasks, handle task failures with retry logic, and provide a status API to query scheduled and completed tasks.

## Goals

- Implement a scheduler that accepts tasks with cron-like expressions or delay durations
- Support one-time and recurring task execution
- Handle task failures with configurable retry policies (max retries, backoff)
- Expose a status endpoint to list scheduled, running, and completed tasks
- Ensure thread-safe access to the task registry using mutexes
- Use channels and context cancellation for graceful shutdown

## Guide

1. **Define the Schedule type** — Create a `Schedule` struct with fields for interval, next run time, and recurrence rule
2. **Build the Scheduler** — Implement a `Scheduler` struct that holds a map of scheduled tasks and uses a `time.Ticker` or `time.AfterFunc` for triggering execution
3. **Add retry logic** — On task failure, retry with exponential backoff up to a configurable max retries
4. **Expose HTTP endpoints** — Add endpoints to list schedules, add new ones, and query task results
5. **Graceful shutdown** — Use `context.WithCancel` to stop the scheduler and wait for in-flight tasks to complete

## Key Concepts

- `time.Ticker` and `time.AfterFunc` for scheduling
- `context.Context` for cancellation and timeouts
- `sync.Mutex` for concurrent access to the schedule registry
- Exponential backoff for retries
- HTTP handler for status queries

## Running

```bash
cd 04-scheduler
go run scheduler.go
```

The scheduler starts on `http://localhost:9090` with endpoints:
- `GET /schedules` — List all scheduled tasks
- `POST /schedules` — Add a new scheduled task
- `GET /schedules/{id}` — Get schedule details and execution history
- `DELETE /schedules/{id}` — Cancel a scheduled task

## Testing

```bash
go test -v
```
