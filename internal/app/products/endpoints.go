package products

import (
	dtos "cmarin20/dnq-ecommerce/internal/app/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Get Controller
	}
)

func NewEndpoints(s Services) Endpoints {
	return Endpoints{
		Get: getProducts(s),
	}
}

func getProducts(s Services) Controller {
	return func(c *gin.Context) {
		var requestParams dtos.RequestParams

		requestParams.Page = c.DefaultQuery("page", "1")
		requestParams.PageSize = c.DefaultQuery("pageSize", "10")
		requestParams.ProductName = c.DefaultQuery("name", "")
		response := s.GetProducts(requestParams)
		if response.Total == 0 {
			c.JSON(http.StatusNotFound, response)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
