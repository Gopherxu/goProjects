package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"remiges/models"
	"time"

	"github.com/lib/pq"
)

func InsertEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		models.RespondError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		models.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate required fields
	if emp.EmpID == "" || emp.FName == "" || emp.DOB == "" || emp.DOJ == "" || emp.ClientReqID == "" {
		models.RespondError(w, http.StatusBadRequest, "All fields are required")
		return
	}

	// Check for null values in optional fields
	if emp.FullName == "" {
		emp.FullName = emp.FName // Assigning first name to full name if full name is not provided
	}

	// Prepare the SQL query
	query := `INSERT INTO employee (empid, fname, fullname, dob, doj, salary, reportsto, deptid, rankid, createdat, client_reqid) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, createdat`
	err = models.DB.QueryRow(query, emp.EmpID, emp.FName, emp.FullName, emp.DOB, emp.DOJ, emp.Salary, emp.ReportsTo, emp.DeptID, emp.RankID, time.Now(), emp.ClientReqID).
		Scan(&emp.ID, &emp.CreatedAt)
	if err != nil {
		// Check for PostgreSQL unique violation error
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code.Name() == "unique_violation" {
			models.RespondError(w, http.StatusConflict, "Employee with the same empid already exists")
			return
		}

		// Check for other SQL errors
		if err == sql.ErrNoRows {
			models.RespondError(w, http.StatusNotFound, "Record not found")
			return
		}

		// Handle any other errors
		models.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	models.RespondSuccess(w, http.StatusCreated, "Employee inserted successfully", emp)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		models.RespondError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	rows, err := models.DB.Query("SELECT id, empid, fname, fullname, dob, doj, salary, reportsto, deptid, rankid, createdat, updatedat, client_reqid FROM employee")
	if err != nil {
		models.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	employees := []models.Employee{}
	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(&emp.ID, &emp.EmpID, &emp.FName, &emp.FullName, &emp.DOB, &emp.DOJ, &emp.Salary, &emp.ReportsTo, &emp.DeptID, &emp.RankID, &emp.CreatedAt, &emp.UpdatedAt, &emp.ClientReqID)
		if err != nil {
			models.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		employees = append(employees, emp)
	}

	models.RespondSuccess(w, http.StatusOK, "Employees retrieved successfully", employees)
}

// update employee Data
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		models.RespondError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	// Parse empid from Path parameters
	empid := r.PathValue("empid")

	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE employee SET fname=$1, fullname=$2, dob=$3, doj=$4, salary=$5, reportsto=$6, deptid=$7, rankid=$8, updatedat=NOW(), client_reqid=$9 
				WHERE empid=$10 RETURNING id, updatedat`
	err = models.DB.QueryRow(query, emp.FName, emp.FullName, emp.DOB, emp.DOJ, emp.Salary, emp.ReportsTo, emp.DeptID, emp.RankID, emp.ClientReqID, empid).Scan(&emp.ID, &emp.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Employee not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	models.RespondSuccess(w, http.StatusOK, "Employee updated successfully", emp)
}

// DeleteEmployee deletes an employee record by empid
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		models.RespondError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	// Parse empid from query parameters
	empid := r.PathValue("empid")

	// Prepare the SQL query
	query := "DELETE FROM employee WHERE empid = $1"

	// Execute the delete query
	result, err := models.DB.Exec(query, empid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// Respond with success message
	models.RespondSuccess(w, http.StatusOK, "Employee deleted successfully", nil)
}

func GetEmployeeDetail(write http.ResponseWriter, read *http.Request) {

	if read.Method != http.MethodGet {
		models.RespondError(write, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	// Get the employee ID from the query parameters
	empid :=
		read.PathValue("empid")
	if empid == "" {
		models.RespondError(write, http.StatusBadRequest, "empid query parameter is required")
		return
	}

	// Define the query to fetch employee details
	query := `
	SELECT e.fname, d.deptname, r.rankdesc
	FROM employee e
	JOIN departments d ON e.deptid = d.id
	JOIN ranks r ON e.rankid = r.id
	WHERE e.empid = $1`

	// Execute the query
	var empDetail models.EmployeeDetail
	err := models.DB.QueryRow(query, empid).Scan(&empDetail.FName, &empDetail.DeptName, &empDetail.RankDesc)
	if err != nil {
		if err == sql.ErrNoRows {
			models.RespondError(write, http.StatusNotFound, "Employee not found")
		} else {
			models.RespondError(write, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Respond with the employee details
	models.RespondSuccess(write, http.StatusOK, "Employee details retrieved successfully", empDetail)
}
