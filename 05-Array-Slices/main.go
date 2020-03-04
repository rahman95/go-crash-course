package main

import "fmt"

func main() {
	// Declaring array
	var fruits [3]string

	// Assigning values
	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[2] = "kiwi"

	// Shorthand - declare and assign
	animals := [2]string{"dog", "cat"}

	fmt.Println(fruits, fruits[1:3], len(fruits))
	fmt.Println(animals, animals[1:2], len(animals))
}
