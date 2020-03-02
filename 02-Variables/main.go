package main

import "fmt"

// Data Types
// - string
// - bool
// - int, int8, int16, int32, int64
// - uint, unit8, uint16, uint32, uint64, uintptr
// - byte = alias for uint8
// - rune = alias for int32
// - float32, float64
// - complex64, complex128

// If no type passed, will be infered

// Global Level
const user = "rahman95"

var isCool = false

func main() {
	// Function Level
	isCool = true
	var age uint8 = 25

	// Shorthand
	// firstname := "Rahman"
	// lastname := "Younus"

	// Multi-assign
	firstname, lastname := "Rahman", "Younus"

	fmt.Println(firstname, lastname, user, age, isCool)
	fmt.Printf("%T, %T, %T, %T, %T \n", firstname, lastname, user, age, isCool)
}
