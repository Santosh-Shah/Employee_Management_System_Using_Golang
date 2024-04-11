package main

import "fmt"

type EmployeeRepository interface {
	CreateEmployee(employee *Employee) error
	GetEmployee(userID int) (*Employee, error)
	UpdateEmployee(userID int, employee *Employee) error
	DeleteEmployee(userID int) error
}

type InMemoryEmployeeRepository struct {
	employees map[int]*Employee
}

func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		employees: make(map[int]*Employee),
	}
}

func (r *InMemoryEmployeeRepository) CreateEmployee(employee *Employee) error {
	r.employees[employee.UserID] = employee
	return nil
}

func (r *InMemoryEmployeeRepository) GetEmployee(userID int) (*Employee, error) {
	employee, ok := r.employees[userID]
	if !ok {
		return nil, fmt.Errorf("employee with userID %d not found", userID)
	}
	return employee, nil
}

func (r *InMemoryEmployeeRepository) UpdateEmployee(userID int, employee *Employee) error {
	r.employees[userID] = employee
	return nil
}

func (r *InMemoryEmployeeRepository) DeleteEmployee(userID int) error {
	delete(r.employees, userID)
	return nil
}
