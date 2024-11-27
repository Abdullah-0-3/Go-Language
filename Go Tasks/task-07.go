package main

import (
	"fmt"
	"os"
)

func fileWordsCount(file *os.File) int {
	var words int
	var word string

	for {
		_, err := fmt.Fscan(file, &word)
		if err != nil {
			break
		}
		words++
	}
	return words
}

func getFileName() (string, error) {
	fmt.Print("Enter the file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	if _, err := os.Stat(fileName); err != nil {
		return "", fmt.Errorf("error checking file: %v", err)
	}
	return fileName, nil
}

func main() {
	fileName, err := getFileName()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Word count:", fileWordsCount(file))
}