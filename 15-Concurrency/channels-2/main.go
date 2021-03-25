package main

import (
	"fmt"
	"time"
)

// See https://gobyexample.com/channels for further info, we will use channels to communicate from and to goroutines
// another example where we use multiple channls and can use select statement to listen for update

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// every 500ms
	go count("sheep", c1, time.Millisecond*500)

	// every 2 secs
	go count("bird", c2, time.Second*2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func count(thing string, channel chan string, delay time.Duration) {
	for i := 0; true; i++ {
		channel <- fmt.Sprint(i, " ", thing)
		time.Sleep(delay)
	}

	close(channel)
}
