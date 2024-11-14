package products

import (
	dtos "cmarin20/dnq-ecommerce/internal/app/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Get     Controller
		GetByID Controller
	}
)

func NewEndpoints(s Services) Endpoints {
	return Endpoints{
		Get:     getProducts(s),
		GetByID: getByID(s),
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

func getByID(s Services) Controller {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		product, err := s.GetProductByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": product})
	}
}
