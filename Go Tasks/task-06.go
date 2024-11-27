// Write a Go program that generates the Fibonacci series up to a user-specified number of terms. 

// The Fibonacci series starts with 0 and 1, and then each subsequent number is the sum of the previous two. For example, the series looks like this:
// 0, 1, 1, 2, 3, 5, 8, 13, ...

// Your program should:

// Prompt the user for the number of terms they want in the Fibonacci series.
// Use a loop to calculate and print the Fibonacci numbers up to that number of terms.
// Handle cases where the input is less than or equal to zero (for example, you can print a message indicating that the number of terms must be positive).

package main

import (
	"fmt"
)



func fibonacci() {
	var n int

	fmt.Println("Enter the number of terms you want in the fibonacci series: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Println(fib(i))
	}
}


func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fibonacci()
}