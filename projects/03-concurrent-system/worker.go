package concurrent_system

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"go-mastery/projects"
)

// WorkerPool manages a pool of goroutines processing tasks concurrently
type WorkerPool struct {
	numWorkers int
	taskQueue  chan *projects.Task
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	onComplete func(*projects.Task)
}

// NewWorkerPool creates a new concurrent worker pool
func NewWorkerPool(numWorkers int, queueSize int, onComplete func(*projects.Task)) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		numWorkers: numWorkers,
		taskQueue:  make(chan *projects.Task, queueSize),
		ctx:        ctx,
		cancel:     cancel,
		onComplete: onComplete,
	}
}

// Start spawns the worker goroutines
func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker represents a single background processing goroutine
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for {
		select {
		case <-wp.ctx.Done():
			return
		case task, ok := <-wp.taskQueue:
			if !ok {
				return
			}
			wp.processTask(id, task)
		}
	}
}

// processTask handles the execution of a single task
func (wp *WorkerPool) processTask(workerID int, task *projects.Task) {
	// Update status to processing
	task.Status = projects.StatusProcessing
	task.UpdatedAt = time.Now()

	// Simulate work duration based on payload length (min 100ms)
	workDuration := time.Duration(100+len(task.Payload)*10) * time.Millisecond
	
	// Create context with timeout for task processing
	taskCtx, taskCancel := context.WithTimeout(wp.ctx, 2*time.Second)
	defer taskCancel()

	// Simulate running the task
	errChan := make(chan error, 1)
	go func() {
		time.Sleep(workDuration)
		errChan <- nil
	}()

	select {
	case <-taskCtx.Done():
		task.Status = projects.StatusFailed
		task.Result = fmt.Sprintf("failed: execution timeout by worker %d", workerID)
	case err := <-errChan:
		if err != nil {
			task.Status = projects.StatusFailed
			task.Result = fmt.Sprintf("failed: %v", err)
		} else {
			task.Status = projects.StatusCompleted
			task.Result = fmt.Sprintf("success: processed %q in worker %d", strings.ToUpper(task.Payload), workerID)
		}
	}

	task.UpdatedAt = time.Now()

	// Invoke callback to persist status or notify API
	if wp.onComplete != nil {
		wp.onComplete(task)
	}
}

// Submit enqueues a task for processing. Returns false if the queue is full.
func (wp *WorkerPool) Submit(task *projects.Task) bool {
	select {
	case wp.taskQueue <- task:
		task.Status = projects.StatusPending
		return true
	default:
		return false // Queue is full
	}
}

// Shutdown stops all workers and waits for outstanding tasks to finish
func (wp *WorkerPool) Shutdown() {
	wp.cancel()
	close(wp.taskQueue)
	wp.wg.Wait()
}
