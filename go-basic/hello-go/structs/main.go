package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "matt",
		contactInfo: contactInfo{
			zip: 3,
		},
	}

	alex.updateFirstName("John")
	alex.print()

}

func (p *person) updateFirstName(newName string) {
	p.firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
