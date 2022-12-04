package client

import (
	"context"
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/domain/repository"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserClient struct {
	conn *gorm.DB
}

func NewUserClient(con *gorm.DB) repository.UserRepository {
	return &UserClient{
		conn: con,
	}
}

func (u *UserClient) Daftar(ctx context.Context, customerData *entity.UserData) error {
	//Chek if user exists
	var customer *entity.CustomerData
	err := u.conn.WithContext(ctx).Where(&entity.CustomerData{Email: customerData.CustomerData.Email,
		NoId: customerData.CustomerData.NoId}).First(&customer).Error
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		found = false
	}
	if found {
		return errors.New("NO ID DAN EMAIL SUDAH TERDAFTAR")
	}

	var user *entity.UserData
	err = u.conn.WithContext(ctx).Where(&entity.UserData{UserName: customerData.UserName}).First(&user).Error
	found = true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		found = false
	}
	if found {
		return errors.New("USERNAME SUDAH TERDAFTAR")
	}
	err = u.conn.WithContext(ctx).Create(&customerData).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserClient) Login(ctx context.Context, login *model.LoginModel) (*entity.UserData, error) {
	var user *entity.UserData
	found := true
	err := u.conn.WithContext(ctx).Where(&entity.UserData{
		UserName: login.Username,
	}).Preload("CustomerData").First(&user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("USERNAME TIDAK TERDAFTAR SILAHKAN DAFTAR TERLEBIH DAHULU")
	}
	if user.Status == 0 {
		return nil, errors.New("User Tidak Aktif")
	}
	if user.UserName == login.Username && user.Password == login.Password {
		loginTime := time.Now()
		loginHistoru := &entity.LoginHistory{
			IdUser:      user.ID,
			LoginTime:   loginTime.Format("2006-01-02 15:04:05"),
			ExpiredTime: loginTime.Add(time.Hour).Format("2006-01-02 15:04:05"),
			Status:      "I",
		}
		err := u.conn.WithContext(ctx).Create(&loginHistoru).Error
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		return nil, errors.New("Username & Password Tidak Sesuai")
	}

}

func (u *UserClient) GantiPassword(ctx context.Context, login *model.UpdateUserModel) error {
	var user *entity.UserData
	found := true
	err := u.conn.WithContext(ctx).Where(&entity.UserData{
		BaseModel: entity.BaseModel{
			ID: login.IdUsername,
		},
	}).First(&user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		found = false
	}
	if !found {
		return errors.New("USER TIDAK DITEMUKAN")
	}
	if user.Password != login.PasswordLama {
		return errors.New("PASSWORD LAMA TIDAK COCOK")
	}
	user.Password = login.PasswordBaru
	err = u.conn.WithContext(ctx).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *UserClient) GetDetailUser(ctx context.Context, id string) (*entity.UserData, error) {
	var user *entity.UserData
	found := true
	err := u.conn.WithContext(ctx).Where(&entity.UserData{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}).Preload("CustomerData").First(&user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		found = false
	}
	if !found {
		return nil, errors.New("USERNAME TIDAK DITEMUKAN")
	}
	return user, nil
}

func (u *UserClient) DeleteUser(ctx context.Context, login string) error {
	var user *entity.UserData
	found := true
	err := u.conn.WithContext(ctx).Where(&entity.UserData{
		BaseModel: entity.BaseModel{
			ID: login,
		},
	}).First(&user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		found = false
	}
	if !found {
		return errors.New("USER TIDAK DITEMUKAN")
	}
	user.Status = 0
	err = u.conn.WithContext(ctx).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}
