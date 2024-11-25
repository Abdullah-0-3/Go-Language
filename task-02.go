// Write a Go program that takes two integers as input from the user and outputs their sum, difference, product, and quotient.

package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Printf("Sum of the two numbers is: %d\n", num1+num2)
	fmt.Printf("Subtraction of the two numbers is: %d\n", num1-num2)
	fmt.Printf("Multiply of the two numbers is: %d\n", num1*num2)
	fmt.Printf("Division of the two numbers is: %d\n", num1/num2)
}