package main

import (
	"fmt"
	mx "private-properties-example/matrix"
)

func main() {
	fmt.Println("Testing private properties analogy in Go using visibility rule")
	matrixPtr := mx.NewMatrix(3, 5)

	fmt.Println("Printing the matrix, initialized using Factory Method:")
	fmt.Printf("%v\n", *matrixPtr)
}
