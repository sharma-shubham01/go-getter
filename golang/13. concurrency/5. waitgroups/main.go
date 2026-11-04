package main

import (
	"errors"
	"fmt"
	"sync"
)

// wait groups are a mechanism for managing concurrency in Go
// they provide a way to wait for a group of goroutines to complete their execution before proceeding further

// Add() is used to add the number of goroutines that need to be waited upon
// Done() is called by each goroutine when it finishes its work, decrementing the internal counter of the wait group
// Wait() is used to block the execution of the goroutine until all the goroutines have called Done()

func worker(id int, wg *sync.WaitGroup, err error) {
	defer wg.Done() // will keep on decrementing the counter as the processing is done for a go routine

	if id == 2 {
		err = errors.New("something unexpected occurred")
		return
	}

	fmt.Println("Worker", id, "completed")
}

func main() {
	var (
		numWorkers = 3
		wg         sync.WaitGroup
		err        error
	)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // will keep on incrementing the counter for wait group to wait

		go worker(i, &wg, err) // start processing in a separate thread (go routines)
	}

	wg.Wait() // blocks the main thread until wait group counter becomes zero

	if err != nil {
		fmt.Println("An error occurred", err)
	} else {
		fmt.Println("All workers completed!!")
	}
}
