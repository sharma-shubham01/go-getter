package main

import "fmt"

// CONCRETE DATA TYPES
// The concrete data types specify the functionality as well as the implementation
// It describes how the type should do it
// Example - Struct and Methods

// ABSTRACT DATA TYPES
// The abstract data types only specify the functionality
// It describes what the type should do but not how it should do it
// Example - Interfaces

// INTERFACES
// Interfaces are abstract types that specify what a type should do or the behaviour
// To write an interface, use the type and interface keyword
// Inside the interface, we declare the signature of methods, known as method set of an interface

// IMPLEMENT AN INTERFACE
// To implement the interface, a type must implement all the methods in the interface
// Names of interfaces from standard lib usually ends with ...er, example - io.Writer
// Interfaces in go are implicit which is different from other languages
// In other languages, you have to specify if your type is implementing an interface in type declaration
// In Go, you don't have to do that
// If types implements the methods listed in the interface, it implements the interface
// A type can implement multiple interfaces

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Type Rectangle implements an interface Shape as it implements Area()

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Type Circle implements an interface Shape as it implements Area()

func main() {
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	shapes := []Shape{rectangle, circle}

	for _, shape := range shapes {
		printArea(shape)
	}

	embeddingInterfaces()
}

func printArea(s Shape) {
	fmt.Println("Area - ", s.Area())
}

// USES OF INTERFACES
// Interfaces are very practical as your code works with functionality and not how it is implemented
// Database interface implements methods but does not matter if it uses Postgres, MySQL etc
// This makes up for a maintainable code
// Decouples implementation from abstraction
// Enables Polymorphism

// EMBEDDING WITH INTERFACES

// Object interface has now two functions - Area() and Name()
// For any type, to implement an object interface, it should implement both the methods
type Object interface {
	Name() string
	Shape
}

func (r Rectangle) Name() string {
	return "Rectangle"
}

func printObject(o Object) {
	fmt.Printf("Area: %f, Name: %s", o.Area(), o.Name())
}

func embeddingInterfaces() {
	rectangle := Rectangle{Width: 5, Height: 3}

	printObject(rectangle)
}
