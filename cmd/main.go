package main

import (
	"cmarin20/dnq-ecommerce/internal/config/db"
	"cmarin20/dnq-ecommerce/internal/config/db/repository"
	"cmarin20/dnq-ecommerce/internal/config/server"
	productsEndpoint "cmarin20/dnq-ecommerce/internal/products/endpoint"
	productsService "cmarin20/dnq-ecommerce/internal/products/service"
	userEndpoint "cmarin20/dnq-ecommerce/internal/user/endpoint"
	userService "cmarin20/dnq-ecommerce/internal/user/service"
	"cmarin20/dnq-ecommerce/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {

	logger := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	mysqlDb := db.NewDbConn(db.NewDbConfig(), logger)

	dbRepository := repository.NewUserRepo(mysqlDb, logger)
	userService := userService.NewService(dbRepository, logger)
	userEndpoints := userEndpoint.NewEndpoints(userService)
	productsService := productsService.NewService(dbRepository, logger)
	productsEndpoint := productsEndpoint.NewEndpoints(productsService)

	//GIN server instance
	server := server.NewServer()
	server.Routes(userEndpoints, productsEndpoint)
	server.Run(logger)
}
