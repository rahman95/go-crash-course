package main

import "fmt"

func main() {
	// If, else block
	x := 10
	y := 15

	if x > y {
		fmt.Printf("%d is larger than %d \n", x, y)
	} else {
		fmt.Printf("%d is larger than %d \n", y, x)
	}

	// If, else if, else block
	color := "blue"

	if color == "blue" {
		fmt.Println("color is blue")
	} else if color == "red" {
		fmt.Println("color is red")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	// Switch block
	animal := "dog"

	switch animal {
	case "dog":
		fmt.Println("animal is dog")
	case "cat":
		fmt.Println("animal is cat")
	default:
		fmt.Println("animal is NOT dog or cat")
	}
}
