// Question 3: Power of Two
// Write a program that takes an integer as input and returns true if the input is a power of two.

package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	// A power of two has only one bit set in its binary representation.
	// Using bitwise AND operation to check if there is only one set bit.
	return (n & (n - 1)) == 0
}

func main() {
	var input int
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	result := isPowerOfTwo(input)
	fmt.Printf("%t",  result)
}
