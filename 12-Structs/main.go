package main

import "fmt"

// Person struct
type Person struct {
	firstName string
	lastName  string
	// Alternative to group them
	// fistName, lastName string
	age int
}

// value receiver
func (person Person) greet() string {
	return fmt.Sprintf("Hello, My name is %s %s and I am %d years old!", person.firstName, person.lastName, person.age)
}

// pointer receiver
func (person *Person) hasBirthday() {
	person.age++
}

// pointer receiver
func (person *Person) changesFirstName(firstName string) {
	person.firstName = firstName
}

// pointer receiver
func (person *Person) changesLastName(lastName string) {
	person.lastName = lastName
}

func main() {
	person := Person{firstName: "Rahman", lastName: "Younus", age: 25}

	// Shorthand
	person1 := Person{"Ada", "Lovelace", 42}

	fmt.Println(person, person1)

	person.hasBirthday()            // 25 -> 26
	person.changesFirstName("Bob")  // Rahman -> Bob
	person.changesLastName("Smith") // Younus -> Smith

	fmt.Println(person.greet())
}
