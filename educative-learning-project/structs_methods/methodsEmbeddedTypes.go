package main

import (
	"fmt"
	"math"
)

// Will be embedded in a struct and used anonymously - Parent Struct
type newPoint struct {
	x, y float64
}

// TASK B 1: Aggregation example

// Parent
type Log struct {
	msg string
}

// Child - with aggregation has-a relationship
type Customer struct {
	Name string
	log  *Log // aggregation/named struct field
}

// Method on Embedded Type Struct
func (p *newPoint) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// Child struct - has the embedded anonymous struct
type NamedPoint struct {
	newPoint `detail1:"Anonymous type - embedded struct"`
	name     string
}

// TASK B 2: Aggregation - basically named embedded struct (with their own methods)
func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

// TASK C: Embedding - basically anonymous embedded struct (with their own methods)

func methodEmbeddedTypeMain() {

	np := new(NamedPoint)

	// Setting values of the struct fields
	np.name = "Cool Point"

	// newPoint{11.0, 12} is a value type where np = new(newPoint) gives us a pointer
	(*np).newPoint = newPoint{11.0, 12}

	fmt.Printf("The absolute value of {%0.2f, %0.2f} is:%0.3f\n", np.x, np.y, np.Abs())

	// TASK B 3: Aggregation example
	c1 := new(Customer)

	(*c1).Name = "Osama Bin Russell"

	(*c1).log = new(Log)

	(*c1).log.msg = "Al-Merecedes Driver"

	// Using the methods under Log structs

	fmt.Println((*c1).log.String())

	c1.log.Add("Is fourth in 2025 championship")

	fmt.Println(c1.log.String())

	//  Or we can also use the Log() method of *Customer to use the methods from *Log Add and String

	c2 := &Customer{"Kimi Talobentalli", &Log{"Second Mercedes Driver"}}

	fmt.Printf("%s\n", c2.Log().String())

	c2.Log().Add("A new line added")

	fmt.Printf("%s\n", c2.Log().String())

}
