package auth

import (
	"cmarin20/dnq-ecommerce/internal/user"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"fmt"
)

type Services interface {
	Login(data interface{}) (string, error)
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

func (s *services) Login(data interface{}) (string, error) {
	s.logger.Info("Logging in...")

	fmt.Println("!!!", data)
	return "", nil
}
