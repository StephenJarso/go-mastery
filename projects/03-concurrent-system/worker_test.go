package concurrent_system

import (
	"sync"
	"testing"
	"time"

	"go-mastery/projects"
)

func TestWorkerPool_Basic(t *testing.T) {
	var mu sync.Mutex
	completedTasks := make(map[string]*projects.Task)
	var wg sync.WaitGroup

	onComplete := func(task *projects.Task) {
		mu.Lock()
		completedTasks[task.ID] = task
		mu.Unlock()
		wg.Done()
	}

	wp := NewWorkerPool(2, 5, onComplete)
	wp.Start()
	defer wp.Shutdown()

	task1 := &projects.Task{
		ID:      "t1",
		Title:   "Task 1",
		Payload: "hello",
	}
	task2 := &projects.Task{
		ID:      "t2",
		Title:   "Task 2",
		Payload: "world",
	}

	wg.Add(2)
	if !wp.Submit(task1) {
		t.Error("Failed to submit task 1")
	}
	if !wp.Submit(task2) {
		t.Error("Failed to submit task 2")
	}

	// Wait with timeout
	c := make(chan struct{})
	go func() {
		wg.Wait()
		close(c)
	}()

	select {
	case <-c:
		// Done
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for tasks to complete")
	}

	mu.Lock()
	defer mu.Unlock()

	if len(completedTasks) != 2 {
		t.Errorf("Expected 2 completed tasks, got %d", len(completedTasks))
	}

	if completedTasks["t1"].Status != projects.StatusCompleted {
		t.Errorf("Expected status completed for task 1, got %s", completedTasks["t1"].Status)
	}

	if completedTasks["t2"].Status != projects.StatusCompleted {
		t.Errorf("Expected status completed for task 2, got %s", completedTasks["t2"].Status)
	}
}
