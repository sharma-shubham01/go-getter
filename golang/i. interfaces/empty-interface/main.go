package main

import "fmt"

// EMPTY INTERFACES
// Empty Interface has no type or value
// The pre-declared identifier "any" is a type alias of an empty interface - interface{}
// It is used when you want to store any kind of value

// USES
// Empty Interfaces in Go are used when type is not known beforehand
// Example - Reading through a JSON file

func main() {
	mixedSlice := []interface{}{42, "Hello", 3.14, true}

	for _, value := range mixedSlice {
		printValue(value)
	}

	typeAssertion()
	typeSwitch()
}

func printValue(value any) {
	fmt.Println(value) // 42, "Hello", 3.14, true
}

// There are times when you need to derive back the type from an empty interface as go is a statically typed language

// TYPE ASSERTION
// Type assertion is used to assert type for an empty interface
// It is different from type conversion and can panic
// Type conversion is done at the compile time whereas type assertion is done at the runtime
// Since, type assertion is done at runtime, the code can panic if an assertion is wrong
// Type Conversion - string(typeDerivedFromString)
// Type Assertion - emptyInterface.(string)

func typeAssertion() {
	var emptyInterface any
	emptyInterface = "Hello World!"

	if str, ok := emptyInterface.(string); ok {
		fmt.Println("The underlying value is a string: ", str) // str -> Hello World!
	}

	if val, ok := emptyInterface.(int); !ok {
		fmt.Println("The underlying value is not string", val) // val -> 0
	}

	// The below piece of code will panic!
	// val := emptyInterface.(bool)
	// fmt.Println(val)
}

// TYPE SWITCH
// Type Switch can be used to check against multiple types
// Type assertion is done within the switch statement instead of switching on the value
// Inside each case, the value will of that type if case matches, otherwise remains empty interface type

func typeSwitch() {
	checkType(4)
	checkType("Hello")
	checkType(3.14)
	checkType(true)
}

func checkType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Type is an int: ", v)
	case float64:
		fmt.Println("Type is an int: ", v)
	case string:
		fmt.Println("Type is an int: ", v)
	default:
		fmt.Println("Unknown Type")
	}
}
