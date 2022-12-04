package client

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
	"gorm.io/gorm"
)

type ProductClient struct {
	con *gorm.DB
}

func NewProductClient(con *gorm.DB) *ProductClient {
	return &ProductClient{
		con: con,
	}
}

func (u *ProductClient) AddProduct(ctx context.Context, in *model.AddProduct) error {
	data := &entity.ProductData{
		KodeProduk:  in.KodeProduk,
		NamaProduck: in.NamaProduck,
		Harga:       in.Harga,
	}
	return u.con.WithContext(ctx).Create(&data).Error
}

func (u *ProductClient) GetProduct(ctx context.Context) ([]*entity.ProductData, error) {
	var out []*entity.ProductData
	err := u.con.WithContext(ctx).Find(&out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
