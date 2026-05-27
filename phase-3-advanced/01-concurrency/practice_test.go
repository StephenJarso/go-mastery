package concurrency

import (
	"testing"
	"context"
	"sync"
	"time"
)


func TestConcurrentSum(t *testing.T) {
	res := ConcurrentSum(100, 4)
	if res != 5050 {
		t.Errorf("expected 5050, got %d", res)
	}
}

func TestSafeMap(t *testing.T) {
	m := NewSafeMap()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			m.Set("key", "val")
			m.Get("key")
		}(i)
	}
	wg.Wait()
	val, ok := m.Get("key")
	if !ok || val != "val" {
		t.Errorf("expected val, got %s", val)
	}
}

func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	WorkerPool(jobs, results, 3)
	close(results)
	sum := 0
	for r := range results {
		sum += r
	}
	if sum != 30 {
		t.Errorf("expected sum of doubles 30, got %d", sum)
	}
}
