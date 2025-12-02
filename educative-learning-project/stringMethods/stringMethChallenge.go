package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	TASK : Solve using the following knowledge

	Test your understanding of strings in Go by implementing various string utility functions to enhance an authentication system.
	These functions will help manipulate and analyze user IDs and emails for security purposes.
	Your goal is to correctly implement each function according to its specification.

	Below is a list of functions with a brief description that need to be implemented:

		IdentifyPrefixPostfix: This function should identify internal identifiers in IDs or emails.
		It should return a boolean indicating whether the provided string contains the specified prefix or postfix.
		For example, IdentifyPrefixPostfix(".io", "evangeline@educative.io") should return true as the email contains “.io” identifier.

		ContainsEducative: This function should check if an email contains the word “educative” and return a boolean indicating its presence.

		MaskUserName: This function should replace the user’s name from an email with asterisks (*) maintaining the first and last characters intact.
		For example, “evangeline@educative.io” should be transformed into “e********e@educative.io”.

		IndexOfAtSymbol: This function should identify the index of the @ symbol in an email and return the index as an integer.
		For example, for “evangeline@educative.io”, the function should return 10.

		TrimAndSplitUserID: This function should trim and split the number from a user ID of the “UID-0000” format.
		It should return the extracted number as a string. For example, “UID-0000” should be transformed into “0000”.

		ConvertStringToInt: This function should convert a string representation of a number to an integer.
		It should return the converted integer.

*/

func IdentifyPrefixPostfix(userID, email string) bool {
	// Implement this function
	if strings.HasPrefix(email, userID) || strings.HasSuffix(email, userID) {
		return true
	} else {
		return false
	}
}

func ContainsEducative(email string) bool {
	// Implement this function
	if strings.Contains(email, "educative") {
		return true
	} else {
		return false
	}
}

func MaskUserName(email string) string {
	// Implement this function
	emailSplit := strings.Split(email, "@")
	username := emailSplit[0]
	emaillastPart := emailSplit[1]
	firstWord := username[:1]
	asteriskWords := len(username[1 : len(username)-1])
	lastWord := username[len(username)-1:]
	censoredS := firstWord + strings.Repeat("*", asteriskWords) + lastWord + "@" + emaillastPart
	return censoredS
}

func IndexOfAtSymbol(email string) int {
	// Implement this function
	return strings.Index(email, "@")
}

func TrimAndSplitUserID(userID string) string {
	// Implement this function
	userID = strings.TrimSpace(userID)
	splitUserID := strings.Split(userID, "-")[1]
	return splitUserID
}

func ConvertStringToInt(str string) int {
	// Implement this function
	val, _ := strconv.Atoi(str)
	return val
}

func CheckChallenge() {
	// Test your functions here
	fmt.Println(IdentifyPrefixPostfix(".io", "evangeline@educative.io")) // true
	fmt.Println(IdentifyPrefixPostfix("UID", "UID-0123"))                // true
	fmt.Println(IdentifyPrefixPostfix("UID", "evangeline@educative.io")) // false
	fmt.Println(ContainsEducative("evangeline@educative.io"))            // true
	fmt.Println(MaskUserName("evangeline@educative.io"))                 // e******e@educative.io
	fmt.Println(IndexOfAtSymbol("evangeline@educative.io"))              // 10
	fmt.Println(TrimAndSplitUserID("UID-0123"))                          // 0123
	fmt.Println(ConvertStringToInt("123"))                               // 123
}
