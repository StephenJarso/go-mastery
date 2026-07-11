package concurrency

import (
	"context"
	"sync"
	"testing"
)

func TestConcurrentSum(t *testing.T) {
	// Sum from 1 to 100 is 5050
	got := ConcurrentSum(100, 4)
	if got != 5050 {
		t.Errorf("ConcurrentSum(100, 4) = %d; want 5050", got)
	}

	gotZero := ConcurrentSum(0, 5)
	if gotZero != 0 {
		t.Errorf("ConcurrentSum(0, 5) = %d; want 0", gotZero)
	}
}

func TestPracticeMap(t *testing.T) {
	pm := NewPracticeMap()

	// Verify basic operations
	pm.Set("name", "Stephen")
	val, ok := pm.Get("name")
	if !ok || val != "Stephen" {
		t.Errorf("Expected 'Stephen', got %v", val)
	}

	pm.Delete("name")
	_, ok = pm.Get("name")
	if ok {
		t.Errorf("Expected key to be deleted")
	}

	// Run concurrent reads and writes to test thread safety (use go test -race)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func(id int) {
			defer wg.Done()
			pm.Set("key", id)
		}(i)
		go func() {
			defer wg.Done()
			pm.Get("key")
		}()
	}
	wg.Wait()
}

func TestProcessTasksWithContext(t *testing.T) {
	// Test Case 1: Finish all tasks successfully
	tasks := make(chan string, 5)
	tasks <- "task1"
	tasks <- "task2"
	close(tasks)

	processed := []string{}
	processor := func(t string) {
		processed = append(processed, t)
	}

	err := ProcessTasksWithContext(context.Background(), tasks, processor)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if len(processed) != 2 || processed[0] != "task1" || processed[1] != "task2" {
		t.Errorf("Processed tasks mismatch: %v", processed)
	}

	// Test Case 2: Cancellation during execution
	ctx, cancel := context.WithCancel(context.Background())
	slowTasks := make(chan string, 5)
	slowTasks <- "taskA"

	slowProcessor := func(t string) {
		cancel() // Cancel context immediately on first task
	}

	err = ProcessTasksWithContext(ctx, slowTasks, slowProcessor)
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled error, got %v", err)
	}
}
