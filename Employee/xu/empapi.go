// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strconv"
// )

// // Define a struct for Employees containing a slice of Employee structs
// type Employees struct {
// 	Employees []Employee `json:"employees"`
// }

// // Define a struct for Employee
// type Employee struct {
// 	Empid int    `json:"empid"`
// 	FName string `json:"fname"`
// 	LName string `json:"lname"`
// 	Age   int    `json:"age"`
// }

// // Define a struct for API responses
// type APIResponse struct {
// 	Status  string      `json:"status"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data,omitempty"`
// }

// // Helper function to send JSON response
// func sendJSONResponse(write http.ResponseWriter, statusCode int, response APIResponse) {
// 	write.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	write.WriteHeader(statusCode)
// 	jsonResponse, _ := json.Marshal(response)
// 	write.Write(jsonResponse)
// }

// // GET Method to read all employees
// func ReadEmployees(write http.ResponseWriter, read *http.Request) {
// 	var employees Employees = ReadEmployeesJSON()

// 	response := APIResponse{
// 		Status:  "success",
// 		Message: "Employees retrieved successfully",
// 		Data:    employees,
// 	}
// 	sendJSONResponse(write, http.StatusOK, response)
// }

// // GET Method to read a specific employee by empid
// func ReadEmployee(write http.ResponseWriter, read *http.Request) {
// 	var employees Employees = ReadEmployeesJSON()

// 	ReqEmpId := read.PathValue("empid")

// 	for _, emp := range employees.Employees {
// 		if ReqEmpId == strconv.Itoa(emp.Empid) {
// 			response := APIResponse{
// 				Status:  "success",
// 				Message: "Employee retrieved successfully",
// 				Data:    emp,
// 			}
// 			sendJSONResponse(write, http.StatusOK, response)
// 			return
// 		}
// 	}

// 	response := APIResponse{
// 		Status:  "error",
// 		Message: "Employee not found",
// 	}
// 	sendJSONResponse(write, http.StatusNotFound, response)
// }

// // Function to read employees from JSON file
// func ReadEmployeesJSON() Employees {
// 	jsonFile, err := os.Open("employee.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Successfully Opened employee.json")
// 	defer jsonFile.Close()

// 	byteValue, _ := io.ReadAll(jsonFile)

// 	var employees Employees
// 	json.Unmarshal(byteValue, &employees)

// 	return employees
// }

// // POST Method to add a new employee
// func AddEmployee(write http.ResponseWriter, read *http.Request) {
// 	var employees Employees = ReadEmployeesJSON()

// 	var newEmployee Employee
// 	err := json.NewDecoder(read.Body).Decode(&newEmployee)
// 	if err != nil {
// 		response := APIResponse{
// 			Status:  "error",
// 			Message: fmt.Sprintf("Error decoding request body: %v", err),
// 		}
// 		sendJSONResponse(write, http.StatusBadRequest, response)
// 		return
// 	}

// 	newEmployee.Empid = employees.Employees[len(employees.Employees)-1].Empid + 1
// 	employees.Employees = append(employees.Employees, newEmployee)

// 	err = WriteEmployeesJSON(employees)
// 	if err != nil {
// 		response := APIResponse{
// 			Status:  "error",
// 			Message: fmt.Sprintf("Error writing JSON: %v", err),
// 		}
// 		sendJSONResponse(write, http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := APIResponse{
// 		Status:  "success",
// 		Message: "Employee added successfully",
// 		Data:    newEmployee,
// 	}
// 	sendJSONResponse(write, http.StatusOK, response)
// }

// // PUT Method to update an existing employee by empid
// // PUT Method to update an existing employee by empid
// func UpdateEmployee(write http.ResponseWriter, read *http.Request) {
// 	var employees Employees = ReadEmployeesJSON()

// 	// Decode request body into a map to handle partial updates
// 	var update map[string]interface{}
// 	err := json.NewDecoder(read.Body).Decode(&update)
// 	if err != nil {
// 		response := APIResponse{
// 			Status:  "error",
// 			Message: fmt.Sprintf("Error decoding request body: %v", err),
// 		}
// 		sendJSONResponse(write, http.StatusBadRequest, response)
// 		return
// 	}

// 	// Extract empid from URL path parameter
// 	ReqEmpId := read.PathValue("empid")
// 	empId, _ := strconv.Atoi(ReqEmpId)

// 	// Find the employee and update the specified fields
// 	for i, emp := range employees.Employees {
// 		if emp.Empid == empId {
// 			if fname, ok := update["fname"].(string); ok {
// 				emp.FName = fname
// 			}
// 			if lname, ok := update["lname"].(string); ok {
// 				emp.LName = lname
// 			}
// 			if age, ok := update["age"].(float64); ok { // JSON numbers are float64
// 				emp.Age = int(age)
// 			}
// 			employees.Employees[i] = emp

// 			err = WriteEmployeesJSON(employees)
// 			if err != nil {
// 				response := APIResponse{
// 					Status:  "error",
// 					Message: fmt.Sprintf("Error writing JSON: %v", err),
// 				}
// 				sendJSONResponse(write, http.StatusInternalServerError, response)
// 				return
// 			}
// 			response := APIResponse{
// 				Status:  "success",
// 				Message: "Employee updated successfully",
// 				Data:    emp,
// 			}
// 			sendJSONResponse(write, http.StatusOK, response)
// 			return
// 		}
// 	}

// 	response := APIResponse{
// 		Status:  "error",
// 		Message: "Employee not found",
// 	}
// 	sendJSONResponse(write, http.StatusNotFound, response)
// }

// // DELETE Method to delete an employee by empid
// func DeleteEmployee(write http.ResponseWriter, read *http.Request) {
// 	var employees Employees = ReadEmployeesJSON()

// 	ReqEmpId := read.PathValue("empid")

// 	for i, emp := range employees.Employees {
// 		if ReqEmpId == strconv.Itoa(emp.Empid) {
// 			employees.Employees = append(employees.Employees[:i], employees.Employees[i+1:]...)
// 			break
// 		}
// 	}

// 	empId := 100
// 	for i := range employees.Employees {
// 		employees.Employees[i].Empid = empId
// 		empId++
// 	}

// 	err := WriteEmployeesJSON(employees)
// 	if err != nil {
// 		response := APIResponse{
// 			Status:  "error",
// 			Message: fmt.Sprintf("Error writing JSON: %v", err),
// 		}
// 		sendJSONResponse(write, http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := APIResponse{
// 		Status:  "success",
// 		Message: "Employee deleted successfully",
// 	}
// 	sendJSONResponse(write, http.StatusOK, response)
// }

// // Function to write employees struct to JSON file
// func WriteEmployeesJSON(employees Employees) error {
// 	jsonData, err := json.MarshalIndent(employees, "", "    ")
// 	if err != nil {
// 		return err
// 	}

// 	err = os.WriteFile("employee.json", jsonData, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func main() {
// 	http.HandleFunc("/employees", ReadEmployees)                // GET all employees
// 	http.HandleFunc("/employee/{empid}", ReadEmployee)          // GET employee by empid
// 	http.HandleFunc("/employee/add", AddEmployee)               // POST add new employee
// 	http.HandleFunc("/employee/update/{empid}", UpdateEmployee) // PUT update employee by empid
// 	http.HandleFunc("/employee/delete/{empid}", DeleteEmployee) // DELETE employee by empid

// 	port := ":8080"
// 	fmt.Printf("Go Lang Server Starting...in 8080\n")
// 	fmt.Printf("http://localhost:8080\n")
// 	http.ListenAndServe(port, nil)
// }
