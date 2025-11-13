package main

import "fmt"

// enumerations
const (
	PI         = 3.14
	Ln2        = 0.693147180559945309417232121458176568075500134360255254120680009
	Log2E      = 1 / Ln2
	BILLION    = 1e9
	HARD_EIGHT = (1 << 100) >> 97
)

// custom type
type Gender int

const (
	UNKNOWN Gender = iota
	FEMALE
	MALE
)

func PrintMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	PrintMessage("Hello, World!")

	var emptyVar string
	var s *string

	fmt.Printf("Value of emptyVar is %s\n", emptyVar)
	fmt.Printf("Value of emptyVar is %p\n", s)

	var number float32 = 5.2
	fmt.Printf("The float of number is: %0.2f\n", number)
	fmt.Printf("The int of float number is: %d\n", int(number))
	fmt.Printf("Constant: %0.2f\n", PI)

	fmt.Println("Printing CONSTANTS")
	fmt.Printf("Printing %0.23f\n", Ln2)
	fmt.Printf("Printing %0.23f\n", Log2E)
	fmt.Printf("Printing %0.0f\n", BILLION)
	fmt.Printf("Printing %d\n", HARD_EIGHT)

	fmt.Println("Printing Enumeration based Constant")
	fmt.Printf("Value of\n1. Unknown: %d\n2. MALE: %d\n3. FEMALE: %d\n", UNKNOWN, MALE, FEMALE)
	fmt.Printf("Type of UNKNOWN:%T\n", UNKNOWN)
}
