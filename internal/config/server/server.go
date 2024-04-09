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

func (s *Server) Routes(userEndpoint userEndpoint.Endpoints) {
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
