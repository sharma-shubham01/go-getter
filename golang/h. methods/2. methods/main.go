package main

import "fmt"

// Methods and Functions
// 1. Methods are attached to a type while functions are standalone
// 2. Methods can be used for functions if the signature of method is same as of function type
// 3. Similar to functions, methods can be saved in variables and passed around

type Calculator struct {
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func ArithmeticOperation(f func(a, b int) int, a, b int) int {
	return f(a, b)
}

func main() {
	// methods can be used for functions if the signature of method is same as of function type
	c := Calculator{}
	fmt.Println(ArithmeticOperation(c.Add, 10, 15))

	// similar to functions, methods can be saved in variables and passed around
	printer := Printer{}
	printFunction := printer.Print
	printFunction("Hello World!")
}

type Printer struct {
}

func (p *Printer) Print(m string) {
	fmt.Println(m)
}
