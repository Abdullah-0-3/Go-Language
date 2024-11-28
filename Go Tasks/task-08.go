package main

import (
	"fmt"
	"net/http"
)

func http_server_home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func http_server_bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye World")
}

func main() {
	http.HandleFunc("/", http_server_home)
	http.HandleFunc("/bye", http_server_bye)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	} 
}