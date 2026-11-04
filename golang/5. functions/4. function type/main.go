package main

import "fmt"

// Function Type
// Passing around function signatures can lead to code that is less readable
// Function signature can be utilised to define unique types

// Example -
// func (a, b int) int
// type op func (a, b int) int
// This enhances code readability

type operation func(int, int) int

func main() {
	var perform operation

	perform = arithmeticOperation("add")
	fmt.Println(perform(1, 2)) // 3

	perform = arithmeticOperation("subtract")
	fmt.Println(perform(7, 3)) // 4

	perform = arithmeticOperation("multiply")
	fmt.Println(perform(3, 4)) // 12
}

func arithmeticOperation(op string) operation {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "subtract":
		return func(a, b int) int {
			return a - b
		}
	default:
		return func(a, b int) int {
			return a * b
		}
	}
}
