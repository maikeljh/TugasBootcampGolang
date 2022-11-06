package product

import (
	"awesomeProject/internal/usecase/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProducts(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		limit := 10

		result, err := useCase.GetProducts(c.Copy().Request.Context(), page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetProduct(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		result, err := useCase.GetProduct(c.Copy().Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
