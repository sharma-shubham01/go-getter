package main

import (
	"fmt"
	"sync"
)

// What is Concurrency?
// 1. Dividing and executing programs in parts
// 2. Concurrency is not parallelism
// 3. Although concurrency may give an illusion of parallelism
// 4. Ex - Maybe the parts of program are sharing their time on single cpu but to human eye it may look like they're running in parallel

// Need of Concurrency
// 1. If implemented correctly, it can increase the performance of a program
// 2. Concurrency is evolution to extract more performance from computers
// 3. Meant to exploit the multiple CPUs now available

// The Challenges
// 1. Concurrency comes with inherent challenges like race conditions
// 2. When two or more concurrent stream of a program try to access a shared resource.
// 3. Race condition cause problems as computer programs needs to be deterministic

// Atomicity
// 1. Atomicity refers to an operation that is indivisible and irreducible
// 2. Database transaction are one of the best examples
// 3. Either they will happen completely or they will not happen at all
// 4. By employing atomic operations, we can ensure that the critical sections of code are executed as a single unit

// Critical Section
// 1. Accessing and modifying the shared data is done in critical sections
// 2. Identify critical sections of our code is very important

func main() {
	var (
		lock    sync.Mutex
		counter int
	)

	for i := 1; i <= 1000; i++ {
		go func() {
			lock.Lock()   // a go routine will apply a lock and then update the counter to avoid race condition
			counter++     // critical section of the code
			lock.Unlock() // after updating the counter, it will release the lock
		}()
	}

	lock.Lock()
	fmt.Println("Expected Counter: ", 1000)  // 1000
	fmt.Println("Actual Counter: ", counter) // some random values like 939, 987, 956 etc...
	lock.Unlock()

	// Open Issue -
	// 1. The above code is not deterministic as we are getting random values for counter
	// 2. Reason - The main thread is not waiting properly for all the go routines to get completed
}

// Go Routines and Channels
// 1. Goroutine are lightweight threads managed by Go run time
// 2. We don't have to worry about low level thread management
// 3. Channels acts as a communication pipeline between go routines
// 4. Channels help in data flow enabling go routine exchange information without complex sync techniques
// 5. Go runtime which is embedded in the binary that we build takes care of scheduling go routines into available CPU cores
// 6. If a go routine is blocked, let's say waiting for I/O, the runtime can switch to another go routine improving overall app
// 7. Hence, it's an added layer of encapsulation provided by go to write concurrent programs efficiently
