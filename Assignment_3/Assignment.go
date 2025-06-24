package main

import (
	"fmt"      // Importing the fmt package for formatted I/O
	"net/http" // Importing the net/http package to create HTTP servers and handle HTTP requests
	"strconv"  // Importing the strconv package to convert strings to numeric types
)

// Function to handle HTTP requests for the /mysum endpoint
func mysum(write http.ResponseWriter, read *http.Request) {
	// Retrieve the 'operator', 'num1', and 'num2' query parameters from the URL
	operator := read.URL.Query().Get("operator")
	num1 := read.URL.Query().Get("num1")
	num2 := read.URL.Query().Get("num2")

	// Convert 'num1' from string to float64
	number1, err1 := strconv.ParseFloat(num1, 64)
	if err1 != nil {
		// If conversion fails, send an HTTP error response
		http.Error(write, "Invalid Number", http.StatusBadRequest)
		return
	}

	// Convert 'num2' from string to float64
	number2, err2 := strconv.ParseFloat(num2, 64)
	if err2 != nil {
		// If conversion fails, send an HTTP error response
		http.Error(write, "Invalid 2nd Number", http.StatusBadRequest)
		return
	}

	// Declare variables for the result and error
	var result float64
	var err error

	// Perform the operation based on the value of 'operator'
	switch operator {
	case "Add":
		result = number1 + number2 // Addition
	case "Substract":
		result = number1 - number2 // Subtraction
	case "Multiply":
		result = number1 * number2 // Multiplication
	case "Divide":
		if number2 == 0 {
			// If the second number is zero, division is not allowed
			err = fmt.Errorf("Can't Divide By Zeroo")
		} else {
			result = number1 / number2 // Division
		}
	default:
		// If the operator is not recognized, set an error
		err = fmt.Errorf("Unsupported Operation : %v !...", operator)
	}

	// If an error occurred, send an HTTP error response
	if err != nil {
		http.Error(write, err.Error(), http.StatusBadRequest)
	} else {
		// Otherwise, send the result back to the client
		fmt.Fprintf(write, "Result of %v %v %v = %v", num1, operator, num2, result)
	}
}

func main() {
	// Set up the /mysum endpoint to be handled by the mysum function
	http.HandleFunc("/mysum", mysum)
	port := ":8090" // Define the port on which the server will listen
	fmt.Println("Server is Running at http://localhost:8090/mysum?operator=+&num1=10&num2=5   ....!!")

	// Start the HTTP server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		// If the server fails to start, print an error message
		fmt.Printf("Failed to Start Server !......")
	}
}
