package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/keima483/fiber-ems/repository"
)

func CreateJWTToken(company repository.Company) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = company.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}
