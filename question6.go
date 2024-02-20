// Question 6: Count Vowels
// Write a program that counts the number of vowels in a sentence.

package main

import (
	"fmt"
	"strings"
)

func numberOfVowels(sentence string) int {
	vowels := "aeiouAEIOU"
	uniqueVowels := make(map[rune]bool)

	for _, char := range sentence {
		if strings.ContainsRune(vowels, char) {
			uniqueVowels[char] = true
		}
	}

	return len(uniqueVowels)
}

func main() {
	inputSentence := "Hello World"
	result := numberOfVowels(inputSentence)
	fmt.Printf("%d", result)
}