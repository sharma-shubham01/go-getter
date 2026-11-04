package main

import "fmt"

// Method Expression
// 1. There is another way of using methods defined on types called method expression
// 2. This requires us to save a method in a variable
// 3. The first argument to this function will be a receiver

type Person struct {
	Name string
	Age  int
}

func (p Person) GetDetails() string {
	return fmt.Sprintf("Name - %s, Age - %d", p.Name, p.Age)
}

type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) SetLength(l int) {
	r.Length = l
}

func main() {
	f1 := Person.GetDetails // method expression

	p := Person{Name: "Shubham", Age: 25}

	// the first argument to this function will be a receiver
	fmt.Println(f1(p))

	// method expression for the set length method of Rectangle
	f2 := (*Rectangle).SetLength

	// creating an instance of Rectangle
	r := &Rectangle{Length: 5, Width: 10}

	// calling the method expression as a function
	f2(r, 1)

	// check is set length modified the length of a rectangle
	fmt.Println(r.Length)
}
