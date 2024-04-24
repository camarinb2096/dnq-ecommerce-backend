package productsEndpoint

import (
	dtos "cmarin20/dnq-ecommerce/internal/dto"
	productsService "cmarin20/dnq-ecommerce/internal/products/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Get Controller
	}
)

func NewEndpoints(s productsService.Services) Endpoints {
	return Endpoints{
		Get: getProducts(s),
	}
}

func getProducts(s productsService.Services) Controller {
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
