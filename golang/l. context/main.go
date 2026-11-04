package main

import "time"

// context
// 1. context allows you to signal that the operation can be canceled
// 2. you can associate a deadline with a context, specifying the maximum time for an operation to complete
// 3. context can be used to carry arbitrary data along with the function calls

// uses
// 1. HTTP Request context can carry values through middlewares and handlers
// 2. Database operations can rely on context for cancellation
// 3. Long-running go routines can be canceled with context

// context interface
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

// 1. context carries a deadline, cancellation signal and request scoped values across API boundaries and go routines
// so its methods are safe for simultaneous use by multiple go routines
// 2. The deadline here returns the time when the context will be canceled if a deadline is set
// 3. The done function returns the channel that is closed when this context is canceled or times out
// 4. The error method indicates why this context was canceled or why the done channel was closed
// 5. Value method returns any associated value with the key that you provided or nil if it doesn't exist

// General Rules
// 1. Context creation should be done at the root of the execution call
// 2. To create context, we use context.Background function that returns an empty context
// 3. By convention, context is the first parameter of a function
// 4. Context should not be saved in a variable for later use
// 5. Context should not be mutated
