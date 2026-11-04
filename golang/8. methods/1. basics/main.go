package main

import "fmt"

// Methods
// 1. Methods are functions that are associated and belong to a type
// 2. You can define methods for both custom types and derived types
// 3. Syntax is slightly different from functions
// 4. We define methods on type in golang as shown below

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// More about methods
// 1. There are no overloaded methods in go
// 2. Each method in go must have a unique name and the method signature including the receiver type
// 3. Methods can be accessed just like field i.e using dot(.) notation

type Rectangle struct {
	Length int
	Width  int
}

// Here the receiver is a value receiver
// Useful when we want to access a field on the type
// If we pass the value type a copy will be created and the difference won't be seen on the original type
func (r Rectangle) area() int {
	return r.Length * r.Width
}

// If the method is set to change the fields or modify the instance of that type
// then we need to have a pointer receiver
func (r *Rectangle) setLength(l int) {
	r.Length = l
}

func main() {
	r := Rectangle{
		Length: 5,
		Width:  4,
	}

	fmt.Println(r.area()) // 20

	r.setLength(10)
	fmt.Println(r.area()) // 40
}
