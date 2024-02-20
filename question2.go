// Question 2: Fibonacci Sequence
// Write a program to generate the Fibonacci sequence up to 100.

package main

import(
	"fmt"
)

func main(){
	fib := fibGen()
	for i := 0; i<11; i++{
		fmt.Print(fib(), " ")
	}
}

func fibGen() func() int{
	f1 := 0
	f2 := 1
	return func() int{f2, f1 =(f1 + f2), f2
	return f1
	}
}