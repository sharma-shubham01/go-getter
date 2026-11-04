package main

import "fmt"

// slices are wrappers around arrays
// they manage the underlying arrays for you,
// they help you overcome the limitations of an array
// length is not attached to the type of slice
// slices have length and capacity

// Length of a slice is the length of elements stored in a slice (use len() function)
// Capacity is the total size of the underlying array (use cap() function )
// usually, when we create a slice, it will always provision an array of greater capacity then required
// the additional capacity is used to store elements that you add to the slice so that it doesn't have to create a new array each time

// Behind the scenes
// when you create a slice of certain length, underlying array is usually created of double the length
// when you reach the capacity of a slice, it will create an array of 50 to 25% larger than the existing array
// and it copies the values from old array to new array
// which can be little problematic as well if the slice is too big

func main() {
	// Declaring slices in golang
	var arr []int
	fmt.Println(arr)

	var arr1 = []int{1, 2, 3, 4}
	fmt.Println(arr1)

	var arr2 = []int{1, 2: 10, 50}
	fmt.Println(arr2)

	var arr3 = make([]int, 0)
	fmt.Println(arr3)

	// make() function
	// the universal make function can be used to create slices
	// it takes type of slice as a first argument, length as second and capacity as the optional third argument

	var a1 = make([]string, 5)
	fmt.Println(len(a1), cap(a1)) // 5 5

	var a2 = make([]string, 5, 10)
	fmt.Println(len(a2), cap(a2)) // 5 10

	var a3 = make([]int, 0, 10)
	fmt.Println(len(a3), cap(a3)) // 0 10

	// append() function
	// the append universal function can be used to add elements to a slice
	// it is an error to not assign the return value of append
	a1 = append(a1, "hello")
	a1 = append(a1, "world", "how", "are", "you")
	fmt.Println(a1)

	// spread operator - used to append slices
	a2 = append(a2, a1...)
	fmt.Println(a2)

	// Nil and Empty Slices
	// nil in Go is different from null of other languages
	// nil means a lack of value and type
	// nil is the zero value of a slice
	// a nil slice has no underlying array
	// an empty slice has underlying array with no values stored
	var (
		nilSlice   []int
		emptySlice = []int{}
	)

	fmt.Println(nilSlice == nil)   // true // a nil slice with no underlying array
	fmt.Println(emptySlice == nil) // false // an empty slice with an underlying array but no values
}
