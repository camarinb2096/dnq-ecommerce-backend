package repository

import (
	dtos "cmarin20/dnq-ecommerce/internal/dto"
	productsModel "cmarin20/dnq-ecommerce/internal/products/model"
	userModel "cmarin20/dnq-ecommerce/internal/user/model"
	"cmarin20/dnq-ecommerce/pkg/logger"

	"gorm.io/gorm"
)

type (
	Repository interface {
		CreateUser(user userModel.User) error
		FindUserByEmail(email string) int
		CountProducts() int
		FindProducts(page, pageSize int) []dtos.Product
	}

	repo struct {
		db     *gorm.DB
		logger *logger.Logger
	}
)

// TODO: manage errors and response
func NewUserRepo(db *gorm.DB, logger *logger.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) FindUserByEmail(email string) int {
	var user userModel.User
	r.db.Where("email = ?", email).First(&user)
	return int(user.ID)
}

func (r *repo) CreateUser(user userModel.User) error {
	r.db.Create(&user)
	return nil
}

func (r *repo) CountProducts() int {
	var count int64
	r.db.Model(&productsModel.Product{}).Count(&count)
	return int(count)
}
func (r *repo) FindProducts(page, pageSize int) []dtos.Product {
	var products []dtos.Product
	offset := (page - 1) * pageSize
	r.db.Limit(pageSize).Offset(offset).Find(&products)
	return products
}
