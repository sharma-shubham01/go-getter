package main

import "fmt"

// Anonymous Functions
// 1. The functions without name are known as anonymous functions
// 2. Anonymous Functions are useful for defining a small piece of functionality inline without the need to give it a name
// 3. Anonymous Functions are also used to provide callbacks

func main() {
	func() {
		fmt.Println("This is an Anonymous Functions")
	}()

	func(msg string) {
		fmt.Println("Message: ", msg)
	}("Hello")

	numbers := []int{1, 2, 3, 4, 5}
	squared := sliceOperation(numbers, func(i int) int {
		return i * i
	})
	fmt.Println(squared) // [1, 4, 6, 16, 25]
}

func sliceOperation(s []int, op func(int) int) []int {
	result := make([]int, len(s))

	for i, v := range s {
		result[i] = op(v)
	}

	return result
}
