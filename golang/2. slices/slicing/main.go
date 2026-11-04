package main

import "fmt"

// slicing notation -> arr[x:y]
// 1st element(x) -> inclusive
// 2ns element(y) -> exclusive

func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	b := a[2:5]
	fmt.Println(b) // 3 4 5

	c := a[:5]
	fmt.Println(c) // 1 2 3 4 5

	d := a[3:]
	fmt.Println(d) // 4 5 6

	e := a[:]
	fmt.Println(e) // 1 2 3 4 5 6

	// Problem with Slices
	f := a[:3]
	fmt.Println(b) // 1 2 3

	// Now, a and b both slices will point to same underlying array
	// any update in a will get reflect in b as well
	// this is known as ripple effect

	f[0] = 10
	fmt.Println(a) // 10 2 3 4 5 6
	fmt.Println(f) // 10 2 3 4 5 6

	g := a[:3]
	fmt.Println(g) // 1 2 3

	g = append(g, 15)
	fmt.Println(a) // 1 2 3 15 5 6
	fmt.Println(g) // 1 2 3 15

	h := a[2:]

	h = append(h, 15) // since the of underlying array is full, it will point to new array now
	fmt.Println(a)    // 1 2 3 4 5 6
	fmt.Println(h)    // 3 4 5 6 15

	// Conclusion - value update in any child slice, will affect the parent slice as well or vice versa

	// Passing a capacity to the slice
	i := a[2:6:6]
	fmt.Println(i)

	// Copy Function -> copy(destination, source)
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := make([]int, 6)

	copy(arr2, arr1)
	fmt.Println(arr2) // 1 2 3 4 5 6

	arr3 := make([]int, 3)

	copy(arr3, arr1)
	fmt.Println(arr3) // 1 2 3

}
