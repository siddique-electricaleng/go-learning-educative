package main

import (
	"fmt"
)

func arrLiteralMain() {
	// TASK A: First array literal variant

	var arrAgeList = [5]int{1, 2, 3, 4, 5}

	var arrAgeListLimit [10]int = [10]int{1, 23, 34} // we can also put [10]int type in the type declaration for arrays

	// Printing section
	fmt.Println("Printing 1st variant - [5]int fixed array size allocation")
	fmt.Printf("Printing arrAgeList:%v\n", arrAgeList)
	fmt.Printf("Printing arrAgeList:%v\n", arrAgeListLimit)

	// TASK B: Second variant for array literals
	var arrLazy = [...]int{3, 1, 4, 5, 6}
	// var arrLazyWrong [...]int = [...]int{3, 1, 4, 5, 6} // this is illegal in Go because [...]int is not a type

	// Printing section
	fmt.Println("\nPrinting 2nd variant: [...]int dynamic array size allocation")
	fmt.Printf("Printing arrLazy:%v\n", arrLazy)

	// TASK C: Third variant for array literals - index:value syntax
	var arrKeyValue = [5]string{3: "Chris", 4: "Deal"}

	// Printing Section
	fmt.Println("\nPrinting 3rd variant - key:value")
	fmt.Printf("Printing arrKeyValue: %v\n", arrKeyValue)

	// TASK D: Pass big array to a function in a memory efficient manner

	var array3 [4]float64 = [4]float64{0.21, 0.32, 0.5, 0.1}

	// Another way to do this directly make it a pointer

	// var array3 *[4]float64 = &[4]float64{0.21, 0.32, 0.5, 0.1}

	fmt.Printf("Sum of items in the array of size %d: %0.4f\n", len(array3), Sum(&array3))

}

func Sum(arr *[4]float64) (sum float64) {
	for _, val := range *arr {
		sum += val
	}
	return
}
