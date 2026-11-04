package main

import "fmt"

// When a function panics, normal flow stops and the call stack begins unwinding, but all deferred functions are still executed.

// Rule: recover() is only effective if called directly inside a deferred function.
// Calling recover() inside normal code or inside a nested sub-function called by defer returns nil and will not halt the panic.

func safeDivide() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	panic("division by zero")
}
