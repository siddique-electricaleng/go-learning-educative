package main

import "fmt"

type FooMap map[string]string

type FooStruct struct {
	fieldOne string
	fieldTwo string
}

func newMakeDiffMain() {
	newStruct := new(FooStruct)
	(*newStruct).fieldOne = "Struct Field One"
	(*newStruct).fieldTwo = "Struct Field Two"
	fmt.Println("Printing the struct")
	fmt.Printf("%#v\n", *newStruct)

	// Why is this an error
	/* anotherNewStruct := make(FooStruct)
	fmt.Println(anotherNewStruct)
	*/
	newMap := make(FooMap)
	newMap["itemOne"] = "Value One"
	newMap["itemTwo"] = "Value Two"
	newMap["itemThree"] = "Value Three"
	fmt.Println("Printing the Map:")
	fmt.Printf("%#v\n", newMap)

	// Assignment to nil map - because new returns a nil pointer and there is no memory address for map, therefore it cannot be given any items inside it
	anotherNewMap := new(FooMap)
	fmt.Println(*anotherNewMap)
	// The following gives a runtime errro - for not having any address allocated for the map
	// (*anotherNewMap)["newItem"] = "ANtoher new item"

}
