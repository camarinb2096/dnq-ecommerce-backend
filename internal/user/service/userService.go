package userService

import (
	userRepo "cmarin20/dnq-ecommerce/internal/user/repository"
	"cmarin20/dnq-ecommerce/pkg/logger"
)

type Services interface {
	CreateUser()
}

type service struct {
	repo   userRepo.Repository
	logger *logger.Logger
}

func NewService(repo userRepo.Repository, logger *logger.Logger) Services {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) CreateUser() {
	s.logger.Info("Creating a new user...")
}
