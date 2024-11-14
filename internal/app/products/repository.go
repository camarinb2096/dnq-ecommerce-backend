package products

import (
	dtos "cmarin20/dnq-ecommerce/internal/app/dto"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type (
	Repository interface {
		CountProducts() int
		FindProducts(name string, page, pageSize int) []dtos.Product
		FindProductByID(id int) (dtos.Product, error)
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

func (r *repo) CountProducts() int {
	var count int64
	r.db.Model(&Product{}).Count(&count)
	return int(count)
}
func (r *repo) FindProducts(name string, page, pageSize int) []dtos.Product {
	var products []dtos.Product
	offset := (page - 1) * pageSize

	likePatter := fmt.Sprintf("%%%s%%", strings.ToLower(name))

	r.db.Where("LOWER (name) LIKE ?", likePatter).Limit(pageSize).Offset(offset).Find(&products)
	return products
}

func (r *repo) FindProductByID(id int) (dtos.Product, error) {
	var product dtos.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
