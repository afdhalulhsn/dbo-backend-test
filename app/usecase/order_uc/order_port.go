package order_uc

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
)

type OrderPort interface {
	GetDetailOrder(ctx context.Context, id string) (*entity.DataOrder, error)
	GetAllOrder(ctx context.Context) ([]*entity.DataOrder, error)
	GetAllOrderWithPaging(ctx context.Context, page, limit int) (*model.OrderDataWithPaging, error)
	SearchDataOrder(ctx context.Context, key string) ([]*entity.DataOrder, error)
	UpdateOrderData(ctx context.Context, data *model.UpdateOrder) error
	DeleteDataOrder(ctx context.Context, id string) error
	InsertDataOrder(ctx context.Context, data *model.InsertOrder) error
}
