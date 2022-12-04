package client

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/repository"
	"gorm.io/gorm"
)

type AuthClient struct {
	conn *gorm.DB
}

func NewAuthClient(con *gorm.DB) repository.AuthManagement {
	return &AuthClient{
		conn: con,
	}
}

func (u *AuthClient) GetLoginData(ctx context.Context) ([]*entity.LoginHistory, error) {
	var out []*entity.LoginHistory
	err := u.conn.Preload("UserData").Preload("UserData.CustomerData").Find(&out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
