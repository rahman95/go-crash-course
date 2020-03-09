package main

import "fmt"

func main() {
	// Define
	users := make(map[string]string)

	// Shorthand Define
	users2 := map[string]string{"Derrick": "derick@mail.com"}
	fmt.Println(users2)

	// Assign
	users["Alice"] = "alice@example.com"
	users["Bob"] = "bob@fastmail.com"
	users["Charles"] = "charlie@gmail.com"
	fmt.Println(users)

	// Get
	fmt.Println(users["Bob"])

	// Update
	users["Bob"] = "bob+new@fastmail.com"
	fmt.Println(users["Bob"])

	// Delete
	delete(users, "Alice")
	fmt.Println(users)

	// Length
	fmt.Println(len(users))
}
