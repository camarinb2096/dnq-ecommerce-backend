package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(userId uint, fkRole int) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"fkRole": fkRole,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
