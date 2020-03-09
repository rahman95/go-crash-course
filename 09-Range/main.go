package main

import "fmt"

func main() {
	// Define
	ids := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	users := map[string]string{"Derrick": "derick@mail.com"}

	// Loop through w/ index
	for i, id := range ids {
		fmt.Printf("%d - %d \n", i, id)
	}

	// Loop w/o index
	for _, id := range ids {
		fmt.Println(id)
	}

	// Sum
	i := 0
	for _, id := range ids {
		i += id
	}
	fmt.Println(i)

	// Itterate maps
	for k, v := range users {
		fmt.Printf("Name: %s - Email: %s \n", k, v)
	}
}
