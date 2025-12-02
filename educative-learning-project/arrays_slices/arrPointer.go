package main

import "fmt"

func arrPtrMain() {

	var arr1 = new([5]int)

	// The following is done bcz arr1 is of type *[5]int so it holds the memory address anyway and shouldn't require dereferencing
	// yet I wanna see what happens when I dereference it vs when I don't
	fmt.Println("Printing arr1 before any assignment - without ampersand - literally gives me value, instead of address, so go is silently dereferencing here")
	fmt.Println(arr1[0])
	fmt.Println("Printing address of arr1 before any assignment - using ampersand")
	fmt.Println(&arr1[0])

	var arr3 [5]int = [5]int{1, 2, 4, 5, 6}

	arr1 = &arr3

	fmt.Println("Printing address of arr1 after any assignment")
	fmt.Println(&arr1[0])

	fmt.Println("Printing address of arr3 it should be same as arr1 because arr1 is pointing to arr3's first element")
	fmt.Println(&arr3[0])

	// implicitly dereferences the pointer
	for _, val := range arr1 {
		fmt.Printf("Printing the arr1 value: %v\n", val)
	}

	for _, val := range arr3 {
		fmt.Printf("Printing the arr3 values:%v\n", val)
	}

}
