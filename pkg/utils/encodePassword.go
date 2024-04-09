package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
