// Write a Go program that takes a string input from the user and counts the number of vowels and consonants in that string. For this task, consider the English vowels to be A, E, I, O, and U (both uppercase and lowercase).

// Your program should ignore any non-alphabetic characters. Here’s how you might approach this:

// Read input from the user.
// Initialize counters for vowels and consonants.
// Loop through each character in the string to determine if it is a vowel, consonant, or neither.
// Print the results.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var count_vowels int = 0
	var count_constants int = 0

	fmt.Print("Enter a string: ")
	fmt.Scanln(&input) 

	input = strings.ToLower(input)

	for i := 0; i < len(input); i++ {
		char := input[i]
		if strings.Contains("aeiou", string(char)) {
			count_vowels += 1
		} else if char >= 'a' && char <= 'z' {
			count_constants += 1
		}
	}

	fmt.Printf("Vowels: %d\n", count_vowels)
	fmt.Printf("Consonants: %d\n", count_constants)
}
