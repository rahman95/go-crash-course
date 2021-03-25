package main

import (
	"fmt"
	"time"
)

// See https://gobyexample.com/goroutines for further info, below are a few examples of using goroutines
func main() {
	// 1 - would never count sheep as the main thread will be stuck in the infinite loop counting birds.
	// count("bird")
	// count("sheep")

	// 2 - using go prefix allows us to use goroutine this will use a seperate thread for execution and the other will run on the main thread
	go count("bird")
	count("sheep")

	// 3 - cant use go prefix for both as that will place both on a seperate thread and since nothing is running on the main it will terminate
	// go count("bird")
	// go count("sheep")
}

func count(thing string) {
	for i := 0; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
