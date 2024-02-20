// Question 1: FizzBuzz
// Write a program that prints the numbers from 1 to 100. For multiples of 3, print "Fizz"; for
// multiples of 5, print "Buzz"; and for numbers that are multiples of both 3 and 5, print "FizzBuzz".

package main

import (
	"fmt"
)

func buzzFizz(){
	n := 1
	for n<100{
		if n%3 == 0 && n%5 ==0 {
        fmt.Println("FizzBuzz")
    } else if n%3 == 0 {
        fmt.Println("Buzz")
    } else if n%5 == 0 {
        fmt.Println("Fizz")
    } else {
        fmt.Println(n)
    } 
   
		n++
	}
}

func main(){
	buzzFizz()
}