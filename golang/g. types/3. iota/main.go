package main

import "fmt"

// IOTA
// 1. Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants
// 2. It can be used to define the set of pre-defined limited values

type Size int

const (
	extraSmallSize = iota
	small
	medium
	large
	extraLarge
)

func main() {
	fmt.Println(extraSmallSize, small, medium, large, extraLarge)
}
