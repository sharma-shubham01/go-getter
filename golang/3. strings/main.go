package main

import "fmt"

// strings are stored as sequence of bytes
// you can access each byte of the string data using index notation
// the slice expression works with strings
// strings in go are immutable, you cannot manipulate a string, once created
// the string, rune, byte type can be converted to each other

func main() {
	var s = "Code & Learn"

	fmt.Println(s[0])  // C
	fmt.Println(s[:5]) // Code
	fmt.Println(s[3:]) // e & Learn
	fmt.Println(s[:])  // Code & Learn

	//s[1] = 'a' // error -> unassignable operand // strings are immutable
}
