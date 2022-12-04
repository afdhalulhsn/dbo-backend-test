package model

import "dbo_backend_test/app/domain/entity"

type OrderDataWithPaging struct {
	List  []*entity.DataOrder
	Total int
}
