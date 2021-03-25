package main

import "fmt"

// See https://gobyexample.com/channel-synchronization for further info, fibonacci examples on main thread and using multiple threads
func main() {
	// 1 - standard fibonacci on main thread
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i, fib(i))
	// }

	// 2 - fibonacci using workerpools and channels
	jobs := make(chan int, 100)
	results := make(chan string, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- string) {
	for n := range jobs {
		results <- fmt.Sprint(n, fib(n))
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
