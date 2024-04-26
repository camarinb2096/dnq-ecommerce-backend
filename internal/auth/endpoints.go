package auth

import "github.com/gin-gonic/gin"

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
		//TODO: Implement login endpoint
		c.JSON(200, gin.H{
			"message": "Login endpoint",
		})
	}
}
