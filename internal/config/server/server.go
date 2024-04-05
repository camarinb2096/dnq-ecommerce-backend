package server

import (
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

func (s *Server) Routes() {
	s.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (s *Server) Run(logger *logger.Logger) {
	logger.Info("Starting the application...")
	s.router.Run()
}
