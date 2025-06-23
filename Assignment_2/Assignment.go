package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct to represent the JSON response body
type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(write http.ResponseWriter, read *http.Request) {
	if read.Method != http.MethodPost {
		http.Error(write, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the query parameter
	name := read.URL.Query().Get("name")
	if name == "" {
		http.Error(write, "Name parameter is required", http.StatusBadRequest)
		return
	}

	// Create the response message
	response := HelloResponse{
		Message: fmt.Sprintf("Hello %s", name),
	}

	// Set the response header to application/json
	write.Header().Set("Content-Type", "application/json")

	// Encode the response as JSON and send it
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		http.Error(write, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	port := ":8090"
	fmt.Println("Server is running at http://localhost:8090/hello")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
