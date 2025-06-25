package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres password=postgres dbname=remiges sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %q", err)
	}

	fmt.Println("Successfully connected to database!")
}

type Employee struct {
	ID          int            `json:"id"`
	EmpID       string         `json:"empid"`
	FName       string         `json:"fname"`
	FullName    string         `json:"fullname"`
	DOB         string         `json:"dob"`
	DOJ         string         `json:"doj"`
	Salary      int            `json:"salary"`
	ReportsTo   int            `json:"reportsto"`
	DeptID      int            `json:"deptid"`
	RankID      int            `json:"rankid"`
	CreatedAt   string         `json:"createdat"`
	UpdatedAt   sql.NullString `json:"updatedat"`
	ClientReqID string         `json:"client_reqid"`
}

type EmployeeDetail struct {
    FName       string `json:"fname"`
    DeptName    string `json:"deptname"`
    RankDesc    string `json:"rankdesc"`
}


type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func RespondSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	response := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	RespondJSON(w, status, response)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	response := ErrorResponse{
		Status:  "error",
		Message: message,
	}
	RespondJSON(w, status, response)
}
