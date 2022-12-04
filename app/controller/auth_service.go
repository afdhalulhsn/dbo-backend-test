package controller

import (
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/usecase/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthService struct {
	uc authentication.AuthenticationPort
}

func NewAuthenticationService(uc authentication.AuthenticationPort) *AuthService {
	return &AuthService{
		uc: uc,
	}
}

// Pendaftaran godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Pendaftaran
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData
// @Param Payload body model.DaftarModel true "DaftarModel"
// @Router /pendaftaran [post]
func (a *AuthService) Pendaftaran(c *gin.Context) {
	var data *model.DaftarModel
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "INVALID OBJECT")
		return
	}
	err = a.uc.DaftarUser(c.Copy(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, model.ResponseData{
		IsError:   false,
		ErrorCode: "00",
		Message:   "SUCCES",
		Data:      nil,
	})
	return
}

// Login godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Login
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{data=entity.UserData}
// @Param Payload body model.LoginModel true "LoginModel"
// @Router /login [post]
func (a *AuthService) Login(c *gin.Context) {
	var data *model.LoginModel
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "INVALID OBJECT")
		return
	}
	res, err := a.uc.Login(c.Copy(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, model.ResponseData{
		IsError:   false,
		ErrorCode: "00",
		Message:   "SUCCES",
		Data:      res,
	})
	return
}

// GetLoginData godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetLoginData
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{data=entity.LoginHistory}
// @Param Payload body model.LoginModel true "LoginModel"
// @Router /login-data [get]
func (a *AuthService) GetLoginData(c *gin.Context) {
	res, err := a.uc.GetLoginData(c.Copy())
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, model.ResponseData{
		IsError:   false,
		ErrorCode: "00",
		Message:   "SUCCES",
		Data:      res,
	})
	return
}
