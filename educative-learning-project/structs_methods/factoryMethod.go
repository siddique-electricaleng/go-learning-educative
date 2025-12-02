package main

import (
	"fmt"
)

type File struct {
	fd   int
	name string
}

// This is the facotry  method like constructors using new in other OO languages

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func structFactoryMethods() {
	// Using the factory method - its function is to initialize a variable of type File struct.
	// returns a pointer of type File struct

	// The factory instantiates an object of the defined type here: File struct
	f := NewFile(10, ".\\test.txt")
	fmt.Println(*f)

	// To apply private properties and methods - use the package visibility rule - capital letter to be exported and small letter to be kept private

}
