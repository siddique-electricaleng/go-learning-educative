package main

import "fmt"

func primeCompositeMain() {
	var testNumbers = [...]int{2, 3, 5, 16, 17, 19, 22, 23, 26}
	for _, val := range testNumbers {
		fmt.Printf("%d Prime? →%t\n", val, isPrime(val))         // true (smallest prime)
		fmt.Printf("%d Composite? →%t\n", val, isComposite(val)) // true (smallest prime)
	}
}

func isPrime(num int) bool {
	if num <= 2 { // 2 is prime, below that are not
		return num == 2
	}
	return primeCheck(num, 2)
}

func isComposite(num int) bool {
	if num < 4 {
		return false // -ve numbers, 0, 1, 2, 3 are not composite
	}
	return !isPrime(num)
}

func primeCheck(num int, divisor int) bool {
	if (divisor * divisor) > num { // a prime is any number that is divisible only by itself and 1
		return true
	}
	if num%divisor == 0 {
		return false
	}
	return compositeCheck(num, divisor+1)
}
func compositeCheck(num int, divisor int) bool {
	if divisor*divisor > num {
		return false
	}
	if num%divisor == 0 {
		return true
	}
	return primeCheck(num, divisor+1)
}
