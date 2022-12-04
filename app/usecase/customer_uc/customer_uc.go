package customer_uc

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/domain/repository"
)

type CustomerUC struct {
	customer repository.CustomerRepository
}

func NewCustomerUC(customer repository.CustomerRepository) *CustomerUC {
	return &CustomerUC{
		customer: customer,
	}
}

func (u *CustomerUC) GetDetailCustomer(ctx context.Context, id string) (*entity.CustomerData, error) {
	return u.customer.GetDetailCustomer(ctx, id)
}
func (u *CustomerUC) GetAllCustomer(ctx context.Context) ([]*entity.CustomerData, error) {
	return u.customer.GetAllCustomer(ctx)
}
func (u *CustomerUC) GetAllCustomerWithPaging(ctx context.Context, page, limit int) (*model.CustomerDataWithPaging, error) {
	return u.customer.GetAllCustomerWithPaging(ctx, page, limit)

}
func (u *CustomerUC) SearchCustomer(ctx context.Context, key string) ([]*entity.CustomerData, error) {
	return u.customer.SearchCustomer(ctx, key)
}

func (u *CustomerUC) UpdateCustomerData(ctx context.Context, data *model.CustomerUpdateModel) error {
	return u.customer.UpdateCustomerData(ctx, data)
}

func (u *CustomerUC) DeleteCustomerData(ctx context.Context, id string) error {
	return u.customer.DeleteCustomerData(ctx, id)
}
