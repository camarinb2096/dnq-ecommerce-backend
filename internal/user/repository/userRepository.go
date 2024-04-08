package userRepo

import (
	"cmarin20/dnq-ecommerce/pkg/logger"

	"gorm.io/gorm"
)

type (
	Repository interface {
		CreateUser()
	}

	repo struct {
		db     *gorm.DB
		logger *logger.Logger
	}
)

func NewUserRepo(db *gorm.DB, logger *logger.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) CreateUser() {
	r.logger.Info("Creating a new user...")
}
