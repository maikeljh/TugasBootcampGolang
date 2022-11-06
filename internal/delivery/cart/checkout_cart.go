package cart

import (
	"awesomeProject/internal/usecase/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Checkout(useCase cart.CartUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := useCase.Checkout(c.Copy().Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
