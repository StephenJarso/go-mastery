package main

import (
	"fmt"
	"sync"
	"time"
)

// While channels are excellent for communication and orchestration (orchestrating ownership),
// low-level synchronization primitives from the "sync" package are better suited for protecting
// shared state (such as cache maps, counters, or singletons).

// Counter represents a thread-safe shared counter.
type Counter struct {
	// A sync.Mutex is used to provide mutual exclusion.
	// It guarantees that only one goroutine can access the critical section at a time.
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter.
func (c *Counter) Increment() {
	c.mu.Lock()
	// Lock blocks until the mutex is available.
	// We defer Unlock() to ensure it's ALWAYS released, preventing deadlocks.
	defer c.mu.Unlock()
	c.value++
}

// Value safely returns the current counter value.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// MutexExample demonstrates basic mutual exclusion.
func MutexExample() {
	fmt.Println("=== sync.Mutex ===")

	var wg sync.WaitGroup
	counter := &Counter{}

	// Spawn 1000 goroutines to increment the counter concurrently.
	// Without a Mutex, this would cause a data race, leading to unpredictable results.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter Value (expected 1000): %d\n", counter.Value())
}

// SafeMap represents a thread-safe map protected by sync.RWMutex.
// sync.RWMutex allows any number of readers OR exactly one writer, which is
// highly efficient for read-heavy workloads.
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]string),
	}
}

// Set writes a key-value pair.
func (sm *SafeMap) Set(key, val string) {
	sm.mu.Lock() // Writer lock
	defer sm.mu.Unlock()
	sm.data[key] = val
}

// Get reads a value by key.
func (sm *SafeMap) Get(key string) (string, bool) {
	sm.mu.RLock() // Reader lock
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}

// RWMutexExample demonstrates concurrent reads and a write.
func RWMutexExample() {
	fmt.Println("\n=== sync.RWMutex ===")

	sm := NewSafeMap()
	sm.Set("status", "Initialized")

	var wg sync.WaitGroup

	// Spawn 5 reader goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// RLock allows concurrent readers
			val, _ := sm.Get("status")
			fmt.Printf("Reader %d: status = %s\n", id, val)
		}(i)
	}

	// Spawn 1 writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond) // Let readers run first
		sm.Set("status", "Running")
		fmt.Println("Writer: Changed status to 'Running'")
	}()

	wg.Wait()
	finalStatus, _ := sm.Get("status")
	fmt.Printf("Final Status: %s\n", finalStatus)
}

// OnceExample demonstrates sync.Once.
// sync.Once guarantees that a function is executed exactly once, regardless of
// how many times it's called or how many goroutines call it concurrently.
func OnceExample() {
	fmt.Println("\n=== sync.Once ===")

	var once sync.Once
	var wg sync.WaitGroup

	initializeDB := func() {
		fmt.Println("DATABASE INITIALIZED! (This message should print exactly once)")
	}

	// Call initializeDB concurrently from 10 goroutines.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Only the first call to Do will invoke initializeDB.
			// All other calls will block until the first one completes, then return immediately.
			once.Do(initializeDB)
			// (Safe to perform database operations here)
		}(i)
	}

	wg.Wait()
}

// CondExample demonstrates sync.Cond.
// sync.Cond implements a condition variable, a rendezvous point for goroutines
// waiting for or announcing the occurrence of an event.
func CondExample() {
	fmt.Println("\n=== sync.Cond ===")

	var mu sync.Mutex
	// sync.Cond is associated with a Locker (usually a Mutex)
	cond := sync.NewCond(&mu)
	ready := false

	var wg sync.WaitGroup

	// Worker goroutines waiting for readiness
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cond.L.Lock()
			// Keep waiting while the condition is false.
			// A loop is required because of spurious wakeups.
			for !ready {
				fmt.Printf("Worker %d: Waiting for ready signal...\n", id)
				cond.Wait() // Wait releases the lock and suspends execution
				// When Wait returns, it has re-acquired the lock!
			}
			fmt.Printf("Worker %d: Running now!\n", id)
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(50 * time.Millisecond)

	cond.L.Lock()
	ready = true
	fmt.Println("Coordinator: All systems ready! Broadcasting signal...")
	cond.Broadcast() // Wake up all waiting workers (Signal() would only wake up one)
	cond.L.Unlock()

	wg.Wait()
}

func main() {
	MutexExample()
	RWMutexExample()
	OnceExample()
	CondExample()
}
