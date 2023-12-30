package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	DateOfBirth    time.Time   `json:"date_of_birth,omitempty"`
	Email          string      `json:"email"`
	Department     string      `json:"department"`
	Address        Address     `json:"address"`
	Phone          json.Number `json:"phone"`
	PpsNumber      string      `json:"pps_number,omitempty"`
	DocumentType   string      `json:"document_type,omitempty"`
	DocumentNumber string      `json:"document_number,omitempty"`
}

type Address struct {
	HouseNameOrNumber string `json:"house_name_or_number,omitempty"`
	Line1             string `json:"line1"`
	Line2             string `json:"line2,omitempty"`
	Town              string `json:"town,omitempty"`
	City              string `json:"city"`
	Country           string `json:"country"`
	PostalCode        string `json:"postal_code,omitempty"`
}

type DB struct {
	data map[string]map[string]string
}

func New(dir string, options interface{}) (*DB, error) {
	return &DB{data: make(map[string]map[string]string)}, nil
}

func (db *DB) Write(collection, key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if _, ok := db.data[collection]; !ok {
		db.data[collection] = make(map[string]string)
	}

	db.data[collection][key] = string(jsonData)
	return nil
}

func (db *DB) ReadAll(collection string) ([]string, error) {
	var records []string
	for _, record := range db.data[collection] {
		records = append(records, record)
	}
	return records, nil
}

func (db *DB) GetAllEmployees() {
	records, err := db.ReadAll("users")
	if err != nil {
		log.Fatalf("Error reading from the database: %s", err)
	}
	fmt.Println(records)
}

func (db *DB) GetEmployeeById(id string) (string, error) {
	record, ok := db.data["users"][id]
	if !ok {
		return "", fmt.Errorf("employee with ID %s not found", id)
	}
	return record, nil
}

func (db *DB) CreateEmployee(id string, employee User) error {
	if _, ok := db.data["users"][id]; ok {
		return fmt.Errorf("employee with ID %s already exists", id)
	}

	return db.Write("users", id, employee)
}

func (db *DB) UpdateEmployee(id string, updatedEmployee User) error {
	if _, ok := db.data["users"][id]; !ok {
		return fmt.Errorf("employee with ID %s not found", id)
	}

	return db.Write("users", id, updatedEmployee)
}

func (db *DB) DeleteEmployee(id string) error {
	if _, ok := db.data["users"][id]; !ok {
		return fmt.Errorf("employee with ID %s not found", id)
	}

	delete(db.data["users"], id)
	return nil
}

func main() {
	dir := "./"
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := map[string]User{
		"1": {"Homer", "Simpson", time.Date(1956, time.May, 12, 0, 0, 0, 0, time.UTC), "homer@example.com", "Nuclear Power Plant", Address{"742", "Evergreen Terrace", "Henderson", "Springfield", "Oregon", "USA", "80085"}, "123456789", "1111111H", "Driver's License", "DL987654"},
		"2": {"Marge", "Simpson", time.Date(1956, time.November, 1, 0, 0, 0, 0, time.UTC), "marge@example.com", "Homemaker", Address{"742", "Evergreen Terrace", "Henderson", "Springfield", "Oregon", "USA", "80085"}, "987654321", "2222222M", "Library Card", "LC123456"},
		// Add more employees as needed
	}

	for id, value := range employees {
		err := db.CreateEmployee(id, value)
		if err != nil {
			log.Fatalf("Error creating employee: %s", err)
		}
	}

	db.GetAllEmployees()

	// Example of getting an employee by ID
	employeeID := "1"
	employeeRecord, err := db.GetEmployeeById(employeeID)
	if err != nil {
		log.Fatalf("Error getting employee by ID: %s", err)
	}
	fmt.Printf("Employee with ID %s: %s\n", employeeID, employeeRecord)

	// Example of updating an employee
	updatedEmployee := User{
		FirstName: "Updated",
		LastName:  "Employee",
		// Update other fields as needed
	}
	err = db.UpdateEmployee(employeeID, updatedEmployee)
	if err != nil {
		log.Fatalf("Error updating employee: %s", err)
	}

	db.GetAllEmployees()

	// Example of deleting an employee
	err = db.DeleteEmployee(employeeID)
	if err != nil {
		log.Fatalf("Error deleting employee: %s", err)
	}

	db.GetAllEmployees()
}
