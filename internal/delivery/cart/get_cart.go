package cart

import (
	"awesomeProject/internal/usecase/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCart(useCase cart.CartUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		limit := 10

		result, err := useCase.GetCart(c.Copy().Request.Context(), page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
