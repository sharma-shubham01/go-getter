package main

import "fmt"

// USES

// MUTABILITY
// 1. Variables can be categorised as either mutable or immutable
// 2. Mutable variables allow their values to be changes
// 3. Immutable variables prohibit any changes to their value
// 4. In Go, within a function, we can enforce immutability with a const and mutability with a var
// 5. Pointers are employed when functions need to alter the state of passed variables
// 6. With the address of the variable, a function can directly update its value

// INDICATION OF NO VALUE
// 1. Occasionally, it's important for the calling function to distinguish between a zero value and no value at all
// 2. When a function returns a pointer, it can use nil to signal absence of a value
// 3. This pattern is widely used in GO, especially when a function returns both a value and a potential error

// PERFORMANCE BENEFITS
// 1. Passing large structs as arguments in functions may lead to performance concerns because entire struct needs to
// be copied over to the function which can be resource intensive
// 2. Instead of passing an entire struct, you can pass a pointer to the struct
// 3. Pointers are allocated in a heap which has a longer lifespan, but we need to define pointer when there is
// an actual need, as optimising memory usage to minimize garbage collector invocations enhance program performance

type student struct {
	firstName string
	lastName  string
}

func main() {
	a := 10

	increament(&a)
	fmt.Println(a)

	s := student{
		firstName: "Shubham",
		lastName:  "Sharma",
	}
	updateLastName(&s, "Kaushik")
	fmt.Println(s)

}

func increament(x *int) {
	*x++
}

func updateLastName(s *student, newLastName string) {
	s.lastName = newLastName // no need to de-reference the struct fields, go will handle it
}
