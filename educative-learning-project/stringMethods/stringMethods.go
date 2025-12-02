package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
2 Package methods:
1. strings
2. strconv
*/

/*
strings package methods:
============================================
Has a Prefix - HasPrefix
Has a Suffix - HasSuffix

Contains substring? - Contains
Index of substring - Index, LastIndex (last occurrence)
Replace substring - Replace
Occurrences of substring - Count

Repeat string - Repeat
Change Case of string - ToUpper, ToLower

Trim a string - TrimSpace (only whitespace), Trim(s, str) - trim a specific string str from string s
Split a string: whitespace, separator
Joining strings over a slice

Reading from a string - skipped
============================================

strconv package methods:
============================================

Conversion to Strings
	-- Float to String
	-- Int to String


Conversion to other types
	-- String to Float
	-- String to Int

*/

func main() {
	testString := "Hello, welcome to the Go programming language. Go programming is fun!"

	// Check for prefix and suffix
	fmt.Println(`String starts with Hel?`, strings.HasPrefix(testString, "Hel")) // true
	fmt.Println(`String ends with tun?`, strings.HasSuffix(testString, "tun!"))  // false

	// Check for substring
	fmt.Println(`String contains "programming" keyword?`, strings.Contains(testString, "programming"))

	// Find index of substring
	// First occurrence
	fmt.Println(`Index of 1st instance of "programming" keyword: `, strings.Index(testString, "programming"))

	// Last occurrence
	fmt.Println(`Index of last instance of "programming" keyword: `, strings.LastIndex(testString, "programming"))

	// Replace all occurrences of a substring
	newString := strings.Replace(testString, "programming", "stupid", -1)

	fmt.Printf("Replacing \"programming\" keyword with \"stupid:\"\n%s\n", newString)

	// Count occurrences of substring
	countMarker := strings.Count(testString, "programming")

	fmt.Printf("\"Programming\" occurs in here %d times\n", countMarker)

	// Printing actual string
	fmt.Printf("\n\n%s\n", testString)

	// Repeating copies of a string
	var originS string = "Ha "
	var newS = strings.Repeat(originS, 3)

	fmt.Printf("Repeated string is\n%s\n", newS)

	/*
		Change case of string to only uppercase for each word:
			use trimming a string
			separate string on whitespace
			change case to both lower and upper
			join strings over a slice
	*/
	var mainS string = "         to liVe Is tO suFFer"
	// trim whitespace from string ends
	mainS = strings.TrimSpace(mainS)
	// make a list of the strings separated over whitespaces
	sepS := strings.Fields(mainS)

	fmt.Println("Printing each members of the string")

	for i := 0; i < len(sepS); i++ {
		sepS[i] = strings.ToUpper(sepS[i][:1]) + strings.ToLower(sepS[i][1:])
		// fmt.Printf("Member Number %d:\n%s\n", i, sepS[i])
	}
	var correctS string = strings.Join(sepS, " ")
	fmt.Printf("The final corrected string is: %s\n", correctS)

	fmt.Println("Size of integer in bits here: ", strconv.IntSize, " bits")
	var intToConv int = 65
	convertS := strconv.Itoa(intToConv)
	fmt.Printf("Type of converted string: %T\n", convertS)
	fmt.Printf("Converted from Int to String: %s\n", convertS)

	CheckChallenge()
}
