package main

import (
	"fmt"
	"sync"
)

//You need to write two functions - one that prints the first 10 odd numbers and the other prints the first 10 even numbers.
//They should print them in the sequence like 0 1 2 3 4 5 .... and so on. It should be synchronous

// Algo
// 1. Create unbuffered channels to provide signal whether to print odd number or an even number
// 2. Whenever even channel has a signal, print even and send signal for odd
// 3. Whenever odd channel has a signal, print odd and send signal for even

func printEven(evenChan, oddChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		<-evenChan
		fmt.Print(i*2, " ")
		oddChan <- true
	}
}

func printOdd(evenChan, oddChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		<-oddChan
		fmt.Print(i*2+1, " ")

		// 0-8: The "odd" goroutine signals the "even" goroutine to start the next iteration
		// on the final iteration, the "odd" goroutine knows there is no "next round," so it skips the signal and finishes cleanly.
		if i < 9 {
			evenChan <- true
		}
	}
}

func main() {
	evenChan := make(chan bool)
	oddChan := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)

	go printEven(evenChan, oddChan, &wg)
	go printOdd(evenChan, oddChan, &wg)

	evenChan <- true
	wg.Wait()
	fmt.Println()
}

// Alternative approach - using select

func usingSelect() {
	var (
		oddChan  = make(chan struct{}, 1)
		evenChan = make(chan struct{}, 1)

		oddCount  = 0
		evenCount = 0
	)

	// Seed the loop with the even turn
	evenChan <- struct{}{}

	for oddCount < 10 || evenCount < 10 {
		select {
		case <-evenChan:
			fmt.Print(evenCount*2, " ")
			evenCount++
			oddChan <- struct{}{}

		case <-oddChan:
			fmt.Print(oddCount*2+1, " ")
			oddCount++
			// Pass turn back as long as more even numbers remain
			if evenCount < 10 {
				evenChan <- struct{}{}
			}
		}
	}
	fmt.Println()
}
