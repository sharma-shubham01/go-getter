package main

import "fmt"

// Inheritance
// 1. In many programming languages, you create a new class by deriving it from another class
// 2. Inheritance allows the derived class and the base class/parent class to be used interchangeably
// 3. The methods defined on the parent class can be used by the child class

// Inheritance in GO - not supported
// 1. In go, the parent and the child types are two distinct types and can't be used interchangeably
// 2. Here, the parent and child types do not share any methods, they both have their own method sets
// 3. Go is an object-oriented language, but it doesn't use inheritance, rather it relies on composition for code reuse

type Age int

func sum(a Age) {
	fmt.Println(a)
}

func main() {
	//var young int = 10

	//sum(young) // Cannot use 'young' (type int) as the type Age

	// Age and int are different types

	sum(10)
}
