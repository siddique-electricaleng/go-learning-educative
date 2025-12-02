package main

import "fmt"

func findFactorialMain() {
	var testNumbers = [...]uint64{5}
	for _, val := range testNumbers {
		fmt.Printf("Factorial of %d: %v\n", val, factorialRecursion(val))
	}
}

func factorialRecursion(num uint64) (fac uint64) {
	if num == 1 {
		return 1
	}
	fac = num * factorialRecursion(num-1)
	return
}
