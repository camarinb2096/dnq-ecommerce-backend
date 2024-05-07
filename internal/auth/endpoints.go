package auth

import (
	dtos "cmarin20/dnq-ecommerce/internal/dto"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Login Controller
	}
)

func NewEndpoints(s Services) Endpoints {
	return Endpoints{
		Login: login(s),
	}
}

func login(s Services) Controller {
	return func(c *gin.Context) {
		var data dtos.UserLogin
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		token, err := s.Login(data)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Login successful",
			"token":   token})

	}
}