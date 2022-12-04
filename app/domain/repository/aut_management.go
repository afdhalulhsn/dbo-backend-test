package repository

import (
	"context"
	"dbo_backend_test/app/domain/entity"
)

type AuthManagement interface {
	GetLoginData(ctx context.Context) ([]*entity.LoginHistory, error)
}
