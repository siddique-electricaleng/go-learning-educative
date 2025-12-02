package main

import "fmt"

func main() {
	var i2 = 200
	var intP *int = &i2
	fmt.Printf("The value of i1:%d\nThe memory address of i1:%p\n", i2, intP)
	i2 = 9011
	fmt.Printf("Pointer variable's value is now:%d\n", *intP)
	fmt.Printf("")
	// var intNewP *int
	fmt.Println(*intP * 3)
	fmt.Println(*intP)
}
