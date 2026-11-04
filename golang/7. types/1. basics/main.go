package main

import "fmt"

// Types
// 1. A user-defined types can be created with the type keyword
// 2. Derive a type from the existing primitive or composite types
// 3. Examples -
// type toDoList []string
// type student struct {...}

// Abstract Types
// 1. Abstract type defined what a type should do
// 2. Example - function type

// Concrete Types
// 1. Concrete type defines what and how the type should do it
// 2. Example - struct with methods

type Person struct {
	Name string
	Age  int
}

func (p Person) GetDetails() string {
	return fmt.Sprintf("Name: %s, Age: %d ", p.Name, p.Age)
}

func main() {
	p := Person{
		Name: "Shubham",
		Age:  25,
	}

	fmt.Println(p.GetDetails())

	derivedTypes()
}

// Derived Types
// 1. We can derive types from other built in types
// 2. We can also derive types from user defined types

type Age int

func derivedTypes() {
	var young Age = 10
	var old Age = 60

	fmt.Println(young + old) // 70

}
