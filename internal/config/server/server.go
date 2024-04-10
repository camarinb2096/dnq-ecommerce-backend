package server

import (
	userEndpoint "cmarin20/dnq-ecommerce/internal/user/endpoint"
	"cmarin20/dnq-ecommerce/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() Server {
	router := gin.Default()
	s := Server{router: router}
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

func (s *Server) Routes(userEndpoint userEndpoint.Endpoints) {
	s.router.Use(configCors())
	user := s.router.Group("/api/v1/user")
	{
		user.POST("/", func(c *gin.Context) {
			userEndpoint.Post(c)
		})
	}
}
func (s *Server) Run(logger *logger.Logger) {
	logger.Info("Starting the application...")
	s.router.Run()
}
