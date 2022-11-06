package delivery

import (
	"awesomeProject/internal/delivery/cart"
	"awesomeProject/internal/delivery/product"
	cart3 "awesomeProject/internal/repository/cart"
	product3 "awesomeProject/internal/repository/product"
	cart2 "awesomeProject/internal/usecase/cart"
	product2 "awesomeProject/internal/usecase/product"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"os"
)

func Routes(r *gin.Engine) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:anak1234@localhost:5432/ecommerce")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	repo1 := product3.NewRepo(conn)
	usecase1 := product2.NewUseCase(repo1)

	repo2 := cart3.NewRepo(conn)
	usecase2 := cart2.NewUseCase(repo2)

	r.GET("/products", product.GetProducts(usecase1))
	r.GET("/products/:id", product.GetProduct(usecase1))
	r.GET("/cart", cart.GetCart(usecase2))
	r.POST("/add-cart", cart.GetCart(usecase2))
	r.DELETE("/checkout", cart.GetCart(usecase2))
}
