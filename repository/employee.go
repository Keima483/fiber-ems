package repository

import (
	"github.com/keima483/fiber-ems/models"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	CompanyId  int
	Name       string
	Age        int
	Department string
	Email      string
	BaseSalary float32
}

func EmployeeFromModel(model models.EmployeeModel) Employee {
	return Employee{
		Name:       model.Name,
		Age:        model.Age,
		Department: model.Department,
		Email:      model.Email,
		BaseSalary: model.BaseSalary,
	}
}
