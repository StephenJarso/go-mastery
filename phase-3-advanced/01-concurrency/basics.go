package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// In Go, concurrency is built directly into the language runtime.
// A "goroutine" is a lightweight execution thread managed by the Go runtime, not the OS.
// Goroutines are multiplexed onto a small number of OS threads. They start with a very small stack
// (typically 2KB) that grows and shrinks as needed.

// PrintNumbers is a helper function to print numbers with a delay.
func PrintNumbers(id string, wg *sync.WaitGroup) {
	// Call Done on the WaitGroup when this function returns.
	// defer guarantees that Done() is called even if the function panics.
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("Goroutine %s: %d\n", id, i)
		// Sleep pauses the current goroutine. This allows other goroutines to run.
		time.Sleep(10 * time.Millisecond)
	}
}

// GoroutineLifecycle demonstrates how goroutines are launched and synchronized.
func GoroutineLifecycle() {
	fmt.Println("=== Goroutine Lifecycle & WaitGroup ===")

	// A sync.WaitGroup is used to wait for a collection of goroutines to finish.
	// It operates like a thread-safe counter.
	var wg sync.WaitGroup

	// We are launching two goroutines, so we add 2 to the WaitGroup counter.
	wg.Add(2)

	fmt.Println("Starting goroutines...")

	// Launching goroutines using the "go" keyword.
	// These run concurrently in the background.
	go PrintNumbers("A", &wg)
	go PrintNumbers("B", &wg)

	fmt.Println("Goroutines launched. Waiting for them to complete...")

	// Wait blocks the calling goroutine (here, the main goroutine) until
	// the WaitGroup counter goes down to zero.
	wg.Wait()

	fmt.Println("All goroutines completed!")
}

// ClosureCaptureIssue demonstrates a very common bug in Go: capturing loop variables in a closure.
func ClosureCaptureIssue() {
	fmt.Println("\n=== Closure Capture Issue ===")

	var wg sync.WaitGroup

	// Let's launch 5 goroutines in a loop.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// BUGGY APPROACH:
		// The goroutine closure captures the variable 'i' by reference, not by value.
		// By the time the goroutines actually run, the loop might have finished,
		// and 'i' will likely have the value 6 for all of them.
		go func() {
			defer wg.Done()
			// This might print '6' repeatedly instead of 1, 2, 3, 4, 5.
			// (Note: in Go 1.22+, loop variables are allocated per iteration,
			// which mitigates this, but passing parameters is still best practice for clarity.)
			fmt.Printf("Buggy Goroutine: i = %d\n", i)
		}()
	}
	wg.Wait()

	fmt.Println("\n=== Correct Closure Capture (Parameters) ===")
	// CORRECT APPROACH:
	// Pass the variable as an argument to the goroutine function.
	// This evaluates the value at the time the go statement is executed and copies it.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Printf("Correct Goroutine: val = %d\n", val)
		}(i) // Pass 'i' here
	}
	wg.Wait()
}

func RunBasicsDemo() {
	GoroutineLifecycle()
	ClosureCaptureIssue()
}
