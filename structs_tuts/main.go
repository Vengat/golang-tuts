package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {
	vengat := person{
		firstName: "Vengat",
		lastName:  "Ramar",
		// contact: contactInfo{
		// 	email: "vengatramanan@gmail.com",
		// 	zipCode: 560087,
		// },
		contactInfo: contactInfo{
			email:   "vengatramanan@gmail.com",
			zipCode: 560087,
		},
	}

	fmt.Println(vengat.firstName)
	fmt.Printf("%+v", vengat)
	vengat.print()
	vengat.updateName("Vengatramanan")
	vengat.print()
	vengatPointer := &vengat
	vengatPointer.updateNameWithPointer("Vengatramanan")
	vengatPointer.print()
	vengat.updateNameWithPointer("Vengatramanan") // also works
}

func (p person) updateName(newFirstName string) {
	// This function creates a copy of the message and stores in another address in the ram
	//Value in the original address never gets updated
	p.firstName = newFirstName //ineffectual this wont work
}

func (pointerToPerson *person) updateNameWithPointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
