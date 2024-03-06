// TASK 3
// Check whether a given string is a palindrome or not. A palindrome is a word,
//  phrase, number, or other sequence of characters that reads the same forward
//  and backward (ignoring spaces, punctuation, and capitalization).

package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(s string) bool {
	// Remove spaces and convert to lowercase
	s = strings.ToLower(s)
	s = removeNonAlphaNumeric(s)

	// Check if the string is a palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func removeNonAlphaNumeric(s string) string {
	var builder strings.Builder
	for _, char := range s {
		if isAlphaNumeric(char) {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func isAlphaNumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}

func main() {
	var input string
	fmt.Print("Enter a string for Palindrome Checking ")
	fmt.Scanln(&input)

	if checkPalindrome(input) {
		fmt.Println("The string is a palindrome")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
