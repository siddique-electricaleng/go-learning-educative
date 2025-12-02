package main

import "fmt"

func sliceMain() {
	arr1 := [...]int{12, 2, 3, 43, 41, 3, 10}

	var slice1 []int = arr1[:3]
	slice2 := arr1[1:5]

	slice3 := [3]int{2, 3, 4}

	slice4 := [...]int{2, 3, 4, 19, 20}

	slice5 := []int{99, 3, 5}

	fmt.Printf("Slice 1 values:%v\n", slice1)
	fmt.Printf("Slice 2 values:%v\n", slice2)
	fmt.Printf("Slice 3 values:%v\n", slice3)
	fmt.Printf("Slice 4 values:%v\n", slice4)
	fmt.Printf("Slice 5 values:%v\n", slice5)

	// Expanding a slice to its maximum size - it accesses the other array values not previously accessed
	slice1MaxSize := cap(slice1)
	slice1New := slice1[:slice1MaxSize]
	fmt.Printf("New Slice 1 with Max Size is: %v\n", slice1New)

	// You cannot move a slice to negative index or below 0, when we reassign a slice. Negative indexing is not allowed in Go

	var dynamicSlice []int = arr1[1 : len(arr1)-2]
	fmt.Printf("Dynamic slice before changing: %v\n", dynamicSlice)

	dynamicSlice = dynamicSlice[1:cap(dynamicSlice)]
	fmt.Printf("Dynamic slice after changing: %v\n", dynamicSlice)

	// Slice to a function
	fmt.Printf("Sum of array items 2nd to 4th: %d\n", sum(arr1[1:4]))

	// Defining a slice using make
	//  - when underlying array is not created
	// 10 is the initial length of the slice, also the length of the array

	s2 := make([]int, 10)
	fmt.Printf("Printing out unassigned slice: %v\n", s2)

	// new() vs make(): new allocates while make initializes

	var p *[]int = new([]int) // *p == nil it is a nil pointer, with len 0 and cap 0
	fmt.Printf("printing the nil pointer p:%v\n", (*p))

	var empty_make []int = make([]int, 10, 50)
	fmt.Printf("printing the empty initialized Make:%v\n", empty_make)

}

// Passing a slice to a function
func sum(a []int) int {
	s := 0
	for _, val := range a {
		s += val
	}
	return s
}
