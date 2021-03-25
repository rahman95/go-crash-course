package main

import (
	"fmt"
	"time"
)

// See https://gobyexample.com/channels for further info, we will use channels to communicate from and to goroutines

func main() {
	messages := make(chan string)

	go count("sheep", messages)
	go count("bird", messages)

	for msg := range messages {
		fmt.Println(msg)
	}
}

func count(thing string, messages chan string) {
	for i := 0; true; i++ {
		messages <- fmt.Sprint(i, " ", thing)
		time.Sleep(time.Millisecond * 500)
	}

	close(messages)
}
