package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestSchedulerAddAndExecute(t *testing.T) {
	s := NewScheduler()
	defer s.Stop()

	var counter int32
	retry := RetryPolicy{MaxRetries: 1, InitialDelay: 10 * time.Millisecond}
	
	sched, err := s.AddSchedule("t1", "Test Task", ScheduleTypeRecurring, 50*time.Millisecond, retry, func(ctx context.Context) error {
		atomic.AddInt32(&counter, 1)
		return nil
	})
	if err != nil {
		t.Fatalf("failed to add schedule: %v", err)
	}

	if sched.ID != "t1" || sched.Name != "Test Task" {
		t.Errorf("unexpected schedule fields: %+v", sched)
	}

	time.Sleep(120 * time.Millisecond)

	count := atomic.LoadInt32(&counter)
	if count < 2 {
		t.Errorf("expected task to run at least 2 times, ran %d times", count)
	}

	got, err := s.GetSchedule("t1")
	if err != nil || got.ID != "t1" {
		t.Errorf("GetSchedule failed: %v", err)
	}
	if len(got.History) == 0 {
		t.Error("expected execution history records")
	}
}

func TestSchedulerRetryOnFailure(t *testing.T) {
	s := NewScheduler()
	defer s.Stop()

	var attempts int32
	retry := RetryPolicy{MaxRetries: 2, InitialDelay: 10 * time.Millisecond, BackoffFactor: 1.0}

	s.AddSchedule("t2", "Failing Task", ScheduleTypeOneTime, 20*time.Millisecond, retry, func(ctx context.Context) error {
		atomic.AddInt32(&attempts, 1)
		return errors.New("temporary error")
	})

	time.Sleep(150 * time.Millisecond)

	totalAttempts := atomic.LoadInt32(&attempts)
	if totalAttempts != 3 { // 1 initial + 2 retries
		t.Errorf("expected 3 total attempts, got %d", totalAttempts)
	}

	sched, _ := s.GetSchedule("t2")
	if len(sched.History) == 0 {
		t.Fatalf("expected history record")
	}
	if sched.History[0].Status != StatusFailed {
		t.Errorf("expected status FAILED, got %s", sched.History[0].Status)
	}
}

func TestSchedulerHTTPRouter(t *testing.T) {
	s := NewScheduler()
	defer s.Stop()

	router := s.Router()

	// POST /schedules
	payload := `{"id":"t-http","name":"HTTP Task","type":"one_time","interval_sec":10,"max_retries":1}`
	req := httptest.NewRequest("POST", "/schedules", bytes.NewBufferString(payload))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", rr.Code)
	}

	// GET /schedules
	req = httptest.NewRequest("GET", "/schedules", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}
	var list []*TaskSchedule
	if err := json.Unmarshal(rr.Body.Bytes(), &list); err != nil || len(list) != 1 {
		t.Errorf("failed to unmarshal schedules list: %v, len=%d", err, len(list))
	}

	// GET /schedules/t-http
	req = httptest.NewRequest("GET", "/schedules/t-http", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}

	// DELETE /schedules/t-http
	req = httptest.NewRequest("DELETE", "/schedules/t-http", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("expected 204 No Content, got %d", rr.Code)
	}
}
