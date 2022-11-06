package cart

import (
	"awesomeProject/internal/usecase/cart"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateCart(useCase cart.CartUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		fmt.Printf("%d", id)
		result, err := useCase.UpdateCart(c.Copy().Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
