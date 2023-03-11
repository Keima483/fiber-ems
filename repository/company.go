package repository

import (
	"github.com/keima483/fiber-ems/models"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name      string
	Email     string
	Address   string
	Employees []*Employee `gorm:"foreignKey:CompanyId"`
	Password  string      `json:"-"`
}

func CompanyFromModel(model models.CompanyModel) Company {
	return Company{
		Name:     model.Name,
		Email:    model.Email,
		Address:  model.Address,
		Password: model.Password,
	}
}

func (c *Company) AddEmployee(emp *Employee) {
	employees := []*Employee {emp}
	if c.Employees != nil {
		employees = append(c.Employees, employees...)
	}
	c.Employees = employees
}

// func (c Company) GetEmployees() []Employee {
// 	employees := []Employee {}
// 	for _, v := range c.Employees {
// 		employees = append(employees, *v)
// 	} 
// 	return employees
// }