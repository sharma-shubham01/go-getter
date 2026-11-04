package main

import "fmt"

// RULES TO REMEMBER

// Rule-1
// Function arguments are evaluated at the moment the defer line is encountered, not when the function actually executes.

func main1() {
	count := 1

	// Arguments are evaluated NOW: count is 1
	defer fmt.Println("Deferred count:", count)

	count = 100
	fmt.Println("Current count:", count)
}

// Output
// Current count: 100
// Deferred count: 1

// Rule-2
// A deferred function can read and modify the enclosing function's named return values before they reach the caller.

// In unnamedReturn(): return result copies the value 1 into an anonymous return slot.
// The deferred function runs and updates the local variable result to 11, but the return slot already holds 1.

// In namedReturn(): The variable result is the return slot.
// The deferred function modifies result directly, so the caller receives 11.

// Unnamed return: caller gets 1
func unnamedReturn() int {
	result := 1
	defer func() {
		result += 10 // modifies local variable, NOT the caller's return slot
	}()
	return result
}

// Named return: caller gets 11
func namedReturn() (result int) {
	result = 1
	defer func() {
		result += 10 // directly modifies the named return slot
	}()
	return result
}

func main2() {
	fmt.Println("unnamed:", unnamedReturn()) // 1
	fmt.Println("named:  ", namedReturn())   // 11
}
