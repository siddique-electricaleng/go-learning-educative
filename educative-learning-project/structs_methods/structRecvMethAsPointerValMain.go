package main

import "fmt"

type BTwo struct {
	b int
}

// TASK B 3: Methods can be called with a value or a pointer - Go automatically switches between them based on what and where it is required.

type List []int

// TASK C 1: Use Methods as values and parameters like functions

type newT struct {
	a int
}

// Methods:

// Receiver as a pointer

func (recv *BTwo) change() {
	recv.b = 1
}

// Receiver as a value

func (recv BTwo) write() string {
	return fmt.Sprint(recv)
}

// TASK B 3: Methods can be called with a value or a pointer - Go automatically switches between them based on what and where it is required.

// give length of the list - receiver as a value type

func (recv List) Len() (res int) {
	res = len(recv)
	return
}

// append values into the list - receiver as a pointer type

func (recv *List) Append(val int) {
	*recv = append(*recv, val)
}

// TASK C 2: Use Methods as values and parameters like functions

func (t newT) print(message string) {
	fmt.Println(message, t.a)
}

func (newT) hello(message string) {
	fmt.Println("Hello!", message)
}

// The function taking in a function as a parameter
func callMethod(t newT, method func(newT, string)) {
	method(t, "A message")
}

func structRecvMethAsPointerValMain() {

	someB := new(BTwo)
	someB.b = 12

	// someB acting as a pointer type receiver

	someB.change()
	fmt.Println(*someB)

	// someB acting as a value type receiver - we are dereferencing and passing the value
	fmt.Println((*someB).write())

	// TASK B 2: Methods can be called with a value or a pointer - Go automatically switches between them based on what and where it is required.

	myList := make(List, 5) // an integer slice of length 5 - make gives us a value type

	fmt.Printf("Length of slice: %d\n", myList.Len())

	myList.Append(12)

	fmt.Println("Printing the list")
	fmt.Println(myList)

	// TASK C 3: Use Methods as values and parameters like functions
	t1 := new(newT)
	t1.a = 10

	t2 := new(newT)
	t2.a = 20

	// setting f = print method
	var f func(newT, string) = newT.print

	callMethod(*t1, f)
	callMethod(*t2, f)

	callMethod(*t1, newT.hello)
}
