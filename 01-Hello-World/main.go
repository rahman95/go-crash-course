package main

import "fmt"

const me string = "rahman95"

func main() {
	var str string = fmt.Sprintf("Hello World from %s 👋", me)
	fmt.Println(str)
}
