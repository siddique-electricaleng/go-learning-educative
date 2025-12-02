package main

import "fmt"

// TASK 1: Your task is to review the code and identify potential issues related to variable scope and usage, with a focus on the concepts listed below:

// Appropriate use of global scope variables

// Appropriate use of local scope variables

// Appropriate use of value type and reference type variables

// Appropriate use of constants and expressions

// Appropriate use of variable declarations and reuse

// What potential issues can you identify?
// MY ANSWER

/*
func main() {
	// Local variable declaration
	var total int
	// Block level local variable
	var sum int

	count := 5
	multiplier := 2

	sum = scaleValue(count, multiplier)
	total = sum

	fmt.Println("Sum:", sum)

	count = 10
	multiplier = 3

	sum = scaleValue(count, multiplier)
	total += sum

	fmt.Println("Sum:", sum)
	fmt.Println("Total:", total)
}

func scaleValue(count int, factor int) int {
	sum := count * factor
	return sum
} */

// Ed's Answer:
const (
	Factor1 = 2
	Factor2 = 3
)

func main() {
	sum1 := scaleValue(5, Factor1)
	fmt.Println("Sum1:", sum1)

	sum2 := scaleValue(5, Factor2)
	fmt.Println("Sum2:", sum2)

	total := sum1 + sum2
	fmt.Println("Total:", total)

	x := 3
	y := 5
	z := (x ^ y) & (x << 2)
	fmt.Println(z)
	somestring := "Hello, World!"
	fmt.Println(len(somestring))
}

func scaleValue(count int, factor int) int {
	sum := count * factor
	return sum
}
