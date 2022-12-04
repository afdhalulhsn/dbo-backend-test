package client

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type OrderClient struct {
	con *gorm.DB
}

func NewOrderClient(con *gorm.DB) *OrderClient {
	return &OrderClient{
		con: con,
	}
}

func (u *OrderClient) GetDetailOrder(ctx context.Context, id string) (*entity.DataOrder, error) {
	var detail *entity.DataOrder
	err := u.con.WithContext(ctx).Preload("ProductData").Preload("CustomerData").Where(&entity.DataOrder{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}).First(&detail).Error
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("ORDER TIDAK DITEMUKAN")
	}
	return detail, nil
}

func (u *OrderClient) GetAllOrder(ctx context.Context) ([]*entity.DataOrder, error) {
	var detail []*entity.DataOrder
	err := u.con.WithContext(ctx).Preload("ProductData").Preload("CustomerData").Find(&detail).Error
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("ORDER TIDAK DITEMUKAN")
	}
	return detail, nil
}

func (u *OrderClient) GetAllOrderWithPaging(ctx context.Context, page, limit int) (*model.OrderDataWithPaging, error) {
	var detail []*entity.DataOrder

	var count int64
	err := u.con.WithContext(ctx).Table("data_order").Count(&count).Error
	if err != nil {
		return nil, err
	}
	err = u.con.WithContext(ctx).Preload("ProductData").Preload("CustomerData").Offset(page - 1).Limit(limit).Find(&detail).Error
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("ORDER TIDAK DITEMUKAN")
	}
	out := &model.OrderDataWithPaging{
		List:  detail,
		Total: int(count),
	}
	return out, nil
}

func (u *OrderClient) SearchDataOrder(ctx context.Context, key string) ([]*entity.DataOrder, error) {
	var detail []*entity.DataOrder
	sb := strings.Builder{}
	if key != "" {
		sb.WriteString("customer_data.nama_customer like '%" + key + "%'")
		sb.WriteString(" or customer_data.email like '%" + key + "%'")
		sb.WriteString(" or customer_data.no_handphone like '%" + key + "%'")
		sb.WriteString(" or customer_data.no_id like '%" + key + "%'")
		sb.WriteString(" or customer_data.date_of_birth like '%" + key + "%'")
		sb.WriteString(" or customer_data.date_of_birth like '%" + key + "%'")

		sb.WriteString(" or product_data.kode_produk like '%" + key + "%'")
		sb.WriteString(" or product_data.nama_produck like '%" + key + "%'")
		sb.WriteString(" or product_data.harga like '%" + key + "%'")

	}
	err := u.con.WithContext(ctx).
		Joins("join customer_data on customer_data.id = data_order.id_customer").
		Joins("join product_data on product_data.id = data_order.id_produck").
		Preload("ProductData").
		Preload("CustomerData").
		Where(sb.String()).
		Find(&detail).Error
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("ORDER TIDAK DITEMUKAN")
	}
	return detail, nil
}

func (u *OrderClient) UpdateOrderData(ctx context.Context, data *model.UpdateOrder) error {
	var detail *entity.DataOrder
	err := u.con.WithContext(ctx).Preload("ProductData").Preload("CustomerData").Where(&entity.DataOrder{
		BaseModel: entity.BaseModel{
			ID: data.IdOrder,
		},
	}).First(&detail).Error
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		found = false
	}
	if !found {
		return errors.New("ORDER TIDAK DITEMUKAN")
	}
	detail.Qty = data.Qty
	return nil
}

func (u *OrderClient) DeleteDataOrder(ctx context.Context, id string) error {
	var detail *entity.DataOrder
	err := u.con.WithContext(ctx).Preload("ProductData").Preload("CustomerData").Where(&entity.DataOrder{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}).Delete(&detail).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *OrderClient) InsertDataOrder(ctx context.Context, data *model.InsertOrder) error {
	order := &entity.DataOrder{
		IdProduct:  data.IdProduct,
		IdCustomer: data.IdCustomer,
		OrderDate:  time.Now(),
		Qty:        data.Qty,
	}

	err := u.con.WithContext(ctx).Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}
