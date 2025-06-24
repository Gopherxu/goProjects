package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct to handle the incoming JSON request
type CalcRequest struct {
	Numbers   []float64 `json:"numbers"`
	Operation string    `json:"operation"`
}

// Struct to handle the outgoing JSON response
type CalcResponse struct {
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}

// Handler function for /mycalc
func MyCalcHandler(write http.ResponseWriter, read *http.Request) {
	var calcReq CalcRequest

	// Read the request body
	body, err := io.ReadAll(read.Body)
	if err != nil {
		http.Error(write, "Invalid request", http.StatusBadRequest)
		return
	}

	// Parse the JSON request
	err = json.Unmarshal(body, &calcReq)
	if err != nil {
		http.Error(write, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Perform the requested operation
	result, err := PerformOperation(calcReq.Numbers, calcReq.Operation)
	if err != nil {
		http.Error(write, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the response
	calcResp := CalcResponse{
		Operation: calcReq.Operation,
		Result:    result,
	}

	// Convert response to JSON
	jsonResp, err := json.Marshal(calcResp)
	if err != nil {
		http.Error(write, "Error generating response", http.StatusInternalServerError)
		return
	}

	// Set headers and write response
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusOK)
	write.Write(jsonResp)
}

// Function to perform the requested operation
func PerformOperation(numbers []float64, operation string) (float64, error) {
	switch operation {
	case "mean":
		return Mean(numbers), nil
	case "min":
		return Min(numbers), nil
	case "max":
		return Max(numbers), nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", operation)
	}
}

// Function to calculate the mean of a list of numbers
func Mean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// Function to find the minimum value in a list of numbers
func Min(numbers []float64) float64 {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

// Function to find the maximum value in a list of numbers
func Max(numbers []float64) float64 {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	http.HandleFunc("/mycalc", MyCalcHandler)
	port := ":8080"
	fmt.Printf("Go Lang Server Starting...on port 8080\n")
	fmt.Printf("http://localhost:8080/mycalc      ...!!!")
	http.ListenAndServe(port, nil)
}
