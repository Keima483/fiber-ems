package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/keima483/fiber-ems/initializers"
	"github.com/keima483/fiber-ems/repository"
)

func CreateJWTToken(company repository.Company) (string, int64, error) {
	config, _ := initializers.LoadConfig(".")
	exp := time.Now().Add(config.JwtExpiresIn).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = company.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}
