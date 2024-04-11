package productsEndpoint

import (
	productsDto "cmarin20/dnq-ecommerce/internal/products/dto"
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
		var requestParams productsDto.RequestParams

		requestParams.Page = c.DefaultQuery("page", "1")
		requestParams.PageSize = c.DefaultQuery("pageSize", "10")
		response := s.GetProducts(requestParams)
		if response.Total == 0 {
			c.JSON(http.StatusNotFound, response)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
