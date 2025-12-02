package main

import (
	"fmt"
	"reflect"
)

type TagTypeStruct struct {
	field1 bool   "An important answer"
	field2 string `The name of the thing`
	field3 int    "How much there are"
}

type T struct {
	a string "This is a tag"
	b int    `A raw string tag`
	c int    `key1:"value1" key2:"value2"` // definition of multiple tags in Go
}

// Exercise
type Book struct {
	Title  string `json:"title" genre:"fiction"`
	Author string `json:"author" genre:"non-fiction"`
}

func refTag(tt TagTypeStruct, ix int) {
	ttType := reflect.TypeOf(tt)
	// ttValue := reflect.ValueOf(tt)

	ixField := ttType.Field(ix)
	fmt.Printf("%s\n", ixField.Name)
	fmt.Printf("%v\n", ixField.Tag)
}

func analyzeStructTags(s reflect.Type) {
	numFields := s.NumField()

	for i := 0; i < numFields; i++ {
		fmt.Printf("Printing information %d field of struct\n", i)
		fmt.Printf("%v\n", s.Field(i))
		fmt.Printf("Printing JSON: %v\n", s.Field(i).Tag.Get("json"))
		fmt.Printf("Printing Genre: %v\n", s.Field(i).Tag.Get("genre"))

	}
}

func structsTagsMain() {
	tt := TagTypeStruct{true, "Barack Obama", 1}

	// Print out each field name and tags
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

	t := T{}
	fmt.Println(reflect.TypeOf(t).Field(0).Tag)

	if field, ok := reflect.TypeOf(t).FieldByName("b"); ok {
		fmt.Println(field.Tag)
	}
	// Using Get method we can get the definition from each tag
	if field, ok := reflect.TypeOf(t).FieldByName("c"); ok {
		fmt.Println(field.Tag.Get("key1"))
		fmt.Println(field.Tag.Get("key2"))
	}

	// if a field does not exist:
	if field, ok := reflect.TypeOf(t).FieldByName("d"); ok {
		fmt.Println(field.Tag)
	} else {
		fmt.Println("Field not found")
	}

	// Doing the exercise:'
	b := Book{}
	analyzeStructTags(reflect.TypeOf(b))
}
