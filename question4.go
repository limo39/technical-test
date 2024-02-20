// Question 4: Capitalize Words
// Write a program that accepts a string as input, capitalizes the first letter of each word in the
// string, and then returns the result string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func capitalizeWords(input string) string {
	words := strings.Fields(input) 
	capitalizedWords := make([]string, len(words))

	for i, word := range words {
		capitalizedWords[i] = strings.Title(word)
	}

	result := strings.Join(capitalizedWords, " ")

	return result
}

func main() {
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	result := capitalizeWords(input)
	fmt.Printf("%s\n", result)
}
