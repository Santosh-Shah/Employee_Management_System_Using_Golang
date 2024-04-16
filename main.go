package main

//import (
//	"database/sql"
//	"fmt"
//	"log"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//type Employee struct {
//	UserID   int
//	Fullname string
//	Username string
//	Password string
//	Phone    string
//	Email    string
//}
//
//func main() {
//	// Connect to the MySQL database
//	db, err := sql.Open("mysql", "root:@Mysql_679#@tcp(localhost:3306)/employees")
//	if err != nil {
//		log.Fatal("Error connecting to database:", err)
//	}
//	defer db.Close()
//
//	// Prompt user for employee details
//	var employee Employee
//	fmt.Println("Enter employee details:")
//	fmt.Print("UserID: ")
//	fmt.Scan(&employee.UserID)
//	fmt.Print("Fullname: ")
//	fmt.Scan(&employee.Fullname)
//	fmt.Print("Username: ")
//	fmt.Scan(&employee.Username)
//	fmt.Print("Password: ")
//	fmt.Scan(&employee.Password)
//	fmt.Print("Phone: ")
//	fmt.Scan(&employee.Phone)
//	fmt.Print("Email: ")
//	fmt.Scan(&employee.Email)
//
//	// Insert employee data into the database
//	_, err = db.Exec("INSERT INTO employees (userID, fullname, username, password, phone, email) VALUES (?, ?, ?, ?, ?, ?)",
//		employee.UserID, employee.Fullname, employee.Username, employee.Password, employee.Phone, employee.Email)
//	if err != nil {
//		log.Fatal("Error inserting employee data:", err)
//	}
//
//	fmt.Println("Employee data inserted successfully!")
//
//	// Query the employees table
//	rows, err := db.Query("SELECT * FROM employees")
//	if err != nil {
//		log.Fatal("Error querying database:", err)
//	}
//	defer rows.Close()
//
//	// Iterate over the rows
//	fmt.Println("\nEmployees:")
//	for rows.Next() {
//		var employee Employee
//		if err := rows.Scan(&employee.UserID, &employee.Fullname, &employee.Username, &employee.Password, &employee.Phone, &employee.Email); err != nil {
//			log.Fatal("Error scanning row:", err)
//		}
//		fmt.Printf("User ID: %d, Fullname: %s, Username: %s, Phone: %s, Email: %s\n", employee.UserID, employee.Fullname, employee.Username, employee.Phone, employee.Email)
//	}
//	if err := rows.Err(); err != nil {
//		log.Fatal("Error iterating over rows:", err)
//	}
//}
