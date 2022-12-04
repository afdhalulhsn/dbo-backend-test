package repository

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
)

type ProductRepo interface {
	AddProduct(ctx context.Context, in *model.AddProduct) error
	GetProduct(ctx context.Context) ([]*entity.ProductData, error)
}
