package auth

import (
	dtos "cmarin20/dnq-ecommerce/internal/app/dto"
	"errors"

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

var ErrUserNotFound = errors.New("user not found")

func login(s Services) Controller {
	return func(c *gin.Context) {
		var data dtos.UserLogin

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		userLoged, err := s.Login(data)
		if err != nil {

			switch err.Error() {
			case "user not found":
				c.JSON(404, gin.H{"error": err.Error()})
				return
			case "invalid password":
				c.JSON(401, gin.H{"error": err.Error()})
				return
			default:
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(200, gin.H{
			"message": "Login successful",
			"data":    userLoged})
	}
}
