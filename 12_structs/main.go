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

//Methods go outside the struct

// value receiver methods are the methods that do only calculations
// pointer receivers when you actually are changing something

// Greetings is a value receiver

func (p Person) greet() string {
	return "hello " + p.firstName + " " + p.lastName + " " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method ()  (pointer receiver)

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Structs (like classes but not really classes)
	// two kinds of methods associated (pointer receivers and value receivers)
	fmt.Println("Structs")

	// Init person using struct

	person1 := Person{firstName: "thomas", lastName: "mueller", city: "Berlin", gender: "F", age: 29}

	person2 := Person{"John", "Doe", "Alexandria", "F", 26}

	fmt.Println(person1)
	fmt.Println(person2)

	// Getting a single field

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

	// pointer receiver

	person1.hasBirthday()

	person1.getMarried("williams")

	//calling the value receiver
	fmt.Println(person1.greet())

}
