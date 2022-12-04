package authentication

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
)

type AuthenticationPort interface {
	Login(ctx context.Context, data *model.LoginModel) (*entity.UserData, error)
	DaftarUser(ctx context.Context, data *model.DaftarModel) error
	GetLoginData(ctx context.Context) ([]*entity.LoginHistory, error)
}
