package main

import (
	"fmt"
	"time"
)

// The Worker Pool pattern is one of the most common concurrency patterns in Go.
// It limits the number of goroutines running concurrently to process a queue of jobs.
// This is crucial for resource management (e.g., limiting database connections or API requests).

// Job represents a unit of work.
type Job struct {
	ID    int
	Value int
}

// Result represents the output of a processed Job.
type Result struct {
	Job    Job
	Output int
	Worker int
}

// worker is a single worker execution block.
// - id: identifies the worker
// - jobs: receive-only channel for tasks (<-chan)
// - results: send-only channel for processed tasks (chan<-)
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d: Started processing job %d\n", id, job.ID)
		
		// Simulate computation time
		time.Sleep(50 * time.Millisecond)
		
		output := job.Value * 2
		
		fmt.Printf("Worker %d: Finished job %d\n", id, job.ID)
		results <- Result{Job: job, Output: output, Worker: id}
	}
}

func RunWorkerPoolDemo() {
	fmt.Println("=== Worker Pool Pattern ===")

	numJobs := 10
	numWorkers := 3

	// Create channels for jobs and results
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// 1. Spawn workers. They immediately block waiting for jobs.
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	fmt.Printf("Spawned %d worker goroutines.\n", numWorkers)

	// 2. Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j * 10}
	}
	fmt.Println("Sent all jobs to workers.")
	
	// Close the jobs channel to signal to the workers that no more jobs are coming.
	// This causes their 'for range jobs' loops to terminate.
	close(jobs)

	// 3. Collect results.
	// Since we know exactly how many jobs we sent, we can receive that exact number of results.
	for r := 1; r <= numJobs; r++ {
		res := <-results
		fmt.Printf("Result collected: Job %d processed by Worker %d -> Output = %d\n", 
			res.Job.ID, res.Worker, res.Output)
	}

	fmt.Println("All jobs completed and results collected.")
}

func main() {
	RunWorkerPoolDemo()
}
