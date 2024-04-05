package main

import (
	"cmarin20/dnq-ecommerce/internal/config/server"
	"cmarin20/dnq-ecommerce/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {

	logger := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
	//GIN server instance
	server := server.NewServer()
	server.Routes()
	server.Run(logger)
}
