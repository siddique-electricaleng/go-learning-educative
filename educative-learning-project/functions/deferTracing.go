package main

/*
Write a tracing mechanism to:

	Identify the sequence of function calls.

	Ensure that resources are released properly in the correct order.

	Implement tracing functions to log messages when entering and leaving each function.

	Use defer to call these tracing functions appropriately.

*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// openFile opens a file and returns a file pointer.
func openFile(filename string) (*os.File, error) {

	trace("openFile")
	fmt.Println("in openFile")
	defer untrace("openFile")

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// processFile processes the opened file.
func processFile(file *os.File) error {

	trace("processFile")
	fmt.Println("in processFile")
	defer untrace("processFile")
	// Do some operations on the file.
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	} else {
		defer fmt.Println("File processed successfully.")
	}
	fmt.Println(string(content))
	return nil
}

// closeFile closes the file.
func closeFile(file *os.File) {

	trace("closeFile")
	fmt.Println("in closeFile")
	defer untrace("closeFile")

	if file != nil {
		file.Close()
	}
}

func trace(s string) {
	fmt.Printf("Entering %q\n", s)
}

func untrace(s string) {
	fmt.Printf("Leaving %q\n", s)
}

func DeferTracingMain() {
	fileName := "example.txt"

	file, err := openFile(fileName)

	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}

	err = processFile(file)

	if err != nil {
		log.Fatalf("Error processing file: %v\n", err)
	}

	closeFile(file)

}
