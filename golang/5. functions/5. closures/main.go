package main

import "fmt"

// Closures
// 1. Closures are functions that close over values
// 2. It captures the values from its surrounding scope usually a function block
// 3. Closure keeps access to the closed over variables ever after the surrounding scope has been executed
// 4. Closure can also be seen as function with state associated

// Benefits
// 1. Closures can access and use variables that are not directly passed as arguments by closing over them
// 2. They can operate upon captured variables independently
// 3. Closures are used for creating memoization functions, implementing callbacks etc.

func main() {
	message := "Hello, "

	greet := func(name string) {
		fmt.Println(message + name)
	}

	greet("Alice") // Hello, Alice // message is a closure or inside another function, we're able to capture message string
	greet("Bob")   // Hello, Bob   // without passing it as an argument

	increament := counter()
	fmt.Println(increament()) // 1
	fmt.Println(increament()) // 2
	fmt.Println(increament()) // 3 // because it was able to capture and maintain its own state of count variable
}

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
