package main

import "fmt"

// Fist Class Functions
// 1. In Go, functions are first class citizens
// 2. First class functions imply that functions can be treated as values
// a. They are assignable to variables
// b. They can be passed as arguments to other functions
// c. They can also be returned from the functions
// 3. This provides great flexibility while writing code

func sayHello(first, last string, fn func(string, string) string) string {
	fullName := fn(first, last)

	return fmt.Sprintf("Welcome %s", fullName)
}

func main() {
	fullName := func(s1, s2 string) string { // function value in a variable
		return fmt.Sprintf("%s %s", s1, s2)
	}

	fmt.Println(sayHello("Shubham", "Sharma", fullName))

	// function passed as an argument
	x := 5
	res := apply(double, x)
	fmt.Println(res)

	res1 := apply(triple, x)
	fmt.Println(res1)

	multiplyByTwo := multiplyBy(2)
	result := multiplyByTwo(5) // 10
	fmt.Println(result)

	multiplyByThree := multiplyBy(3)
	result1 := multiplyByThree(7)
	fmt.Println(result1) // 21
}

func double(x int) int {
	return x * 2
}

func triple(x int) int {
	return x * 3
}

func apply(f func(int) int, x int) int {
	return f(x)
}

// A Function returned from a function
func multiplyBy(multiplier int) func(int) int {
	return func(i int) int {
		return i * multiplier
	}
}
