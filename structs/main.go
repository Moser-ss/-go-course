package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p := person{firstName: "John",
		lastName: "Doe",
		contactInfo: contactInfo{
			email:   "john.doe@example.com",
			zipCode: 14254},
	}

	p.updateName("Jim")
	p.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
