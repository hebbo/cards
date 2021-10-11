package structs

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func Call() {
	alex := person{firstName: "Alex", lastName: "Anderson", contactInfo: contactInfo{email: "email.com", zipCode: 2481}}
	alex.print()

	alex.updateName("tom")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}
