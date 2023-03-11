package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/keima483/fiber-ems/models"
	"github.com/keima483/fiber-ems/repository"
	"github.com/keima483/fiber-ems/service"
)

func AddEmployee(c *fiber.Ctx) error {
	employeeDetail := new(models.EmployeeModel)
	if err := c.BodyParser(employeeDetail); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	employee, err := service.AddEmployee(id, employeeDetail)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	} 
	return c.JSON(employee)
}

func GetEmployees(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	employees, err := service.GetEmployees(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	} 
	return c.JSON(employees)
}

func UpdateEmployee(c *fiber.Ctx) error {
	var emp repository.Employee
	if err := c.BodyParser(&emp); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	emp, err := service.UpdateEmployee(emp)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(emp)
}

func DeleteEmployee(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := service.DeleteEmployee(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	} 
	return c.SendString("Successfuly deleted")
}