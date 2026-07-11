package concurrency

import (
	"context"
	"sync"
	"testing"
	"time"
)

// TestWaitGroup verifies that sync.WaitGroup correctly waits for goroutines.
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	sum := 0

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			mu.Lock()
			sum += val
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	expected := 55 // sum of 1 to 10
	if sum != expected {
		t.Errorf("Expected sum to be %d, got %d", expected, sum)
	}
}

// TestUnbufferedChannel demonstrates simple synchronization via channels.
func TestUnbufferedChannel(t *testing.T) {
	ch := make(chan bool)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch <- true
	}()

	select {
	case val := <-ch:
		if !val {
			t.Errorf("Expected true from channel, got false")
		}
	case <-time.After(100 * time.Millisecond):
		t.Errorf("Timeout waiting for channel receive")
	}
}

// TestMutexCounter verifies the thread-safety of our Counter implementation.
func TestMutexCounter(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	numGoroutines := 100
	incrementsPerGoroutine := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	expected := numGoroutines * incrementsPerGoroutine
	if counter.Value() != expected {
		t.Errorf("Expected counter to be %d, got %d", expected, counter.Value())
	}
}

// TestContextTimeout verifies that a context timeout triggers correctly.
func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(50 * time.Millisecond):
		t.Error("Expected timeout to occur, but operation completed")
	case <-ctx.Done():
		if ctx.Err() != context.DeadlineExceeded {
			t.Errorf("Expected error to be context.DeadlineExceeded, got %v", ctx.Err())
		}
	}
}
