package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func populate_json() {
	p := Person{"John", 30, "john30@john.com"}
	err := json.NewEncoder(os.Stdout).Encode(p)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func read_json() {
	var p Person
	r := strings.NewReader(`{"Name": "John", "Age": 30, "Email": "john30@john.com"}`)
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Decoded JSON: %v\n", p)
}

func main() {
	populate_json()
	read_json()
}