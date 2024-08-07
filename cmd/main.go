package main

import (
	"cmarin20/dnq-ecommerce/internal/di"
	"log"
)

func main() {

	app, err := di.Initialize()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	app.Server.Run()
}
