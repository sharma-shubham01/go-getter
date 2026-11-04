package main

import "fmt"

// Demonstrate polymorphism using interfaces
type Animal interface {
	speak()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) speak() {
	fmt.Println("Hello Everyone! My name is ", d.name)
}

func (c Cat) speak() {
	fmt.Println("Hi Guys!, Myself ", c.name)
}

func polymorphismUsingInterface(a Animal) {
	a.speak()
}

// Implement a custom error using interface error
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error [%d]: %s", e.Code, e.Message)
}

func throwError(val int) error {
	if val < 0 {
		return &MyError{
			Code:    400,
			Message: "Invalid Param",
		}
	}

	return nil
}

// Implement writer-like interface & mock it for testing.
type Writer interface {
	Write(data string) error
}

type ConsoleWriter struct{}

func (c ConsoleWriter) Write(data string) error {
	fmt.Println("Console output ", data)
	return nil
}

func main() {
	dog := Dog{name: "Pommy"}
	cat := Cat{name: "Mickey"}
	animal := []Animal{dog, cat}

	for _, val := range animal {
		polymorphismUsingInterface(val)
	}

	err := throwError(-5)
	if err != nil {
		fmt.Println(err)
	}

	c := ConsoleWriter{}
	err = c.Write("Hello World!")
	if err != nil {
		fmt.Println(err)
	}
}
