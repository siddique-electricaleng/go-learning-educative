package main

import (
	"fmt"
	"reflect"
)

type innerS struct {
	int1 int
	int2 int
}
type outerS struct {
	b int
	c float32

	// the following are anonymous fields - because they only have types and not names
	int
	innerS // a struct type = embedded struct is similar to composition - Go favors composition over traditional inheritance
}

type A struct{ a int }
type B struct{ a, b int }
type C struct {
	A
	B
}
type D struct {
	B
	b float32
}

func anonymousEmbeddedStructsMain() {

	// TASK 1 : Embedded Structs Example
	outer := new(outerS)
	outer.b = 12
	outer.c = 12.12
	outer.int = 53
	outer.int1 = 12
	outer.int2 = 14

	printStruct(*outer)

	// A literal struct
	outer2 := outerS{12, 2.3, 1, innerS{12, 1}}
	printStruct(outer2)

	// TASK 2: Conflicting field names or methods
	/*
		Rules on confliction field name or method - for inner and outer structs ie embedded structs
			1. An outer name hides an inner name. This provides a way to override a field or method.
			2. If the same name appears twice at the same level, it is an error if the program uses the name. If it’s not used, it doesn’t matter. There are no rules to resolve the ambiguity; it must be fixed.
	*/

	var c C
	fmt.Println(c.A.a)
	fmt.Println(c.B.a)
	fmt.Println(c.B.b)

	var d D
	fmt.Println(d.b) //float32 for d

	// TASK 3: Anonymous Structs - structs without the keyword 'type' hence t he type name

	anonymousStruct := struct {
		name, surname string
	}{"Donald", "Trump"}

	fmt.Println(anonymousStruct.name, anonymousStruct.surname)

	// Usefulness of anonymous structs - instead of a fixed struct with fixed fields anonymous structs can be used to create dynamic adaptable data structures on the fly
	// e.g. there could be servers with changing configurations and this is how we would declare them
	/*
		serverConfig := struct {
			Hostname    string
			IP          string
			Port        int
			Environment string
			Credentials struct {
				Username string
				Password string
			}
		}{
			Hostname:    "localhost",
			IP:          "192.168.0.1",
			Port:        8080,
			Environment: "production",
			Credentials: struct {
				Username string
				Password string
			}{
				Username: "root",
				Password: "1234",
			},
		}
	*/
}

func printStruct(s outerS) {
	fmt.Println("Printing New Struct")
	structType := reflect.TypeOf(s)
	structValue := reflect.ValueOf(s)
	numFields := structType.NumField()
	fmt.Println("Field Name: Field Value")
	for i := 0; i < numFields; i++ {
		fmt.Printf("%v :%v\t", structType.Field(i).Name, structValue.Field(i))
	}
	fmt.Println()
}
