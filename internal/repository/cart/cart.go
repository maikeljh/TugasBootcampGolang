package cart

import (
	"awesomeProject/internal/entity"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type Repo interface {
	GetCart(ctx context.Context, page int, limit int) ([]entity.Cart, error)
	UpdateCart(ctx context.Context, id int64) (result []entity.Cart, err error)
	Checkout(ctx context.Context) (result []entity.Cart, err error)
}

type repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) repo {
	return repo{db: db}
}

func (r repo) GetCart(ctx context.Context, page int, limit int) (result []entity.Cart, err error) {
	const query = "SELECT * FROM cart"
	err = pgxscan.Select(ctx, r.db, &result, query)
	return result, err
}

func (r repo) UpdateCart(ctx context.Context, id int64) (result []entity.Cart, err error) {
	const query = "SELECT * FROM products WHERE id=$1"
	var temp entity.Product
	err = pgxscan.Select(ctx, r.db, &temp, query, id)

	if err == nil {
		// add product to cart
	}
	return result, err // return new cart
}

func (r repo) Checkout(ctx context.Context) (result []entity.Cart, err error) {
	const query = "SELECT * FROM cart"
	err = pgxscan.Select(ctx, r.db, &result, query)
	return result, err
}
