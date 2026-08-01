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
	"time"
)

// TaskFunc is the signature of a task execution function.
type TaskFunc func(ctx context.Context) error

// ScheduleType represents whether a schedule is one-time or recurring.
type ScheduleType string

const (
	ScheduleTypeOneTime   ScheduleType = "one_time"
	ScheduleTypeRecurring ScheduleType = "recurring"
)

// ExecutionStatus represents the result of a single task execution.
type ExecutionStatus string

const (
	StatusSuccess ExecutionStatus = "SUCCESS"
	StatusFailed  ExecutionStatus = "FAILED"
)

// ExecutionRecord stores historical metadata for task runs.
type ExecutionRecord struct {
	Timestamp time.Time       `json:"timestamp"`
	Status    ExecutionStatus `json:"status"`
	Error     string          `json:"error,omitempty"`
	Duration  string          `json:"duration"`
	Attempts  int             `json:"attempts"`
}

// RetryPolicy defines how failed tasks should be retried.
type RetryPolicy struct {
	MaxRetries int           `json:"max_retries"`
	InitialDelay time.Duration `json:"initial_delay"`
	BackoffFactor float64    `json:"backoff_factor"`
}

// TaskSchedule represents a scheduled job inside the engine.
type TaskSchedule struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Type        ScheduleType      `json:"type"`
	Interval    time.Duration     `json:"interval"`
	NextRun     time.Time         `json:"next_run"`
	RetryPolicy RetryPolicy       `json:"retry_policy"`
	Status      string            `json:"status"` // "SCHEDULED", "RUNNING", "CANCELLED", "COMPLETED"
	History     []ExecutionRecord `json:"history"`
	
	fn          TaskFunc
	cancelFunc  context.CancelFunc
}

// Scheduler coordinates scheduled tasks.
type Scheduler struct {
	mu        sync.RWMutex
	schedules map[string]*TaskSchedule
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
}

// NewScheduler initializes a new TaskScheduler engine.
func NewScheduler() *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &Scheduler{
		schedules: make(map[string]*TaskSchedule),
		ctx:       ctx,
		cancel:    cancel,
	}
}

// AddSchedule registers a new task into the scheduler.
func (s *Scheduler) AddSchedule(id, name string, schedType ScheduleType, interval time.Duration, retry RetryPolicy, fn TaskFunc) (*TaskSchedule, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.schedules[id]; exists {
		return nil, fmt.Errorf("schedule with ID %s already exists", id)
	}

	taskCtx, taskCancel := context.WithCancel(s.ctx)
	sched := &TaskSchedule{
		ID:          id,
		Name:        name,
		Type:        schedType,
		Interval:    interval,
		NextRun:     time.Now().Add(interval),
		RetryPolicy: retry,
		Status:      "SCHEDULED",
		History:     make([]ExecutionRecord, 0),
		fn:          fn,
		cancelFunc:  taskCancel,
	}

	s.schedules[id] = sched
	s.wg.Add(1)
	go s.runLoop(taskCtx, sched)

	return sched, nil
}

// CancelSchedule stops and removes a scheduled task.
func (s *Scheduler) CancelSchedule(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	sched, exists := s.schedules[id]
	if !exists {
		return errors.New("schedule not found")
	}

	if sched.cancelFunc != nil {
		sched.cancelFunc()
	}
	sched.Status = "CANCELLED"
	return nil
}

// GetSchedule returns details of a single task schedule.
func (s *Scheduler) GetSchedule(id string) (*TaskSchedule, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	sched, exists := s.schedules[id]
	if !exists {
		return nil, errors.New("schedule not found")
	}
	return sched, nil
}

// ListSchedules returns all registered schedules.
func (s *Scheduler) ListSchedules() []*TaskSchedule {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]*TaskSchedule, 0, len(s.schedules))
	for _, sched := range s.schedules {
		list = append(list, sched)
	}
	return list
}

// Stop gracefully stops the scheduler and waits for all active tasks to complete.
func (s *Scheduler) Stop() {
	s.cancel()
	s.wg.Wait()
}

func (s *Scheduler) runLoop(ctx context.Context, sched *TaskSchedule) {
	defer s.wg.Done()

	timer := time.NewTimer(sched.Interval)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-s.ctx.Done():
			return
		case <-timer.C:
			s.executeTask(ctx, sched)

			if sched.Type == ScheduleTypeOneTime {
				s.mu.Lock()
				sched.Status = "COMPLETED"
				s.mu.Unlock()
				return
			}

			s.mu.Lock()
			sched.NextRun = time.Now().Add(sched.Interval)
			s.mu.Unlock()
			timer.Reset(sched.Interval)
		}
	}
}

func (s *Scheduler) executeTask(ctx context.Context, sched *TaskSchedule) {
	s.mu.Lock()
	sched.Status = "RUNNING"
	s.mu.Unlock()

	start := time.Now()
	var execErr error
	attempts := 0
	maxAttempts := sched.RetryPolicy.MaxRetries + 1
	if maxAttempts <= 0 {
		maxAttempts = 1
	}

	delay := sched.RetryPolicy.InitialDelay
	if delay <= 0 {
		delay = 10 * time.Millisecond
	}
	factor := sched.RetryPolicy.BackoffFactor
	if factor <= 0 {
		factor = 2.0
	}

	for attempts < maxAttempts {
		attempts++
		execErr = sched.fn(ctx)
		if execErr == nil {
			break
		}

		if attempts < maxAttempts {
			select {
			case <-ctx.Done():
				return
			case <-time.After(delay):
				delay = time.Duration(float64(delay) * factor)
			}
		}
	}

	duration := time.Since(start)
	status := StatusSuccess
	errStr := ""
	if execErr != nil {
		status = StatusFailed
		errStr = execErr.Error()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	sched.Status = "SCHEDULED"
	sched.History = append(sched.History, ExecutionRecord{
		Timestamp: start,
		Status:    status,
		Error:     errStr,
		Duration:  duration.String(),
		Attempts:  attempts,
	})
}

// Router sets up HTTP endpoints for managing and querying schedules.
func (s *Scheduler) Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/schedules", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(s.ListSchedules())

		case http.MethodPost:
			var req struct {
				ID          string       `json:"id"`
				Name        string       `json:"name"`
				Type        ScheduleType `json:"type"`
				IntervalSec int          `json:"interval_sec"`
				MaxRetries  int          `json:"max_retries"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if req.ID == "" {
				req.ID = fmt.Sprintf("task-%d", time.Now().UnixNano())
			}
			if req.IntervalSec <= 0 {
				req.IntervalSec = 1
			}

			retry := RetryPolicy{
				MaxRetries:    req.MaxRetries,
				InitialDelay:  50 * time.Millisecond,
				BackoffFactor: 1.5,
			}

			taskFunc := func(ctx context.Context) error {
				log.Printf("Executing scheduled task: %s (%s)", req.Name, req.ID)
				return nil
			}

			sched, err := s.AddSchedule(req.ID, req.Name, req.Type, time.Duration(req.IntervalSec)*time.Second, retry, taskFunc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(sched)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/schedules/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/schedules/")
		if id == "" {
			http.Error(w, "missing task ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			sched, err := s.GetSchedule(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(sched)

		case http.MethodDelete:
			if err := s.CancelSchedule(id); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}

func main() {
	scheduler := NewScheduler()

	// Register sample tasks
	scheduler.AddSchedule("sys-backup", "System Backup", ScheduleTypeRecurring, 5*time.Second, RetryPolicy{MaxRetries: 2, InitialDelay: 100 * time.Millisecond}, func(ctx context.Context) error {
		log.Println("[Backup Task] Running periodic database backup...")
		return nil
	})

	scheduler.AddSchedule("cleanup-logs", "Log Cleanup", ScheduleTypeOneTime, 2*time.Second, RetryPolicy{MaxRetries: 1}, func(ctx context.Context) error {
		log.Println("[Cleanup Task] Cleaning temporary logs...")
		return nil
	})

	port := 9090
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: scheduler.Router(),
	}

	log.Printf("Starting Task Scheduler server on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
