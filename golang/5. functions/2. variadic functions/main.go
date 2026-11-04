package main

import "fmt"

// Variadic Functions
// 1. Go permits the passing of variadic parameters to a function
// 2. The number of parameters you pass to a function is not fixed for variadic parameter function
// 3. Variadic parameters, in essence, allow for passing of any number of function arguments

// Rules - variadic parameters
// 1. Syntax - Empty three dots ... (must be placed before the type of the parameter)
// 2. It must be the final parameter - because it captures all remaining arguments
// 3. Within the function, these parameters are accessible as a slice

// Syntax
// func functionName(paramName ...paramType) returnType { }

// Example -
// fmt.Println(a ...any)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	s := sum(nums...) // nums... - separate each value of the slice into individual values
	fmt.Println(s)
}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// Default Parameters
// 1. Go lacks named or default parameters, a feature found in some other languages
// 2. Aim to limit the number of parameters in a function to just a couple for clarity and simplicity
