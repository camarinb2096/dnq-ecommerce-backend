package auth

import (
	"cmarin20/dnq-ecommerce/internal/app/auth/jwt"
	dtos "cmarin20/dnq-ecommerce/internal/app/dto"
	"cmarin20/dnq-ecommerce/internal/app/user"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"fmt"

	"cmarin20/dnq-ecommerce/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	Login(data dtos.UserLogin) (dtos.UserLoged, error)
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

func (s *services) Login(data dtos.UserLogin) (dtos.UserLoged, error) {

	if !utils.IsValidEmail(data.Email) {
		return dtos.UserLoged{}, fmt.Errorf("invalid email")
	}

	user := s.repo.FindUserByEmail(data.Email)
	if user.Email != data.Email {
		return dtos.UserLoged{}, fmt.Errorf("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return dtos.UserLoged{}, fmt.Errorf("invalid password")
	}

	token, err := jwt.GenerateToken(user.ID, user.FkRole)
	if err != nil {
		return dtos.UserLoged{}, fmt.Errorf("error generating token")
	}

	return dtos.UserLoged{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}, nil
}
