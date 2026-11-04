package main

import "fmt"

// Memory
// 1. Programs are allocated memory by the operating system for execution
// 2. The allocated memory is divided into two key segments - stack and heap

// Stack
// 1. The stack is a region of memory that is used for execution of function calls & management of local variables
// 2. The stack is like a call history, it keeps track of function calls, local variables within the function calls
// 3. When function is called, its information is pushed onto the stack
// 4. When function call finishes, its information is popped off the stack

// Heap
// 1. The heap on the other hand is used for dynamic memory allocation which allows for creation and deletion of variables as needed
// 2. The heap however is a free storage area where you can create and delete variables dynamically
// 3. It allows for more flexible memory management but requires manual cleanup in most languages

// Garbage Collector in Go
// 1. The Go runtime on the other hand comes with a garbage collector,
// 2. which means that the Go runtime system automatically manages memory allocation and deallocation for us
// 3. It keeps track of which part of memory are in use and which are not
// 4. It periodically reclaims the memory that is not being used by the program

// Pointers
// 1. Initializing a value in a program allocates memory to store a value
// 2. A pointer is a variable that holds the memory address of another variable
// 3. Pointers of any data type has the same size
// 4. The zero value for a pointer variable is nil
// 5. Trying to access a value stored at nil pointer will give you runtime panic

// Pointer Operators
// 1. When the asterisk(*) is used as a prefix with a variable type, it signifies that it stores a pointer
// 2. It is also employed to retrieve the value stored at the address pointed to by a pointer, known as dereferencing
// 3. The ampersand(&) operator, when used as a prefix, provides the address of the variable
// 4. Pointers cannot be created for constants of primitive literals
// contants and primitive literals are created only at compile time, so they don't have any memory address

func main() {
	var intPtr *int

	age := 10
	intPtr = &age

	fmt.Println(intPtr) // stores memory address of age

	fmt.Println(*intPtr) // de-referencing // return the value stored at age variable
}
