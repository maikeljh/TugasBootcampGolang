package cart

import (
	"awesomeProject/internal/entity"
	"awesomeProject/internal/repository/cart"
	"context"
)

type CartUseCase interface {
	GetCart(ctx context.Context, page int, limit int) ([]entity.Cart, error)
	UpdateCart(ctx context.Context, id int64) ([]entity.Cart, error)
	Checkout(ctx context.Context) ([]entity.Cart, error)
}

type usecase struct {
	repo cart.Repo
}

func NewUseCase(repo cart.Repo) usecase {
	return usecase{repo: repo}
}

func (u usecase) GetCart(ctx context.Context, page int, limit int) ([]entity.Cart, error) {
	return u.repo.GetCart(ctx, page, limit)
}

func (u usecase) UpdateCart(ctx context.Context, id int64) ([]entity.Cart, error) {
	return u.repo.UpdateCart(ctx, id)
}

func (u usecase) Checkout(ctx context.Context) ([]entity.Cart, error) {
	return u.repo.Checkout(ctx)
}
