package main

import "fmt"

// arrays are data structures used to store data in contiguous block of memory
// each piece of data of an array is called the element of an array
// each element of an array needs to be of the same type
// we use index of the element in the array to access the element
// arrays are zero indexed

// Benefits of arrays
// 1. arrays are very efficient to fetch elements from memory
// 2. used to store collection of data, such as list of numbers etc
// 3. uses include search and sort algorithms, buffers etc

// Declaring an array

// 1. Array Declaration Basic Form
// declares an array of size 5
var a [5]int

// 2. Array Declaration with Elements
// declares an array of size 5 and assign values to it
var a1 = [5]int{1, 2, 3, 4, 5}

// 3. Sparse Array Declaration
// declares an array of size 5
// assign value 1 at 0th index, value 10 at 2nd index and then value 50
// by default, the values will be zero at remaining indices
var a2 = [5]int{1, 2: 10, 50}

// 4. Implicit Length Declaration
// declare an array with ellipse notation
// it will determine the size of an array based on the values assigned to it
var a3 = [...]int{1, 4, 7}
var a4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 2D ARRAYS
var a5 = [2][2]int{{1, 2}, {3, 4}}

// len() function is used to calculate the length of an array
func main() {
	fmt.Println("a -", a)   // {0, 0, 0, 0, 0}
	fmt.Println("a1 -", a1) // {1, 2, 3, 4, 5}
	fmt.Println("a2 -", a2) // {1, 0, 10, 50, 0}
	fmt.Println(len(a3))    // 3
	fmt.Println(len(a4))    // 10
	fmt.Println("2D Array -", a5)

	// Accessing Array Elements
	fmt.Println("First Element of a4 -", a4[0]) // 1
	fmt.Println("Last Element of a4 -", a4[9])  // 10

	// Manipulate Array Elements
	a4[1] = 20
	fmt.Println(a4) // {1, 20, 3, 4, 5, 6, 7, 8, 9, 10}
}

// Limitations

// 1. once declared, the size cannot be altered
// 2. the size or length of the array are stored in the type of the array
// 3. you cannot use variable to declare the size of an array, it should be constant
// 4. most cases slices are favoured over arrays
