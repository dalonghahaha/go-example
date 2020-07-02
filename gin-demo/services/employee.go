package services

import (
	"gin-demo/models/hr"
)

type EmployeeService struct {
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

func (s *EmployeeService) GetEmployee(id int64) *hr.YwEmployee {
	employee := new(hr.YwEmployee)
	err := employee.GeByID(id)
	if err != nil {
		return nil
	} else {
		return employee
	}
}
