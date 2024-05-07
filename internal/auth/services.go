package auth

import (
	dtos "cmarin20/dnq-ecommerce/internal/dto"
	"cmarin20/dnq-ecommerce/internal/user"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"fmt"

	"cmarin20/dnq-ecommerce/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	Login(data dtos.UserLogin) (string, error)
}

type services struct {
	repo   user.Repository
	logger *logger.Logger
}

func NewService(repo user.Repository, logger *logger.Logger) Services {
	return &services{
		repo:   repo,
		logger: logger,
	}
}

func (s *services) Login(data dtos.UserLogin) (string, error) {

	if !utils.IsValidEmail(data.Email) {
		return "", fmt.Errorf("invalid email")
	}

	user := s.repo.FindUserByEmail(data.Email)
	if user.Email != data.Email {
		return "", fmt.Errorf("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	return "", nil
}
