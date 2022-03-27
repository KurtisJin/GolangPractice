package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
	phone   int
}

type human struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	person1 := human{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(person1)
	person2 := human{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "test@test.com",
			zipCode: 75001,
		},
	}
	person1.print()

	person2.updateName("Bob")
	person2.print()
}

func (h human) print() {
	fmt.Println(h)
}

func (h human) updateName(newFirstName string) {
	h.firstName = newFirstName
}
