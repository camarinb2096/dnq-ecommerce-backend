package main

import (
	"cmarin20/dnq-ecommerce/internal/auth"
	"cmarin20/dnq-ecommerce/internal/config/db"
	"cmarin20/dnq-ecommerce/internal/config/server"
	"cmarin20/dnq-ecommerce/internal/products"
	user "cmarin20/dnq-ecommerce/internal/user"
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

	userRepo := user.NewRepository(mysqlDb, logger)
	userService := user.NewService(userRepo, logger)
	userEndpoints := user.NewEndpoints(userService)

	productsRepo := products.NewRepository(mysqlDb, logger)
	productsService := products.NewService(productsRepo, logger)
	productsEndpoint := products.NewEndpoints(productsService)

	authService := auth.NewService(userRepo, logger)
	authEndpoints := auth.NewEndpoints(authService)

	//GIN server instance
	server := server.NewServer()
	server.Routes(userEndpoints, productsEndpoint, authEndpoints)
	server.Run(logger)
}
