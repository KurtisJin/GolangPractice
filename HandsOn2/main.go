package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	sa bool
}

type bothSpeak interface {
	speak() string
}

func main() {

	per := person{
		firstName: "Bob",
		lastName:  "James",
	}
	sa := secretAgent{
		person{
			firstName: "James",
			lastName:  "Bond",
		},
		true,
	}

}

func (p person) speak() {
	fmt.Println("Hello", p.firstName, "is your last name", p.lastName)

}

func (s secretAgent) speak() {
	fmt.Println(s.firstName, s.lastName, "is a secret agent?", s.sa)

}
