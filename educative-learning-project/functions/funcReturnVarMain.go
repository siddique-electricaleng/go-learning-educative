package main

import "fmt"

func funcReturnVarMain() {

	f1 := Add2()
	fmt.Printf("Adding 2 to 5: %d\n", f1(5))

	f2 := Adder(5)
	fmt.Printf("Adding 3 to 5: %d\n", f2(3))

	f3 := AdderClosure()
	fmt.Println(f3(3))
	fmt.Println(f3(21))
	fmt.Println(f3(4))
	fmt.Println(f3(10))
}

// The following teach how to use functions as return variables
func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

// Actual Closer in Go
func AdderClosure() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
