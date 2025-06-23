package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {

	// Print Hello World
	fmt.Fprintf(writer, "Hello World")
}

func main() {
	http.HandleFunc("/hello", hello)
	port := ":8089"
	fmt.Println("Starting Server at http://localhost:8089/hello  !!!.....")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
