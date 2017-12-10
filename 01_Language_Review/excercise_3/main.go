package main

import (
	"fmt"
)

// Person Struct
type Person struct {
	firstName string
	lastName  string
	title     string
}

func (p Person) speak() {
	fmt.Println(p.title, " ", p.firstName, " ", p.lastName, " is speaking!")
}

// SecretAgent Struct
type SecretAgent struct {
	person        Person
	licenseToKill bool
}

func (sa SecretAgent) speak() {
	fmt.Println(sa.person.title, " ", sa.person.firstName, " ", sa.person.lastName, " is speaking!")
}

// Human Interface
type Human interface {
	speak()
}

func main() {
	person := Person{"Foo", "Bar", "Mr."}
	agent := SecretAgent{
		Person{"James", "Bond", "Mr."},
		true,
	}

	person.speak()
	agent.speak()
}
