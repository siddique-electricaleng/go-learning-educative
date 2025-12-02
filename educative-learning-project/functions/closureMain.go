package main

import "fmt"

/*
1. Lambda function
2. Closures
*/

func closureMain() {

	// Lambda Function = also called closures in Go wtf
	lambdaFunc := func(x, y int) int { return x + y }
	fmt.Printf("2 + 2 = %d\n", lambdaFunc(2, 2))

	// Function invoking indirectly
	fmt.Printf("1 + 2 = %d\n", func(x, y int) int { return x + y }(1, 2))
}
