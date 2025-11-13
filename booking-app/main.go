package main

import "fmt"

// Packages are collections of source files

func PrintMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	PrintMessage("Hello, World!")
	var number float32 = 5.2
	fmt.Println(number)
	fmt.Println(int(number))
}
