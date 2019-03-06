package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
// A value receiver doesn't change any values, just takes them in
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
// Because we are changing, we have to use the pointer with *
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "F", age: 29}
	person2 := Person{"Amber", "Johnson", "Denver", "F", 29}

	fmt.Println(person1, person2)
	fmt.Println(person1.firstName)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("Herpdorp")
	fmt.Println(person1.greet())
}
