package main

import (
	"fmt"
	"sync"
)

// Write a Go program to process 50 jobs concurrently using goroutines.
// Make sure that at most 5 goroutines execute the job-processing section simultaneously.
// Each job should calculate the cube of an integer and send the result to an output channel.

func worker(id int, input, output chan int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	for i := range input {
		sem <- struct{}{}

		output <- i * i * i

		fmt.Printf("Processing completed by worker: %d\n", id)

		<-sem
	}
}

func main() {
	var (
		numJobs    = 50
		numWorkers = 50

		input  = make(chan int, 50)
		output = make(chan int, 50)

		wg sync.WaitGroup

		limit = 5
		sem   = make(chan struct{}, limit)
	)

	for i := 1; i <= numJobs; i++ {
		input <- i
	}

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go worker(i, input, output, &wg, sem)
	}

	close(input)

	wg.Wait()

	close(output)

	for res := range output {
		fmt.Printf("Result: %d\n", res)
	}
}
