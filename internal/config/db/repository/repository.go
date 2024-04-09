package repository

import (
	userModel "cmarin20/dnq-ecommerce/internal/user/model"
	"cmarin20/dnq-ecommerce/pkg/logger"

	"gorm.io/gorm"
)

type (
	Repository interface {
		CreateUser(user userModel.User) error
		FindUserByEmail(email string) int
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

func (r *repo) FindUserByEmail(email string) int {
	r.logger.Info("Finding user by email...")
	var user userModel.User
	r.db.Where("email = ?", email).First(&user)
	return int(user.ID)
}

func (r *repo) CreateUser(user userModel.User) error {
	r.logger.Info("Creating a new user...")
	r.db.Create(&user)
	return nil
}
