package main

import "fmt"

// Deadlocks and Blocking
// occur when using channels if goroutines become blocked indefinitely. Few scenarios -
// if a goroutine is waiting to send a value on an unbuffered channel, but there is no receiver to receive the value
// if a goroutine is waiting to receive a value from an unbuffered channel, but there is no sender to send the value
// if a goroutine is waiting to send a value on a buffered channel that is already full, but there is no receiver to receive the value
// How to avoid?

// 1. Close Channels Properly
// closing a channel is important to signal to the receiver that all the expected values have been sent
// a closed channel can still be received from, but it will always return zero values after the last value has been received

func main() {
	ch := make(chan int) // Creating an unbuffered channel

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // Sending values to the channel
		}

		close(ch) // Closing the channel after all values are sent
	}()

	for x := range ch { // iterating over a channel
		fmt.Println(x)
	}
}

// 2. Synchronisation (Using Wait Groups)
