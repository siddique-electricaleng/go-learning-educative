package main

import (
	"fmt"
	"time"
)

type S struct {
	a int
}

type SType S    // a new type
type SAlias = S // this is an alias of S

type IntType int // new type

/*
It is not possible to redefine the methods on the int Alias type because it is the same as the original but we can define new methods
on the new type based on the int above
*/

type IntAlias = int // alias of int

type TwoInts struct {
	a int
	b int
}

// TASK IMP A: Way/Workaround of the error of defining methods in non-local packages - alias or anynonymous field in a struct

type myTime struct {
	time.Time //anonymous field
}

// Methods:
/*
The general format to write a method is:

	func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

	The method is called using:
		recv.methodName(argument list)
*/

// Example of overloading - In Go overloading is possible bcz different receiver types can still have the same func name
// it doesn't mean they use the same function but different function definition, just with the same name as shown below in a simple example

func (recv S) print() {
	// recv is the receiver that receives this method and is of type S or an alias of S
	fmt.Printf("%[1]v\n", recv)
}

func (recv SType) print() {
	// recv is the receiver that receives this method and is of type S or an alias of S
	fmt.Printf("%[1]v\n", recv)
}

func (recv IntType) print() {
	fmt.Printf("%[1]v\n", recv)
}

func (recv *TwoInts) AddThem() (res int) {
	res = recv.a + recv.b
	return
}

func (recv *TwoInts) AddToParam(num int) (res int) {
	res = recv.a + recv.b + num
	return
}

// TASK IMP B: Way/Workaround of the error of defining methods in non-local packages - alias or anynonymous field in a struct

func (t *myTime) first3Chars() string {
	return t.String()[0:3]
}

// The main function
func structMethodsMain() {
	a := S{10}
	a.print()

	b := SType{20}
	b.print()

	c := IntType(40)
	c.print()

	d := new(TwoInts)
	d.a = 11
	d.b = 45

	e := TwoInts{20, 20}

	fmt.Printf("The sum is: %d\n", d.AddThem())
	fmt.Printf("The sum is: %d\n", e.AddThem())

	fmt.Printf("Add 30 to d: %d\n", d.AddToParam(30))
	fmt.Printf("Add 30 to e: %d\n", e.AddToParam(30))

	fmt.Println(d)
	fmt.Println(e)

	// TASK IMP C: Way/Workaround of the error of defining methods in non-local packages - alias or anynonymous field in a struct
	t := new(myTime)
	t.Time = time.Now()
	fmt.Printf("Full time now: %v\n", t.String())
	fmt.Println(t.first3Chars())
}
