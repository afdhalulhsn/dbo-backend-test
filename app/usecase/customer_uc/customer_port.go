package customer_uc

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
)

type CustomerPort interface {
	GetDetailCustomer(ctx context.Context, id string) (*entity.CustomerData, error)
	GetAllCustomer(ctx context.Context) ([]*entity.CustomerData, error)
	GetAllCustomerWithPaging(ctx context.Context, page, limit int) (*model.CustomerDataWithPaging, error)
	SearchCustomer(ctx context.Context, key string) ([]*entity.CustomerData, error)
	UpdateCustomerData(ctx context.Context, data *model.CustomerUpdateModel) error
	DeleteCustomerData(ctx context.Context, id string) error
}
