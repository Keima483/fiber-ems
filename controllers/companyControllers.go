package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/fiber-ems/models"
	"github.com/keima483/fiber-ems/service"
	"github.com/keima483/fiber-ems/util"
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
	token, exp, err := util.CreateJWTToken(company)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"token": token, "exp": exp, "company": company})
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
	token, exp, err := util.CreateJWTToken(company)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token, "exp": exp, "company": company})
}

