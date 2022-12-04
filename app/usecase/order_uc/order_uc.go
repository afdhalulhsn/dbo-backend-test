package order_uc

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/domain/repository"
)

type OrderUc struct {
	order repository.OrderRepository
}

func NewOrderUc(order repository.OrderRepository) *OrderUc {
	return &OrderUc{
		order: order,
	}
}

func (u *OrderUc) GetDetailOrder(ctx context.Context, id string) (*entity.DataOrder, error) {
	return u.order.GetDetailOrder(ctx, id)
}
func (u *OrderUc) GetAllOrder(ctx context.Context) ([]*entity.DataOrder, error) {
	return u.order.GetAllOrder(ctx)
}
func (u *OrderUc) GetAllOrderWithPaging(ctx context.Context, page, limit int) (*model.OrderDataWithPaging, error) {
	return u.order.GetAllOrderWithPaging(ctx, page, limit)

}
func (u *OrderUc) SearchDataOrder(ctx context.Context, key string) ([]*entity.DataOrder, error) {
	return u.order.SearchDataOrder(ctx, key)
}

func (u *OrderUc) UpdateOrderData(ctx context.Context, key *model.UpdateOrder) error {
	return u.order.UpdateOrderData(ctx, key)
}

func (u *OrderUc) DeleteDataOrder(ctx context.Context, key string) error {
	return u.order.DeleteDataOrder(ctx, key)
}

func (u *OrderUc) InsertDataOrder(ctx context.Context, data *model.InsertOrder) error {
	return u.order.InsertDataOrder(ctx, data)
}
