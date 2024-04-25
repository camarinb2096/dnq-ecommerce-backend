package products

import (
	dtos "cmarin20/dnq-ecommerce/internal/dto"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"strconv"
)

type Services interface {
	GetProducts(requestParams dtos.RequestParams) dtos.Response
}

type service struct {
	repo   Repository
	logger *logger.Logger
}

func NewService(repo Repository, logger *logger.Logger) Services {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) GetProducts(requestParams dtos.RequestParams) dtos.Response {
	s.logger.Info("Getting products...")

	total := s.repo.CountProducts()
	if total == 0 {
		return dtos.Response{
			Message: "No products found",
			Total:   0,
			Page:    1,
		}
	}

	page, _ := strconv.Atoi(requestParams.Page)
	pageSize, _ := strconv.Atoi(requestParams.PageSize)
	name := requestParams.ProductName

	products := s.repo.FindProducts(name, page, pageSize)

	return dtos.Response{
		Message:  "Products retrieved successfully",
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Data:     products,
	}
}
