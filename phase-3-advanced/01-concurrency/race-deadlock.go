package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// A data race occurs when two or more goroutines access the same memory location
// concurrently, at least one of these accesses is a write, and there is no
// synchronization (like a mutex or channel) to coordinate the accesses.
//
// Go includes a powerful data race detector. Run tests or build with the `-race` flag:
//   go run -race race-deadlock.go
//   go test -race ./...

// DemonstrateRaceCondition shows a simple data race.
// Note: If you run this with the race detector enabled, it will print a warning and terminate.
func DemonstrateRaceCondition() {
	fmt.Println("=== Data Race Demonstration ===")
	
	var sharedCounter int
	var wg sync.WaitGroup

	// We spawn 2 goroutines that read and write to sharedCounter concurrently.
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sharedCounter++ // Read-modify-write
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sharedCounter++ // Read-modify-write
		}
	}()

	wg.Wait()
	fmt.Printf("Finished. Counter value: %d (might not be 2000 due to race)\n", sharedCounter)
}

// DemonstrateDeadlock shows how a deadlock can occur.
// A deadlock occurs when a set of goroutines are blocked because each goroutine
// is holding a resource and waiting for another resource held by some other goroutine.
func DemonstrateDeadlock() {
	fmt.Println("\n=== Deadlock Demonstration ===")

	var mu1 sync.Mutex
	var mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	// Goroutine 1 acquires mu1, then mu2
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("Goroutine 1: Acquired Lock 1")
		
		time.Sleep(10 * time.Millisecond) // force switch

		fmt.Println("Goroutine 1: Waiting for Lock 2...")
		mu2.Lock() // blocks indefinitely if Goroutine 2 holds mu2
		fmt.Println("Goroutine 1: Acquired Lock 2")
		
		mu2.Unlock()
		mu1.Unlock()
	}()

	// Goroutine 2 acquires mu2, then mu1 (LOCK ORDER VIOLATION)
	go func() {
		defer wg.Done()
		mu2.Lock()
		fmt.Println("Goroutine 2: Acquired Lock 2")

		time.Sleep(10 * time.Millisecond) // force switch

		fmt.Println("Goroutine 2: Waiting for Lock 1...")
		mu1.Lock() // blocks indefinitely if Goroutine 1 holds mu1
		fmt.Println("Goroutine 2: Acquired Lock 1")

		mu1.Unlock()
		mu2.Unlock()
	}()

	// We'll set a timeout to print a message since Go's deadlock detector
	// only triggers if ALL goroutines in the system are blocked.
	// Since the main goroutine would otherwise wait forever on wg.Wait(),
	// the whole program will crash with "fatal error: all goroutines are asleep - deadlock!".
	// To let this script finish gracefully for demonstration, we will run the wait in a goroutine
	// and use a timeout select block.
	
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Completed without deadlock (unlikely!).")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("TIMEOUT! A deadlock occurred: Goroutine 1 holds Lock 1 and wants Lock 2; Goroutine 2 holds Lock 2 and wants Lock 1.")
		fmt.Println("Prevention Tip: Always acquire locks in the same order across all goroutines!")
	}
}

func RunRaceDeadlockDemo() {
	// Running DemonstrateRaceCondition shows why we need Mutex/Channels.
	DemonstrateRaceCondition()
	// Running DemonstrateDeadlock shows lock ordering issue.
	DemonstrateDeadlock()
}
