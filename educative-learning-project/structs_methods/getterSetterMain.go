package main

import (
	"fmt"
	pr "goTut/structs_methods/person"
)

func getSetMain() {
	// Instantiating the struct like an object
	p := new(pr.Person)

	// Using the setters
	p.SetFirstName("Nate")
	p.SetLastName("Higgerson")

	fmt.Println("Printing the field values")
	fmt.Printf("FirstName: %s\n", p.GetFirstName())
	fmt.Printf("LastName: %s\n", p.GetLastName())
}
