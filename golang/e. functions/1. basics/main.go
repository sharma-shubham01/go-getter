package main

import "fmt"

// Functions
// 1. Functions are modular, reusable blocks of code designed to accomplish specific tasks
// 2. Functions take input(optional), process it and returns output(optional)
// 3. Functions enable us to structure and organise the code
// 4. A program is basically an orchestration of function calls
// 5. Functions in Go can return multiple values

// Void Function
// 1. Functions that don't return any values are called void functions
// 2. In go, the entry point function main takes no parameters and returns no values

func main() {
	sayHi()
	sayHiToSomeone("Shubham")

	fullName := printFullName("shubham", "sharma")
	sayHiToSomeone(fullName)

	fn, l := printFullNameWithLength("shubham", "sharma")
	fmt.Println(fn, l)
}

// void function
func sayHi() {
	fmt.Println("Hi!")
}

// accepts single parameter
func sayHiToSomeone(name string) {
	fmt.Println("Hi! ", name)
}

// accepts multiple parameters and return single parameter
func printFullName(first, last string) string {
	return fmt.Sprintf("%s %s", first, last)
}

// accepts multiple parameters and return multiple parameters
func printFullNameWithLength(first, last string) (string, int) {
	fn := fmt.Sprintf("%s %s", first, last)
	return fn, len(fn)
}
