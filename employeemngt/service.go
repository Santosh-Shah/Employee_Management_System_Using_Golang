package main

type EmployeeService interface {
	CreateEmployee(employee *Employee) error
	GetEmployee(userID int) (*Employee, error)
	UpdateEmployee(userID int, employee *Employee) error
	DeleteEmployee(userID int) error
}

type EmployeeServiceImpl struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{repo: repo}
}

func (s *EmployeeServiceImpl) CreateEmployee(employee *Employee) error {
	return s.repo.CreateEmployee(employee)
}

func (s *EmployeeServiceImpl) GetEmployee(userID int) (*Employee, error) {
	return s.repo.GetEmployee(userID)
}

func (s *EmployeeServiceImpl) UpdateEmployee(userID int, employee *Employee) error {
	return s.repo.UpdateEmployee(userID, employee)
}

func (s *EmployeeServiceImpl) DeleteEmployee(userID int) error {
	return s.repo.DeleteEmployee(userID)
}
