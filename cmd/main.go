package main

import (
	"cmarin20/dnq-ecommerce/internal/di"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	
	app, err := di.Initialize()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	app.Server.Run()
}
