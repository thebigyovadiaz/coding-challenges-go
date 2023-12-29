package reversestr

import "fmt"

func reverseString(input string) string {
	chars := []rune(input)
	count := len(chars)

	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	reversedString := string(chars)
	return reversedString
}

func Reversestr() {
	originalString := "Hello, World!"
	reversedString := reverseString(originalString)
	fmt.Println("Reverse-String >>", reversedString)
}
