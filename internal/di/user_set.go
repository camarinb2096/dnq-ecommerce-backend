//go:build wireinject
// +build wireinject

package di

import (
	"cmarin20/dnq-ecommerce/internal/app/auth"
	"cmarin20/dnq-ecommerce/internal/app/products"
	"cmarin20/dnq-ecommerce/internal/app/user"
	"cmarin20/dnq-ecommerce/internal/config/db"
	"cmarin20/dnq-ecommerce/internal/config/server"
	"cmarin20/dnq-ecommerce/pkg/logger"

	"github.com/google/wire"
)

var CommonSet = wire.NewSet(
	logger.NewLogger,
	db.NewDbConfig,
	db.NewDbConn,
)

var UserSet = wire.NewSet(
	user.NewRepository,
	user.NewService,
	user.NewEndpoints,
)

var ProductSet = wire.NewSet(
	products.NewRepository,
	products.NewService,
	products.NewEndpoints,
)

var AuthSet = wire.NewSet(
	auth.NewService,
	auth.NewEndpoints,
)

var ServerSet = wire.NewSet(
	server.NewServer,
)

var AppSet = wire.NewSet(
	CommonSet,
	UserSet,
	ProductSet,
	AuthSet,
	ServerSet,
	wire.Struct(new(App), "Server"),
)
