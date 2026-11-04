package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// errors.Is()
// 1. The function accepts two parameters, the first is the error and second is the instance of the error we are checking against
// 2. Function signature - func Is(err, target error) bool
// 3. Iterate through the errors in the error chain and return true if matching instance found

var ErrFirstError = CustomError{Message: "first error"} // sentinel error

type CustomError struct { // custom error
	Message string
}

func (e CustomError) Error() string { // to satisfy the errors interface
	return e.Message
}

func someFunction() error {
	return ErrFirstError
}

// errors.As()
// 1. The function accepts two parameters, the first is the error and second a non-nil pointer to either a type that implements
// error or to any interface type
// 2. Function signature - func As(err error, target any) bool
// 3. Iterate through the errors in the error chain and matching target type, if found returns true and set the target to that error value

func main() {
	err := someFunction()
	if err != nil {
		if errors.Is(err, ErrFirstError) {
			fmt.Println("Sentinel Error Found")
		}
	}

	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path", pathError.Path)
		}
	}
}

// Summary
// 1. Both errors.Is and errors.As functions are used to compare errors and work with wrapped errors
// 2. The errors.Is function is used to check for error instance, the error.As function is used to check for the error type
