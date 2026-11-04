package main

import "fmt"

// Panic
// 1. Panic occurs when a Go program has reached a state of no return
// 2. Ex - trying to access an element that is out of bounds of an array causes runtime panic
// 3. We can raise our own panic with the panic() function
// 4. The panic function accepts an argument of type any
// 5. When panic is called, the code execution stop and panic prints the error stack of where the panic occurred
// 6. The defer functions are guaranteed to be called if there are any incase of even a panic

func function1() {
	defer func() {
		fmt.Println("Func1 deferred function called")
	}()

	function2()
}

func function2() {
	defer func() {
		fmt.Println("Func2 deferred function called")
	}()

	panic("the function panicked")
}

func main() {
	defer func() {
		fmt.Println("main function deferred")
	}()

	function1()

	fmt.Println("Start of the program")
	panicExample()
	fmt.Println("End of the program")
}

// Recover
// 1. The recover() function can be used to recover from panics
// 2. Function recover() needs to be called in a deferred function - bcz only deferred functions are called during panic
// 3. The return value of calling recover() is the argument provided to panic
// 4. Should be limited to graceful shutdown

func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	panic("something went wrong")
}
