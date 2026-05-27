package solutions

import (
	"sync"
)


type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func ConcurrentSum(n int, numGoroutines int) int {
	if n <= 0 || numGoroutines <= 0 {
		return 0
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	total := 0
	chunkSize := n / numGoroutines
	if chunkSize == 0 {
		chunkSize = 1
		numGoroutines = n
	}

	for i := 0; i < numGoroutines; i++ {
		start := i*chunkSize + 1
		end := (i + 1) * chunkSize
		if i == numGoroutines-1 {
			end = n
		}
		wg.Add(1)
		go func(s, e int) {
			defer wg.Done()
			subSum := 0
			for j := s; j <= e; j++ {
				subSum += j
			}
			mu.Lock()
			total += subSum
			mu.Unlock()
		}(start, end)
	}
	wg.Wait()
	return total
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]string),
	}
}

func (m *SafeMap) Set(key, val string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}

func (m *SafeMap) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.data[key]
	return val, ok
}

func WorkerPool(jobs <-chan int, results chan<- int, numWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- job * 2
			}
		}()
	}
	wg.Wait()
}
