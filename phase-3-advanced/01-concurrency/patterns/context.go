package patterns

import (
	"context"
	"fmt"
	"time"
)

// The "context" package makes it easy to pass request-scoped values, cancelation signals,
// and deadlines across API boundaries and between goroutines.
//
// In Go, context is typically passed as the first parameter to functions:
//   func DoSomething(ctx context.Context, ...)

// ContextCancellation demonstrates basic cancellation propagation.
func ContextCancellation() {
	fmt.Println("=== Context Cancellation ===")

	// Create a parent context and a cancel function.
	ctx, cancel := context.WithCancel(context.Background())

	// Spawn a worker that runs indefinitely until canceled.
	go func() {
		for {
			select {
			case <-ctx.Done():
				// Done() returns a channel that is closed when cancel() is called.
				fmt.Println("Worker: Received cancellation signal. Cleaning up and exiting...")
				return
			default:
				fmt.Println("Worker: Working...")
				time.Sleep(20 * time.Millisecond)
			}
		}
	}()

	time.Sleep(50 * time.Millisecond)
	fmt.Println("Main: Triggering cancellation...")
	cancel() // This closes ctx.Done(), notifying the worker goroutine.

	// Give the worker time to print cleanup message.
	time.Sleep(10 * time.Millisecond)
}

// ContextTimeout demonstrates canceling operations that exceed a deadline.
func ContextTimeout() {
	fmt.Println("\n=== Context Timeout ===")

	// Create a context that automatically cancels after 50 milliseconds.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel() // Good practice to release resources when operation completes early.

	// Simulated database query that takes 100 milliseconds.
	queryDatabase := func(c context.Context) (string, error) {
		select {
		case <-time.After(100 * time.Millisecond):
			return "Query Results", nil
		case <-c.Done():
			// c.Err() returns context.DeadlineExceeded or context.Canceled.
			return "", c.Err()
		}
	}

	fmt.Println("Main: Executing database query...")
	result, err := queryDatabase(ctx)
	if err != nil {
		fmt.Printf("Main: Query failed: %v\n", err)
	} else {
		fmt.Printf("Main: Query success: %s\n", result)
	}
}

// ContextValues demonstrates propagating request-scoped values.
func ContextValues() {
	fmt.Println("\n=== Context Values ===")

	// Define a custom key type to prevent collisions with other packages.
	type favContextKey string
	const keyUserID favContextKey = "userID"

	// Create a context carrying a value.
	ctx := context.WithValue(context.Background(), keyUserID, "user_12345")

	// Helper function that accesses the value
	printUserSession := func(c context.Context) {
		if val := c.Value(keyUserID); val != nil {
			fmt.Printf("Session: Found User ID = %v\n", val)
		} else {
			fmt.Println("Session: No User ID found in context.")
		}
	}

	printUserSession(ctx)
	printUserSession(context.Background()) // No value present
}

func RunContextDemo() {
	ContextCancellation()
	ContextTimeout()
	ContextValues()
}
