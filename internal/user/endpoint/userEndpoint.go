package userEndpoint

import "github.com/gin-gonic/gin"

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		CreateUser Controller
	}
)

func NewEndpoints() Endpoints {
	return Endpoints{
		CreateUser: createUser,
	}
}

func createUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create User",
	})
}
