package strmanipulation

import (
	"fmt"
	"strings"
)

// ReverseWords reverses the order of words in a sentence.
func ReverseWords(sentence string) string {
	words := strings.Fields(sentence)
	reversedWords := make([]string, len(words))

	for i, j := len(words)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversedWords[j] = words[i]
	}

	return strings.Join(reversedWords, " ")
}

// FirstNonRepeatedChar finds the first non-repeated character in a string.
func FirstNonRepeatedChar(input string) rune {
	charCount := make(map[rune]int)

	for _, char := range input {
		charCount[char]++
	}

	for _, char := range input {
		if charCount[char] == 1 {
			return char
		}
	}

	return 0 // If no non-repeated character is found
}

func StrManipulation() {
	// Example usage for reversing words in a sentence:
	sentence := "Hello, World! Greetings from Go."
	reversedSentence := ReverseWords(sentence)
	fmt.Println("Reversed Sentence:")
	fmt.Println(reversedSentence)

	// Example usage for finding the first non-repeated character:
	inputString := "abacabad"
	firstNonRepeated := FirstNonRepeatedChar(inputString)
	fmt.Printf("First non-repeated character in \"%s\": %c\n", inputString, firstNonRepeated)
}
