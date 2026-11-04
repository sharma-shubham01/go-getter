package main

import "fmt"

// DEFER
// 1. In go, the defer statement is used to schedule a function call to be executed later after the surrounding
// function or the function containing the defer statement completes its execution
// 2. This ensures that the deferred function is called regardless of how the surrounding function finishes

// CONDITIONS for DEFERRED FUNCTION EXECUTION
// 1. surrounding function returns
// 2. code execution reaches the end of function body
// 3. function panics

// IMPORTANT POINTS
// 1. The deferred call's arguments are evaluated immediately, but the function call is not executed right away
// 2. Multiple deferred functions are executed in last in first out manner (LIFO)
// means the function that was deferred last will be executed first
// 3. Deferred functions are executed after any result parameters are set by that return statement but before the function return to its caller
// 4. Used for cleanup like closing the file or releasing the resources and panics

func main() {
	message := "hello"

	greetingFn := func(name string) {
		fmt.Println(message + name)
	}

	defer greetingFn("Alice")
	defer greetingFn("Bob")

	message = "Hi, "
}
