package main

import "fmt"

// Error
// 1. Go has different way of error handling
// 2. There is no try-catch but uses if for error handling
// 3. The last return value is error by convention
// 4. The error string should not be capitalized and does not end with punctuation

// Error Interface
// 1. The error type is an interface type in Go
// 2. It has single method called error returning a string
// 3. It becomes really easy to implement your own error types
// 4. All we need to do to implement the error interface is to implement the Error method
// 5. This is especially useful when we want to create custom error messages

//type error interface {
//	Error() string
//}

// Creating Errors
// 1. Implement the error interface
// 2. Use the New function from errors package
// 3. Use the Errorf function from fmt package

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Error %d : %s", c.Code, c.Message)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, CustomError{Code: 1, Message: "cannot divided by zero"} // custom error is implementing error interface
	}

	return a / b, nil
}

func main() {
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("result: ", result)
}
