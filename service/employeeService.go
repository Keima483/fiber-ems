package service

import (
	"errors"
	"github.com/keima483/fiber-ems/models"
	"github.com/keima483/fiber-ems/repository"
)

func AddEmployee(companyId int, em *models.EmployeeModel) (repository.Employee, error) {
	if em.Name == "" && em.Department == "" && em.Email == "" {
		return repository.Employee{}, errors.New("enter name, department and email atleast")
	}
	company := new(repository.Company)
	repository.DB.First(&company, companyId)
	if company.Email == "" {
		return repository.Employee{}, errors.New("no company with such id")
	}
	emp := repository.EmployeeFromModel(*em)
	company.AddEmployee(&emp)
	repository.DB.Save(&company)
	return emp, nil
}

func GetEmployees(companyId int) ([]repository.Employee, error) {
	company := new(repository.Company)
	repository.DB.First(&company, companyId)
	if company.Email == "" {
		return []repository.Employee{}, errors.New("no company with such id")
	}
	var employees []repository.Employee
	repository.DB.Where("company_id = ?", companyId).Find(&employees)
	return employees, nil
}

func UpdateEmployee(emp repository.Employee) (repository.Employee, error) {
	var employee repository.Employee
	repository.DB.Find(&employee, emp.ID)
	if employee.Email == "" {
		return repository.Employee{}, errors.New("no such employee exist")
	}
	repository.DB.Save(&emp)
	return emp, nil
}

func DeleteEmployee(empId int) (bool, error) {
	var employee repository.Employee
	repository.DB.Find(&employee, empId)
	if employee.Email == "" {
		return false, errors.New("no employee with this id")
	}
	repository.DB.Delete(&employee)
	return true, nil
}