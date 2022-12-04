package product_uc

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/domain/repository"
)

type ProductUc struct {
	repo repository.ProductRepo
}

func NewProductUC(repo repository.ProductRepo) *ProductUc {
	return &ProductUc{
		repo: repo,
	}
}

func (u *ProductUc) AddProduct(ctx context.Context, in *model.AddProduct) error {
	return u.repo.AddProduct(ctx, in)
}

func (u *ProductUc) GetProduct(ctx context.Context) ([]*entity.ProductData, error) {
	return u.repo.GetProduct(ctx)
}
