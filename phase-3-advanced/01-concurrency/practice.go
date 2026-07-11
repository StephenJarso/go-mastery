package concurrency

import (
	"context"
	"sync"
)

// PRACTICE EXERCISE #1: Concurrent Sum
// Implement a function that calculates the sum of integers from 1 to N
// using multiple concurrent goroutines. Each goroutine should sum a sub-range
// and return the sub-total. Use sync.WaitGroup and a Mutex (or channel) to sum them.

func ConcurrentSum(n int, numGoroutines int) int {
	if n <= 0 || numGoroutines <= 0 {
		return 0
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalSum := 0

	// Calculate the range size for each goroutine
	rangeSize := n / numGoroutines
	if rangeSize == 0 {
		rangeSize = 1
		numGoroutines = n
	}

	for i := 0; i < numGoroutines; i++ {
		start := i*rangeSize + 1
		end := (i + 1) * rangeSize
		if i == numGoroutines-1 {
			end = n // Ensure we cover up to n
		}

		wg.Add(1)
		go func(s, e int) {
			defer wg.Done()
			subTotal := 0
			for val := s; val <= e; val++ {
				subTotal += val
			}

			mu.Lock()
			totalSum += subTotal
			mu.Unlock()
		}(start, end)
	}

	wg.Wait()
	return totalSum
}

// PRACTICE EXERCISE #2: Thread-Safe Concurrent Map
// Implement a simple concurrent-safe map structure called PracticeMap.
// It should support Concurrent Get, Set, and Delete using sync.RWMutex.

type PracticeMap struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewPracticeMap() *PracticeMap {
	return &PracticeMap{
		data: make(map[string]interface{}),
	}
}

func (pm *PracticeMap) Set(key string, val interface{}) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.data[key] = val
}

func (pm *PracticeMap) Get(key string) (interface{}, bool) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	val, ok := pm.data[key]
	return val, ok
}

func (pm *PracticeMap) Delete(key string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	delete(pm.data, key)
}

// PRACTICE EXERCISE #3: Context-Aware Worker
// Implement a worker function that processes tasks from a channel.
// The worker should stop processing and return context.Canceled if the context is canceled.
// If the tasks channel is closed, the worker should exit cleanly returning nil.

func ProcessTasksWithContext(ctx context.Context, tasks <-chan string, processor func(string)) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case task, ok := <-tasks:
			if !ok {
				// channel closed
				return nil
			}
			processor(task)
		}
	}
}
