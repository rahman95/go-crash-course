package main

import (
	"fmt"
	"sync"
	"time"
)

// See https://gobyexample.com/waitgroups for further info, we will use WaitGroup to wait for go routines to complete on all thread before terminating.
func main() {
	// we define a waitGroup and add go routines to them, this allows us to wait on the main thread until both go routines are complete
	var wg sync.WaitGroup

	// add goroutine to waitGroup and execute goroutine
	wg.Add(1)
	go count("sheep", &wg)

	// add goroutine to waitGroup and execute goroutine
	wg.Add(1)
	go count("birds", &wg)

	// wait for all waitGroup items to complete
	wg.Wait()
}

func count(thing string, wg *sync.WaitGroup) {
	// the defer prefix will allow us to run this statement last once the rest of execution in this method has ran.
	defer wg.Done()

	for i := 0; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
