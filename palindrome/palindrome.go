package palindrome

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	// Convert the string to lowercase and remove spaces
	cleanedInput := strings.ReplaceAll(strings.ToLower(input), " ", "")

	// Convert the string to a rune slice for easy comparison
	chars := []rune(cleanedInput)

	// Get the length of the rune slice
	length := len(chars)

	// Check if the string is a palindrome
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}

	return true
}

func Palindrome() {
	testString := "A man a plan a canal Panama"
	result := isPalindrome(testString)
	fmt.Printf("Palindrome >> Is \"%s\" a palindrome? %t\n", testString, result)
}
