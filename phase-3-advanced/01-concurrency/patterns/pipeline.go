package main

import "fmt"

// A Pipeline is a series of stages connected by channels, where each stage
// is a group of goroutines running the same function.
//
// In each stage, the goroutine:
// 1. Receives values from inbound channels.
// 2. Performs an operation on that data (usually producing new values).
// 3. Sends values outbound on channels.
//
// By using pipelines, we decouple execution stages and can process data in a streaming fashion.

// Stage 1: Generator
// Converts a variadic list of integers into a channel that emits them.
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Stage 2: Doubler
// Receives integers from in, doubles them, and sends them to out.
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// Stage 3: Stringify
// Receives integers, formats them as strings, and sends them to out.
func stringify(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			out <- fmt.Sprintf("Number: %d", n)
		}
		close(out)
	}()
	return out
}

func RunPipelineDemo() {
	fmt.Println("=== Pipeline Pattern ===")

	// Setup the pipeline stages:
	// generate -> double -> stringify
	intsChan := generate(1, 2, 3, 4, 5)
	doubledChan := double(intsChan)
	stringifiedChan := stringify(doubledChan)

	// Consume the final output of the pipeline:
	for result := range stringifiedChan {
		fmt.Println(result)
	}

	fmt.Println("Finished Pipeline.")
}

func main() {
	RunPipelineDemo()
}
