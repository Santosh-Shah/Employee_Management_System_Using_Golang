package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	UserID   int
	Fullname string
	Username string
	Password string
	Phone    string
	Email    string
}

var employees []*Employee

func main() {
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
	return input[:len(input)-1]
}

func createEmployee() {
	userID, _ := strconv.Atoi(readInput("Enter userID: "))
	// Check if employee with same userID already exists
	for _, emp := range employees {
		if emp.UserID == userID {
			fmt.Println("Employee with same userID already exists!")
			return
		}
	}

	fullname := readInput("Enter fullname: ")
	username := readInput("Enter username: ")
	password := readInput("Enter password: ")
	phone := readInput("Enter phone: ")
	email := readInput("Enter email: ")

	employee := &Employee{
		UserID:   userID,
		Fullname: fullname,
		Username: username,
		Password: password,
		Phone:    phone,
		Email:    email,
	}

	employees = append(employees, employee)
	fmt.Println("Employee created successfully!")
}

func getEmployee() {
	userID, _ := strconv.Atoi(readInput("Enter userID: "))

	for _, emp := range employees {
		if emp.UserID == userID {
			fmt.Println("Employee Details:")
			fmt.Printf("User ID: %d\n", emp.UserID)
			fmt.Printf("Fullname: %s\n", emp.Fullname)
			fmt.Printf("Username: %s\n", emp.Username)
			fmt.Printf("Phone: %s\n", emp.Phone)
			fmt.Printf("Email: %s\n", emp.Email)
			return
		}
	}

	fmt.Println("Employee not found!")
}

func updateEmployee() {
	userID, _ := strconv.Atoi(readInput("Enter userID: "))

	for _, emp := range employees {
		if emp.UserID == userID {
			fmt.Println("Update Employee Details:")
			emp.Fullname = readInput("Enter new fullname: ")
			emp.Username = readInput("Enter new username: ")
			emp.Password = readInput("Enter new password: ")
			emp.Phone = readInput("Enter new phone: ")
			emp.Email = readInput("Enter new email: ")
			fmt.Println("Employee details updated successfully!")
			return
		}
	}

	fmt.Println("Employee not found!")
}

func deleteEmployee() {
	userID, _ := strconv.Atoi(readInput("Enter userID: "))

	for i, emp := range employees {
		if emp.UserID == userID {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted successfully!")
			return
		}
	}

	fmt.Println("Employee not found!")
}

func getAllEmployees() {
	fmt.Println("All Employees:")
	for _, emp := range employees {
		fmt.Printf("User ID: %d\n", emp.UserID)
		fmt.Printf("Fullname: %s\n", emp.Fullname)
		fmt.Printf("Username: %s\n", emp.Username)
		fmt.Printf("Phone: %s\n", emp.Phone)
		fmt.Printf("Email: %s\n", emp.Email)
		fmt.Println("-------------------")
	}
}
