package main

import (
	"fmt"
	"sync"
)

//Create a fixed number of worker goroutines
//Distribute the work of squaring each integer among the workers using channels.
//Collect the results and print them in any order

func worker(id int, input chan int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range input {
		result <- i * i

		fmt.Printf("Worker Pool %d has finished processing\n", id)
	}

}

func main() {
	var (
		num        = 10
		numWorkers = 10
		input      = make(chan int, num)
		result     = make(chan int, num)
		wg         sync.WaitGroup
	)

	for i := 1; i <= num; i++ {
		input <- i
	}

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, input, result, &wg)
	}

	close(input)

	wg.Wait()

	close(result)

	for r := range result {
		fmt.Println(r)
	}
}
