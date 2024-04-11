package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var employeeService EmployeeService

func main() {
	repo := NewInMemoryEmployeeRepository()
	employeeService = NewEmployeeService(repo)

	r := gin.Default()

	r.POST("/employee/create", CreateEmployeeHandler)
	r.GET("/employee/get/:userID", GetEmployeeHandler)
	r.PUT("/employee/update/:userID", UpdateEmployeeHandler)
	r.DELETE("/employee/delete/:userID", DeleteEmployeeHandler)

	fmt.Println("Server listening on port 8080...")
	r.Run(":8080")
}
