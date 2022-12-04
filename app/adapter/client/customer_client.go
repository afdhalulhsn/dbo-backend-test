package client

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type CustomerClient struct {
	con *gorm.DB
}

func NewCustomerClient(con *gorm.DB) *CustomerClient {
	return &CustomerClient{
		con: con,
	}
}

func (u *CustomerClient) GetDetailCustomer(ctx context.Context, id string) (*entity.CustomerData, error) {
	var detail *entity.CustomerData
	err := u.con.WithContext(ctx).Where(&entity.CustomerData{
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
		return nil, errors.New("CUSTOMER TIDAK DITEMUKAN")
	}
	return detail, nil
}

func (u *CustomerClient) GetAllCustomer(ctx context.Context) ([]*entity.CustomerData, error) {
	var detail []*entity.CustomerData
	err := u.con.WithContext(ctx).Where("status = 1").Find(&detail).Error
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("CUSTOMER TIDAK DITEMUKAN")
	}
	return detail, nil
}

func (u *CustomerClient) GetAllCustomerWithPaging(ctx context.Context, page, limit int) (*model.CustomerDataWithPaging, error) {
	var detail []*entity.CustomerData
	var count int64
	err := u.con.WithContext(ctx).Table("customer_data").Count(&count).Error
	if err != nil {
		return nil, err
	}

	err = u.con.WithContext(ctx).Where("status = 1").Offset(page - 1).Limit(limit).Find(&detail).Error
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("CUSTOMER TIDAK DITEMUKAN")
	}
	out := &model.CustomerDataWithPaging{
		ListCustomer:  detail,
		TotalCustomer: int(count),
	}
	return out, nil
}

func (u *CustomerClient) SearchCustomer(ctx context.Context, key string) ([]*entity.CustomerData, error) {
	var detail []*entity.CustomerData
	var err error
	found := true
	sb := strings.Builder{}
	if key != "" {
		sb.WriteString("customer_data.status = 1")
		sb.WriteString("customer_data.nama_customer like '%" + key + "%'")
		sb.WriteString(" or customer_data.email like '%" + key + "%'")
		sb.WriteString(" or customer_data.no_handphone like '%" + key + "%'")
		sb.WriteString(" or customer_data.no_id like '%" + key + "%'")
		sb.WriteString(" or customer_data.date_of_birth like '%" + key + "%'")
	}
	query := u.con.WithContext(ctx)
	if key != "" {
		err = query.Where(sb.String()).Find(&detail).Error
	} else {
		err = query.Find(&detail).Error
	}
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("CUSTOMER TIDAK DITEMUKAN")
	}
	return detail, nil
}

func (u *CustomerClient) DeleteCustomerData(ctx context.Context, id string) error {
	var detail *entity.CustomerData
	err := u.con.WithContext(ctx).Where(&entity.CustomerData{
		BaseModel: entity.BaseModel{
			ID: id,
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
		return errors.New("CUSTOMER TIDAK DITEMUKAN")
	}
	detail.Status = 0
	err = u.con.Updates(&detail).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *CustomerClient) UpdateCustomerData(ctx context.Context, data *model.CustomerUpdateModel) error {

	var detail *entity.CustomerData
	err := u.con.WithContext(ctx).Where(&entity.CustomerData{
		BaseModel: entity.BaseModel{
			ID: data.IdCustomer,
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
		return errors.New("CUSTOMER TIDAK DITEMUKAN")
	}
	detail.Email = data.Email
	detail.NoId = data.NoId
	detail.NoHp = data.NoHp
	detail.DateOfBirth = data.DateOfBirth
	detail.Gender = data.Gender
	err = u.con.Updates(&detail).Error
	if err != nil {
		return err
	}
	return nil
}
