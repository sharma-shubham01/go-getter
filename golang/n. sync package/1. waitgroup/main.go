package main

import "sync"

// sync.WaitGroup waits for a collection of goroutines to finish executing. It maintains an internal counter.

// Core Mechanics
// Add(delta int): Increments (or decrements) the counter.
// Done(): Decrements the counter by 1 (equivalent to Add(-1)).
// Wait(): Blocks the caller until the counter drops to 0.

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // Always call Add before launching the goroutine
		go func(id int) {
			defer wg.Done()
			// do work
		}(i)
	}

	wg.Wait()
}

// Few points -
// 1. Calling Add() inside the goroutine: Causes a race condition where Wait() might return before the goroutines start and invoke Add().

// 2. Passing by value: WaitGroup contains a noCopy struct.
// If passed by value into functions or goroutines, a copy is made, which creates a detached counter and leads to deadlocks.
// Pass it by reference (*sync.WaitGroup).

// 3. Negative counter: Calling Done() more times than Add() panics (negative WaitGroup counter).
