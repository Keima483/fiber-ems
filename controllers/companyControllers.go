package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/fiber-ems/models"
	"github.com/keima483/fiber-ems/service"
)

func Login(c *fiber.Ctx) error {
	loginDetails := new(models.LoginModel)
	if err := c.BodyParser(loginDetails); err != nil {
		return err
	}
	company, err := service.Login(loginDetails)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(company)
}

func Signup(c *fiber.Ctx) error {
	companyDetails := new(models.CompanyModel)
	if err := c.BodyParser(companyDetails); err != nil {
		return err
	}
	company, err := service.SignUp(companyDetails)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	} 
	return c.JSON(company)
}

