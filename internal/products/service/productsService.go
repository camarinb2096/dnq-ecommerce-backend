package productsService

import (
	"cmarin20/dnq-ecommerce/internal/config/db/repository"
	productsDto "cmarin20/dnq-ecommerce/internal/products/dto"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"strconv"
)

type Services interface {
	GetProducts(requestParams productsDto.RequestParams) productsDto.Response
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

func (s *service) GetProducts(requestParams productsDto.RequestParams) productsDto.Response {
	s.logger.Info("Getting products...")

	total := s.repo.CountProducts()
	if total == 0 {
		return productsDto.Response{
			Message: "No products found",
			Total:   0,
			Page:    1,
		}
	}

	page, _ := strconv.Atoi(requestParams.Page)
	pageSize, _ := strconv.Atoi(requestParams.PageSize)
	products := s.repo.FindProducts(page, pageSize)

	return productsDto.Response{
		Message:  "Products retrieved successfully",
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Data:     products,
	}
}
