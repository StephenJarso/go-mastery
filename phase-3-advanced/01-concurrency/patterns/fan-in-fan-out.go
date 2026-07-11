package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Fan-out, Fan-in is a design pattern used to distribute workload and combine results.
// - Fan-out: Multiple goroutines read from the same channel until it's closed, parallelizing a task.
// - Fan-in: A coordinator goroutine multiplexes multiple result channels into a single output channel.

// generator creates a stream of integers.
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// squareWorker reads from the input channel, squares the numbers, and sends them to a results channel.
func squareWorker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			// Simulate intensive computation
			time.Sleep(20 * time.Millisecond)
			out <- n * n
		}
		close(out)
	}()
	return out
}

// fanIn combines multiple channels into a single output channel.
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.
	// This goroutine reads from c and sends to out.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func RunFanDemo() {
	fmt.Println("=== Fan-out / Fan-in Pattern ===")

	// 1. Generate inputs
	in := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// 2. Fan-out: Distribute work across multiple concurrent workers.
	// Each worker operates on the SAME input channel.
	worker1 := squareWorker(in)
	worker2 := squareWorker(in)
	worker3 := squareWorker(in)

	fmt.Println("Fanned out to 3 workers.")

	// 3. Fan-in: Merge the channels back into a single result channel.
	results := merge(worker1, worker2, worker3)

	fmt.Println("Merging results (Fan-in)...")

	// Read and print all merged results
	for res := range results {
		fmt.Printf("Result: %d\n", res)
	}

	fmt.Println("Finished Fan-out / Fan-in.")
}
