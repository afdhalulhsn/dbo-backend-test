package authentication

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/domain/repository"
	"dbo_backend_test/app/helper/jwt_token"
)

type AuthenticationUC struct {
	userRepo repository.UserRepository
	//customer repository.CustomerRepository
	aut repository.AuthManagement
}

func NewAuthenticationUC(user repository.UserRepository,
	aut repository.AuthManagement) AuthenticationPort {
	return &AuthenticationUC{
		userRepo: user,
		aut:      aut,
	}
}

func (uc *AuthenticationUC) DaftarUser(ctx context.Context, data *model.DaftarModel) error {
	userdata := &entity.UserData{
		UserName: data.UserName,
		Password: data.Password,
		Role:     data.Role,
		Status:   1,
		CustomerData: entity.CustomerData{
			Name:        data.Name,
			DateOfBirth: data.DateOfBirth,
			Gender:      data.Gender,
			NoHp:        data.NoHp,
			NoId:        data.NoId,
			Email:       data.Email,
			Status:      1,
		},
	}
	err := uc.userRepo.Daftar(ctx, userdata)
	return err
}

func (uc *AuthenticationUC) Login(ctx context.Context, data *model.LoginModel) (*entity.UserData, error) {
	resp, err := uc.userRepo.Login(ctx, data)
	if err != nil {
		return nil, err
	}
	resTok, errTok := jwt_token.GenerateToken(jwt_token.TokenRequest{
		Username: resp.UserName,
		Role:     resp.Role,
	})
	if errTok != nil {
		return nil, errTok
	}

	resp.Token = resTok.Token
	return resp, err
}

func (uc *AuthenticationUC) GetLoginData(ctx context.Context) ([]*entity.LoginHistory, error) {
	return uc.aut.GetLoginData(ctx)
}
