// models.go
package main

// Define a struct for Employees containing a slice of Employee structs
type Employees struct {
	Employees []Employee `json:"employees"`
}

// Define a struct for Employee
type Employee struct {
	Empid    int    `json:"empid"`
	FName    string `json:"fname"`
	LName    string `json:"lname"`
	Age      int    `json:"age"`
	FullName string `json:"fullname"`
}

// Define a struct for API responses
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
