package main

import (
	"fmt"
	"strings"
)

// Struct declaration - type
type s struct {
	a string
	b string
}

type Point struct {
	x int
	y int
}

// Struct containing other structs
type RectangleA struct {
	vertexA Point
	vertexB Point
}

// Since this struct of structs holds pointers meaning they are not in contiguous memory
type RectangleB struct {
	vertexA *Point
	vertexB *Point
}

func structsMethodsMainFirst() {
	// TASK 1: Initialize struct in various ways

	// Field value assigning using structs
	// structname.fieldname = value

	// t is a value of type s struct
	// Struct as a pointer
	var tPtr *s = new(s) // returns a pointer

	// Using dot-notation also called selector-notation to assign values to fields, of the value t, of type struct s
	tPtr.a = "jhasd"
	tPtr.b = "JBAJDBAS"
	// (*tPtr).b = 89

	upStruct(tPtr)
	fmt.Printf("Printing out the struct:%v\n", *tPtr)

	// Another way to initialize a struct value - as a struct literal
	tNewPtr := &s{"hbhdba", "ENB1"}

	upStruct(tNewPtr)
	fmt.Printf("Printing out the struct:%v\n", *tNewPtr)

	// Another way - as a value type
	tLiteral := s{"jajda", "hAAUBD"}

	upStruct(&tLiteral)
	fmt.Printf("Printing out the struct:%v\n", tLiteral)

	// TASK 2 :Struct of Structs
	// just initializing a value type
	rectA := RectangleA{Point{10, 20}, Point{30, 20}}
	fmt.Printf("Printing out rectA:%v\n", rectA)

	// just initiliazing as a pointer
	var rectB *RectangleB = new(RectangleB)

	// Recall in the declaration of how this struct of structs was created - each vertex of the rectangle is a direct pointer to the struct Pointer and not a value type
	rectB.vertexA = &Point{10, 10}
	rectB.vertexB = &Point{30, 10}
	// The following literally gives me addresses at each field name: vertexA and vertexB for RectangleB
	fmt.Printf("Printing out rectB: %+v\n", *rectB) // can't do it 1 shot to get the values of vertexA and vertexB
	// The following gives me values
	fmt.Printf("Printing out rectB:{%+v, %+v}\n", *rectB.vertexA, *rectB.vertexB)
}

// expects a struct pointer or address
func upStruct(p *s) {
	p.a = strings.ToUpper(p.a)
	p.b = strings.ToLower(p.b)
}
