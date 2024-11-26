// Write a Go program that functions as a simple calculator. Your program should:

// Prompt the user to enter two numbers.
// Prompt the user to choose an operation (addition, subtraction, multiplication, or division).
// Perform the chosen operation on the numbers and print the result.
// Here’s a quick guide to get you started:

// Use a loop to allow multiple calculations until the user decides to exit.
// Handle division by zero in the case of division operation.
// Use functions for different operations to keep your code organized.

package main

import (
	"fmt"
)

func main() {
	calculator_function()
}

func calculator_function() {
	for {
		var operator string // Declare operator here to fix the error
		if operator == "0" {
			break
		} else {
			var x, y int

			fmt.Print("Enter first number: ")
			fmt.Scanln(&x)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&y)
			fmt.Print("Enter operator (+, -, *, /, %): 0(exit) ")
			fmt.Scanln(&operator)

			result := calculator(x, y, operator)
			fmt.Printf("Result: %d\n", result)
		}
	}
}

func calculator(x, y int, operator string) int {
	switch operator {
	case "+":
		return add(x, y)
	case "-":
		return sub(x, y)
	case "*":
		return mul(x, y)
	case "/":
		return div(x, y)
	case "%":
		return mod(x, y)
	default:
		fmt.Println("Invalid operator")
	}
	return 0
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func mod(x, y int) int {
	return x % y
}