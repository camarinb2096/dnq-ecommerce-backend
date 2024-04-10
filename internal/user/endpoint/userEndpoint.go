package userEndpoint

import (
	userDto "cmarin20/dnq-ecommerce/internal/user/dto"
	userService "cmarin20/dnq-ecommerce/internal/user/service"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Post Controller
	}
)

func NewEndpoints(s userService.Services) Endpoints {
	return Endpoints{
		Post: createUser(s),
	}
}

func createUser(s userService.Services) Controller {
	return func(c *gin.Context) {
		var user userDto.CreateRequest
		json.NewDecoder(c.Request.Body).Decode(&user)

		err := s.CreateUser(user)
		if err != nil {
			if err.Error() == "name is required" {
				c.JSON(400, gin.H{"error": "Name is required"})
				return
			}
			if err.Error() == "email is required" {
				c.JSON(400, gin.H{"error": "Email is required"})
				return
			}
			if err.Error() == "password is required" {
				c.JSON(400, gin.H{"error": "Password is required"})
				return
			}
			if err.Error() == "invalid email" {
				c.JSON(400, gin.H{"error": "Invalid email"})
				return
			}
			if err.Error() == "the password must contain at least 8 characters, 1 uppercase letter, 1 lowercase letter, 1 number and 1 special character" {
				c.JSON(400, gin.H{"error": "The password must contain at least 8 characters, 1 uppercase letter, 1 lowercase letter, 1 number and 1 special character"})
				return
			}
			if err.Error() == "user already exists" {
				c.JSON(409, gin.H{"error": "User already exists"})
				return
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(200, gin.H{"message": "User created successfully"})
	}
}
