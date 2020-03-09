package main

import "fmt"

func main() {
	// Long
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		fmt.Println(i)
	}
}
