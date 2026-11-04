package main

import (
	"errors"
	"fmt"
)

// Sentinel Errors
// 1. Indicates occurrence of a unique event
// 2. Exported package level variables
// 3. Example from database/sql package:
// var ErrNoRows = erros.New("sql: no rows in result set")
// 4. To check for specific error you can do
// if err == sql.ErrNoRows

// Error Wrapping
// 1. Wrap errors to add additional information
// 2. An error e wraps another error if e's type has one of the methods
// a. Unwrap() error
// b. Unwrap() []error
// 3. An error chain refers to the series of errors that are wrapped together also called error tree

// How to wrap Error
// 1. We can use fmt.Errorf() with %w to wrap errors
// 2. If we don't want to wrap error but just want to add error from, use the %v placeholder in fmt.Errorf()

func firstFunction() error {
	return fmt.Errorf("original error: something went wrong in the first function")
}

func secondFunction() error {
	firstErr := firstFunction()
	if firstErr != nil {
		return fmt.Errorf("failed in second function: %w", firstErr)
	}

	return nil
}

// 3. Use errors.Join()
// a. Join wraps provided error, discarding nil errors
// b. If all errors are nil, it returns nil
// c. A non nil error returned by Join implements the unwrap() []error method

func joinErrors() error {
	firstErr := firstFunction()
	if firstErr != nil {
		err := errors.New("something went wrong in second function")
		return errors.Join(firstErr, err)
	}

	return nil
}

// Unwrapping Errors
// 1. Unwrap calls the unwrap method implemented on the error or returns nil
// 2. fmt.Errorf() - creates error implements the Unwrap method
// 3. Please note that Unwrap does not unwrap the error returned by errors.Join()
// 4. If e.Unwrap() returns a non nil error w or a slice containing w, then we say that e wraps w
// 5. It is invalid for Unwrap method to return an []error containing a nil error value
// 6. A nil error returned from e.Unwrap() indicates that e does not wrap any error

func main() {
	err := secondFunction()
	fmt.Println("Original error: ", err) // failed in second function: original error: something went wrong in the first function

	fmt.Println(joinErrors()) // original error: something went wrong in the first function
	// something went wrong in second function

	innerErr := errors.Unwrap(err)
	fmt.Println("Inner error: ", innerErr) //  original error: something went wrong in the first function
}
