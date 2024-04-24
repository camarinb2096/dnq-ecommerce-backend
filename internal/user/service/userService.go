package userService

import (
	"cmarin20/dnq-ecommerce/internal/config/db/repository"
	dtos "cmarin20/dnq-ecommerce/internal/dto"
	userModel "cmarin20/dnq-ecommerce/internal/user/model"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"cmarin20/dnq-ecommerce/pkg/utils"
	"fmt"
	"strings"
)

type Services interface {
	CreateUser(user dtos.CreateRequest) error
}

type service struct {
	repo   repository.Repository
	logger *logger.Logger
}

func NewService(repo repository.Repository, logger *logger.Logger) Services {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) CreateUser(user dtos.CreateRequest) error {
	var userModel userModel.User
	s.logger.Info("Creating a new user...")

	if user.Name == "" {
		return fmt.Errorf("name is required")
	}

	if user.Email == "" {
		return fmt.Errorf("email is required")
	} else if !utils.IsValidEmail(user.Email) {
		return fmt.Errorf("invalid email")
	}

	if user.Password == "" {
		return fmt.Errorf("password is required")
	} else if !utils.IsStrongPassword(user.Password) {
		return fmt.Errorf("the password must contain at least 8 characters, 1 uppercase letter, 1 lowercase letter, 1 number and 1 special character")
	}

	existingUser := s.repo.FindUserByEmail(user.Email)
	if existingUser != 0 {
		return fmt.Errorf("user already exists")
	}

	hashedPassword, err := utils.EncodePassword(user.Password)
	if err != nil {
		s.logger.Error("Error encoding password: %v", err)
	}

	userModel.Name = strings.ToUpper(user.Name)
	userModel.Email = user.Email
	userModel.Prefix = user.Prefix
	userModel.Phone = user.Phone
	userModel.Address = user.Address
	userModel.Password = hashedPassword
	userModel.FkRole = 1

	err = s.repo.CreateUser(userModel)
	if err != nil {
		return err
	}

	return nil
}
