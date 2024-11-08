//go:build wireinject
// +build wireinject

package di

import (
	"cmarin20/dnq-ecommerce/internal/config/server"

	"github.com/google/wire"
)

type App struct {
	Server *server.Server
}

func Initialize() (*App, error) {
	wire.Build(AppSet)
	return &App{}, nil
}
