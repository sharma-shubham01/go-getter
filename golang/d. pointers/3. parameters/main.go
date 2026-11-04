package main

import "fmt"

// Call By Value
// 1. Call By Value - Pass the value of the parameter
// 2. Inside function, a copy of parameter is created
// 3. Any change made inside a function, do not have any effect on the actual value
// 4. In many languages, call by value is used for most primitive data types

// Call By Reference
// 1. Call By Reference - Pass the memory address of the parameter
// 2. Inside the function, you have the access to the actual memory where data is stored
// 3. Any change made inside the function are visible on the actual variables
// 4. In most languages, call by reference is used for complex data structures like classes and arrays

// Go - Strictly Call By Value
// 1. Even when you pass a pointer inside a function, a copy of pointer is created
// 2. Maps and slices are curious cases in golang
// 3. Maps and slices are internally implemented as data structures that hold memory address

func main() {
	a := 10
	aPtr := &a

	increament(aPtr)
	fmt.Println(aPtr) // 0x14000110018

	s := []int{1, 2}
	fmt.Println(len(s)) // 2

	somethingSlice(s)
	fmt.Println(len(s)) // 2
}

func increament(x *int) {
	// We are working with a copy of pointer
	// Hence, it will not reflect in the main function
	// even though we are updating the pointer to nil
	x = nil
}

// Maps
// 1. Maps are part of the reference types in golang
// 2. Data Structure that holds the maps has a reference of memory where data is stored
// 3. When passing to function, a copy of the structure defining a map is created
// 4. The copy also includes the reference to memory when data is stored
// 5. Updating the map inside a function, will update the original map

// Slices
// 1. Slices are also reference data types along with maps and pointers
// 2. A slice is a data structure that holds length, capacity and pointer to the memory where data is stored
// 3. Inside a function, again copy of slice structure is created
// 4. Directly changing the value will update the original slice, but append is different
// 5. If we append anything beyond the length of an original slice, it won't show up in the original slice
// as the slice field length still hasn't been updated since we are working with the copy inside a function

func somethingSlice(s []int) {
	s = append(s, 1000)
	fmt.Println(len(s)) // 3 // we are working with a copy
}