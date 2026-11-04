package main

import (
	"fmt"
	"time"
)

// channels
// allow goroutines to communicate and synchronize their execution

// unbuffered channel - synchronous communication
// an unbuffered channel has a capacity of zero
// every send operation on the channel blocks until another goroutine is ready to receive the value
// every receive operation blocks until another goroutine is ready to send a value
func executeUnbufferedChannel() {
	ch := make(chan int) // created an unbuffered channel

	go func() {
		time.Sleep(time.Second)

		ch <- 5 // sending value to a channel
		// this will be a blocking call until n unless some receiver is ready
	}()

	x := <-ch // read value from channel and store in x

	fmt.Println(x)
}

// buffered channel - asynchronous communication
// a buffered channel has a specified capacity greater than zero
// they can hold a certain number of values before blocking send operations
// they decouple the sender and receiver, allowing for asynchronous communication
func executeBufferedChannel() {
	ch := make(chan int, 2) // created a buffered channel with capacity as 2

	ch <- 1 // sending value to a channel
	ch <- 2 // sending another value to a channel

	x := <-ch // reading first value from a channel
	fmt.Println(x)

	y := <-ch // reading another value from a channel
	fmt.Println(y)
}

// unidirectional channel
// they restrict the direction of data flow, allowing only sending or receiving operations
func executeUnidirectionalChannel() {
	ch := make(chan<- int)  // created an unbuffered send only channel
	ch1 := make(<-chan int) // created an unbuffered read only channel

	go executeSendOnlyChannel(ch)
	go executeReceiveOnlyChannel(ch1)
}

// goroutines that have access to a send-only channel can send values to the channel, but they cannot receive or read from it
func executeSendOnlyChannel(ch chan<- int) {
	ch <- 5 // sending value to a channel
}

// goroutines that have access to a receive-only channel can read values from the channel, but they cannot send or write to it
func executeReceiveOnlyChannel(ch <-chan int) {
	x := <-ch

	fmt.Println(x)
}

// bidirectional channel
// allow both sending and receiving operations, enabling data flow in both directions
func executeBidirectionalChannel() {
	ch := make(chan int) // created a bidirectional buffered channel

	go executeSendOnlyChannel(ch)
	go executeReceiveOnlyChannel(ch)

}

func main() {
	executeUnbufferedChannel()
	executeBufferedChannel()
	executeUnidirectionalChannel()
	executeBidirectionalChannel()
}
