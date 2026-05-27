package concurrency

import (
	"context"
	"sync"
	"errors"
)


var _ = context.Background
var _ sync.Mutex
var _ = errors.New

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

// Exercise 1: Concurrent Sub-range Sum
// Sum integers from 1 to N using multiple concurrent workers (goroutines).
// Each worker sums a sub-range and accumulates to total sum. Use WaitGroup and Mutex.
func ConcurrentSum(n int, numGoroutines int) int {
	// TODO: Implement
	return 0
}

// Exercise 2: Thread-Safe Map
// Implement Set, Get, Delete on SafeMap using sync.RWMutex.
func NewSafeMap() *SafeMap {
	// TODO: Implement
	return nil
}

func (m *SafeMap) Set(key, val string) {
	// TODO: Implement
}

func (m *SafeMap) Get(key string) (string, bool) {
	// TODO: Implement
	return "", false
}

// Exercise 3: Worker Pool with Channels
// Run workers (goroutines) to process input jobs channel, sending doubled values to results channel.
func WorkerPool(jobs <-chan int, results chan<- int, numWorkers int) {
	// TODO: Implement
}
