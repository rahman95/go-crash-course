package main

import "fmt"

func greeting(name string) string {
	return "Hello 👋 from " + name
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("rahman95"))
	fmt.Println(sum(15, 10))
}
