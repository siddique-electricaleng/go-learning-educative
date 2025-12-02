package main

import "fmt"

func MulRecursiveMain() {
	fmt.Printf("The number %d is even? %t\n", 16, even(16))
	fmt.Printf("The number %d is odd? %t\n", 18, odd(18))
	fmt.Printf("The number %d is odd? %t\n", 17, odd(17))
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(AbsVal(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(AbsVal(nr) - 1)
}

func AbsVal(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
