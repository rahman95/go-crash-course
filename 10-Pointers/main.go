package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T %T \n", a, b)

	// To get actual value instead of memory address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Update value of pointer
	*b = 20
	fmt.Println(a, *b)
}
