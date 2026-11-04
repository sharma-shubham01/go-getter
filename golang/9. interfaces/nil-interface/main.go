package main

import "fmt"

// NIL Interface
// The zero value for an interface type is nil
// A nil value for a concrete type that the variable has no value or the value is absent
// For an interface to be a nil, both the type and value should be nil
// This is because, in Go interfaces are implemented as a pair of pointers one to the type and one to the value

func main() {
	var a *int
	var i interface{}

	fmt.Println(a == nil) // true
	fmt.Println(i == nil) // true

	i = a
	fmt.Println(i == nil) // false

	// why false?
	// variable - i was assigned a type which is an int pointer
	// interfaces have two pointers - one to the type and one to the value
	// for an interface to be nil, both the values should be nil
	// but here since the interface has a type assigned to it
}
