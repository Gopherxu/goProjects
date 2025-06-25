package main

import (
	"fmt"
	"net/http"
	"remiges/handlers"
	"remiges/models"
)

func main() {
	// Initialize the database connection
	models.InitDB()
	defer models.DB.Close()

	// Create a new router
	// r := mux.NewRouter()

	// Define the routes
	http.HandleFunc("/employees", handlers.GetEmployees)
	http.HandleFunc("/employee/add", handlers.InsertEmployee)
	http.HandleFunc("/employee/update/{empid}", handlers.UpdateEmployee)
	http.HandleFunc("/employee/delete/{empid}", handlers.DeleteEmployee)
	http.HandleFunc("/employee/details/{empid}", handlers.GetEmployeeDetail)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Go Lang Server Starting...in 8080\n")
	fmt.Printf("http://localhost:8080\n")
	http.ListenAndServe(port, nil)
}
