package model

import "dbo_backend_test/app/domain/entity"

type CustomerDataWithPaging struct {
	ListCustomer  []*entity.CustomerData
	TotalCustomer int
}
