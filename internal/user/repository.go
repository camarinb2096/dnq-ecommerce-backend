package user

import (
	"cmarin20/dnq-ecommerce/pkg/logger"

	"gorm.io/gorm"
)

type (
	Repository interface {
		CreateUser(user User) error
		FindUserByEmail(email string) User
	}

	repo struct {
		db     *gorm.DB
		logger *logger.Logger
	}
)

// TODO: manage errors and response
func NewRepository(db *gorm.DB, logger *logger.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) FindUserByEmail(email string) User {
	var user User
	r.db.Where("email = ?", email).First(&user)
	return user
}

func (r *repo) CreateUser(user User) error {
	r.db.Create(&user)
	return nil
}
