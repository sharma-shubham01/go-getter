package main

import (
	"fmt"
	"sync"
	"time"
)

// a worker pool is a concurrency pattern where you spawn a fixed number of worker goroutines, and they pull tasks from a job queue (channel)
// How it Works
// 1. Job Queue: A channel that holds incoming tasks (jobs).
// 2. Workers: Fixed number of goroutines that continuously read from the job channel, process jobs, and send results to a result channel.
// 3. Results: Another channel where workers send outputs.
// 4. WaitGroup: Ensures we wait until all jobs are processed.

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobId  int
	Output int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs { // reading the jobs from read only channel
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)

		time.Sleep(time.Second)
		results <- Result{JobId: job.ID, Output: job.Value * 2} // storing the results in write only channel

		fmt.Printf("Worker %d finished processing job %d\n", id, job.ID)
	}
}

func main() {
	var (
		numWorkers = 3
		numJobs    = 5

		jobs    = make(chan Job, numJobs)
		results = make(chan Result, numJobs)

		wg sync.WaitGroup
	)

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j}
	}

	close(jobs)

	wg.Wait()

	close(results)

	for res := range results {
		fmt.Printf("Job %d result: %d\n", res.JobId, res.Output)
	}
}
