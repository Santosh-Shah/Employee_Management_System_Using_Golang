package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	UserID   int
	Fullname string
	Username string
	Password string
	Phone    string
	Email    string
}

var db *sql.DB

func main() {
	// Connect to the MySQL database
	var err error
	// Update the database connection string with the correct username and password
	db, err = sql.Open("mysql", "root:@Mysql_679#@tcp(localhost:3306)/employees")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Welcome to Employee Management System!")
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add/Create Employee")
		fmt.Println("2. Get Employee")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Get All Employees")
		fmt.Println("6. Exit")

		option := readInput("Enter option: ")

		switch option {
		case "1":
			createEmployee()
		case "2":
			getEmployee()
		case "3":
			updateEmployee()
		case "4":
			deleteEmployee()
		case "5":
			getAllEmployees()
		case "6":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}

func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func createEmployee() {
	userID, _ := strconv.Atoi(readInput("Enter userID: "))
	// Check if employee with same userID already exists
	if employeeExists(userID) {
		fmt.Println("Employee with same userID already exists!")
		return
	}

	fullname := readInput("Enter fullname: ")
	username := readInput("Enter username: ")
	// Check if employee with same username already exists
	if usernameExists(username) {
		fmt.Println("Employee with same username already exists!")
		return
	}

	password := readInput("Enter password: ")
	phone := readInput("Enter phone: ")
	email := readInput("Enter email: ")
	// Check if employee with same email already exists
	if emailExists(email) {
		fmt.Println("Employee with same email already exists!")
		return
	}

	_, err := db.Exec("INSERT INTO employees (userID, fullname, username, password, phone, email) VALUES (?, ?, ?, ?, ?, ?)",
		userID, fullname, username, password, phone, email)
	if err != nil {
		fmt.Println("Error creating employee:", err)
		return
	}

	fmt.Println("Employee created successfully!")
}

func getEmployee() {
	userID, _ := strconv.Atoi(readInput("Enter userID: "))
	var (
		fullname string
		username string
		phone    string
		email    string
	)

	err := db.QueryRow("SELECT fullname, username, phone, email FROM employees WHERE userID=?", userID).Scan(&fullname, &username, &phone, &email)
	if err != nil {
		fmt.Println("Error getting employee details:", err)
		return
	}

	fmt.Println("Employee Details:")
	fmt.Printf("User ID: %d\n", userID)
	fmt.Printf("Fullname: %s\n", fullname)
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Phone: %s\n", phone)
	fmt.Printf("Email: %s\n", email)
}

func updateEmployee() {
	userID, _ := strconv.Atoi(readInput("Enter userID: "))
	// Check if employee exists
	if !employeeExists(userID) {
		fmt.Println("Employee not found!")
		return
	}

	fmt.Println("Update Employee Details:")
	fullname := readInput("Enter new fullname: ")
	username := readInput("Enter new username: ")
	password := readInput("Enter new password: ")
	phone := readInput("Enter new phone: ")
	email := readInput("Enter new email: ")

	_, err := db.Exec("UPDATE employees SET fullname=?, username=?, password=?, phone=?, email=? WHERE userID=?",
		fullname, username, password, phone, email, userID)
	if err != nil {
		fmt.Println("Error updating employee details:", err)
		return
	}

	fmt.Println("Employee details updated successfully!")
}

func deleteEmployee() {
	userID, _ := strconv.Atoi(readInput("Enter userID: "))
	// Check if employee exists
	if !employeeExists(userID) {
		fmt.Println("Employee not found!")
		return
	}

	_, err := db.Exec("DELETE FROM employees WHERE userID=?", userID)
	if err != nil {
		fmt.Println("Error deleting employee:", err)
		return
	}

	fmt.Println("Employee deleted successfully!")
}

func getAllEmployees() {
	rows, err := db.Query("SELECT userID, fullname, username, phone, email FROM employees")
	if err != nil {
		fmt.Println("Error getting employees:", err)
		return
	}
	defer rows.Close()

	fmt.Println("All Employees:")
	for rows.Next() {
		var (
			userID   int
			fullname string
			username string
			phone    string
			email    string
		)
		err := rows.Scan(&userID, &fullname, &username, &phone, &email)
		if err != nil {
			fmt.Println("Error scanning employee:", err)
			continue
		}
		fmt.Printf("User ID: %d\n", userID)
		fmt.Printf("Fullname: %s\n", fullname)
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Phone: %s\n", phone)
		fmt.Printf("Email: %s\n", email)
		fmt.Println("-------------------")
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func employeeExists(userID int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM employees WHERE userID=?", userID).Scan(&count)
	if err != nil {
		fmt.Println("Error checking if employee exists:", err)
		return false
	}
	return count > 0
}

func usernameExists(username string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM employees WHERE username=?", username).Scan(&count)
	if err != nil {
		fmt.Println("Error checking if username exists:", err)
		return false
	}
	return count > 0
}

func emailExists(email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM employees WHERE email=?", email).Scan(&count)
	if err != nil {
		fmt.Println("Error checking if email exists:", err)
		return false
	}
	return count > 0
}
