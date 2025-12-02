package main

import "fmt"

/*
Common Terminologies and Concepts used in Higher Order Functions

Higher Order Functions accept other functions as arguments or returns a function as its result

1. Function as a value
2. Function as a parameter = callback
3. Function as a filter
*/

type flt func(int) bool

/* Callback example - a higher order function accepting another func as parameter */

func Add(a int, b int) (res int) {
	res = a + b
	return
}

func callback(y int, f func(int, int) int) int {

	// Calling the callback function
	return f(y, 2)
}

// functions to check if it is even or odd

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

// function as a filter
func filter(s []int, f flt) (res []int) {
	for _, val := range s {
		if f(val) {
			res = append(res, val)
		}
	}
	return
}

func higherOrderFuncMain() {
	fmt.Printf("Call back returned: %d\n", callback(3, Add))

	slice := []int{2, 3, 4, 5, 6, 1}

	fmt.Println("Printing the even numbers")
	fmt.Println(filter(slice, isEven))

	fmt.Println("Printing the odd numbers")
	fmt.Println(filter(slice, isOdd))
}
