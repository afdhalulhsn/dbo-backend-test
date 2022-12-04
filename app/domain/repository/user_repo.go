package repository

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
)

type UserRepository interface {
	Login(ctx context.Context, model *model.LoginModel) (*entity.UserData, error)
	Daftar(ctx context.Context, customerData *entity.UserData) error
	GantiPassword(ctx context.Context, data *model.UpdateUserModel) error
	GetDetailUser(ctx context.Context, id string) (*entity.UserData, error)
	DeleteUser(ctx context.Context, id string) error
}
