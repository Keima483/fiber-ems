package service

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/keima483/fiber-ems/models"
	"github.com/keima483/fiber-ems/initializers"
	"github.com/keima483/fiber-ems/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(lm *models.LoginModel) (repository.Company, error) {
	company := new(repository.Company)
	initializers.DB.Where("email = ?", lm.Email).First(&company)
	if company.Email == "" {
		return repository.Company{}, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(company.Password), []byte(lm.Password)); err != nil {
		return repository.Company{}, fiber.NewError(fiber.StatusBadRequest, "Incorrect Password")
	}
	return *company, nil
}

func SignUp(cm *models.CompanyModel) (repository.Company, error) {
	if cm.Email == "" && cm.Password == "" {
		return repository.Company{}, errors.New("enter email and password atleast")
	}
	company := repository.CompanyFromModel(*cm)
	hash, err := bcrypt.GenerateFromPassword([]byte(company.Password), bcrypt.DefaultCost)
	if err != nil {
		return repository.Company{}, errors.New("not able to encrypt password")
	}
	company.Password = string(hash)
	initializers.DB.Save(&company)
	return company, nil
}

