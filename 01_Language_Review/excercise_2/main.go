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

func (p Person) pSpeak() {
	fmt.Println(p.title, " ", p.firstName, " ", p.lastName, " is speaking!")
}

// SecretAgent Struct
type SecretAgent struct {
	person        Person
	licenseToKill bool
}

func (sa SecretAgent) saSpeak() {
	fmt.Println(sa.person.title, " ", sa.person.firstName, " ", sa.person.lastName, " is speaking!")
}

func main() {
	person := Person{"Foo", "Bar", "Mr."}
	agent := SecretAgent{
		Person{"James", "Bond", "Mr."},
		true,
	}

	fmt.Println(person.lastName)
	person.pSpeak()

	fmt.Println(agent.person.lastName)
	agent.saSpeak()

	agent.person.pSpeak() //calling embeded type method
}
