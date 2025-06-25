package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int     `xml:"id"`
	Name     string  `xml:"name"`
	Position string  `xml:"position"`
	Salary   float64 `xml:"salary"`
	HireDate string  `xml:"hire_date"`
}

type Employees struct {
	XMLName   xml.Name   `xml:"employees"`
	Employees []Employee `xml:"employee"`
}

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres password=postgres dbname=employees sslmode=disable"
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

func main() {
	InitDB()
	defer DB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/employees/xml", GenerateXML).Methods("GET")
	r.HandleFunc("/download/xml", DownloadXMLFile).Methods("GET")
	http.Handle("/", r)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fetchEmployees() ([]Employee, error) {
	rows, err := DB.Query("SELECT id, name, position, salary, hire_date FROM employees")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary, &emp.HireDate)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}
		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %v", err)
	}

	return employees, nil
}

func GenerateXML(w http.ResponseWriter, r *http.Request) {
	employees, err := fetchEmployees()
	if err != nil {
		http.Error(w, "Unable to fetch employees", http.StatusInternalServerError)
		return
	}

	employeesXML := Employees{Employees: employees}

	fileName := "employees.xml"
	file, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Error creating XML file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	xmlData, err := xml.MarshalIndent(employeesXML, "", "  ")
	if err != nil {
		http.Error(w, "Error generating XML", http.StatusInternalServerError)
		return
	}

	_, err = file.Write(xmlData)
	if err != nil {
		http.Error(w, "Error writing to XML file", http.StatusInternalServerError)
		return
	}

	file.Write([]byte("\n"))

	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(fmt.Sprintf("XML file created successfully: %s", fileName)))
}

func DownloadXMLFile(w http.ResponseWriter, r *http.Request) {
	filePath := "employees.xml"
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found.", 404)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("Content-Disposition", "attachment;filename=employees.xml")

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error serving XML file", http.StatusInternalServerError)
		return
	}
}
