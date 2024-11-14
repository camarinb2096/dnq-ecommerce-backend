package server

import (
	"cmarin20/dnq-ecommerce/internal/app/auth"
	"cmarin20/dnq-ecommerce/internal/app/products"
	"cmarin20/dnq-ecommerce/internal/app/user"

	// userEndpoint "cmarin20/dnq-ecommerce/internal/app/user"
	"cmarin20/dnq-ecommerce/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	logger *logger.Logger
}

func NewServer(
	userEndpoint user.Endpoints,
	productsEndpoint products.Endpoints,
	authEndpoint auth.Endpoints,
	logger *logger.Logger) *Server {
	router := gin.Default()
	s := &Server{
		router: router,
		logger: logger,
	}

	// Configurar rutas
	s.Routes(userEndpoint, productsEndpoint, authEndpoint)
	return s
}

func configCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (s *Server) Routes(userEndpoint user.Endpoints, productsEndpoint products.Endpoints, authEndpoint auth.Endpoints) {
	s.router.Use(configCors())
	user := s.router.Group("/api/v1/user")
	{
		user.POST("/", func(c *gin.Context) {
			userEndpoint.Post(c)
		})
	}

	auth := s.router.Group("/api/v1/login")
	{
		auth.POST("/", func(c *gin.Context) {
			authEndpoint.Login(c)
		})
	}

	product := s.router.Group("/api/v1/products")
	{
		product.GET("/", func(c *gin.Context) {
			productsEndpoint.Get(c)
		})
		product.GET("/:id", func(c *gin.Context) {
			productsEndpoint.GetByID(c)
		})
	}
}
func (s *Server) Run() {
	s.logger.Info("Starting the application...")
	s.router.Run()
}
