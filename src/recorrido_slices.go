package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var text_reverse string

	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		text_reverse += string(text[i])
	}

	if text == text_reverse {
		fmt.Println("The text is palindrome")
	} else {
		fmt.Println("The text is not palindrome")
	}
}

func main() {
	slice := []string{"Hi", "What", "are", "you", "doing"}

	for index, value := range slice {
		fmt.Println(index, value)
	}

	isPalindrome("Ama")

}
