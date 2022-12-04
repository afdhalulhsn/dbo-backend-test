package controller

import (
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/usecase/customer_uc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerController struct {
	uc customer_uc.CustomerPort
}

type CustomerDetailBindigObject struct {
	ID string `uri:"id" binding:"required"`
}

func NewCustomerService(uc customer_uc.CustomerPort) *CustomerController {
	return &CustomerController{
		uc: uc,
	}
}

// GetDetailCustomer godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetDetailCustomer
// @Accept */*
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.ResponseData{data=entity.CustomerData}
// @Router /admin/customer-admin/detail/:id [get]
func (s *CustomerController) GetDetailCustomer(c *gin.Context) {
	var person CustomerDetailBindigObject
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, "ID TIDAK BOLEH KOSONG")
		return
	}
	res, err := s.uc.GetDetailCustomer(c.Copy(), person.ID)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, model.ResponseData{
		IsError:   false,
		ErrorCode: "00",
		Message:   "SUCCES",
		Data:      res,
	})
	return
}

// GetAllCustomer godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetAllCustomer
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{data=[]entity.CustomerData}
// @Router /admin/customer-admin [get]
func (s *CustomerController) GetAllCustomer(c *gin.Context) {
	res, err := s.uc.GetAllCustomer(c.Copy())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, model.ResponseData{
		IsError:   false,
		ErrorCode: "00",
		Message:   "SUCCES",
		Data:      res,
	})
	return
}

// GetAllCustomerPagination godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetAllCustomerPagination
// @Accept */*
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} model.ResponseData{data=[]entity.CustomerData}
// @Router /admin/customer-admin/paging [get]
func (s *CustomerController) GetAllCustomerPagination(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	res, err := s.uc.GetAllCustomerWithPaging(c.Copy(), intPage, intLimit)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, model.ResponseData{
		IsError:   false,
		ErrorCode: "00",
		Message:   "SUCCES",
		Data:      res,
	})
	return
}

// SearchAllCustomer godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags SearchAllCustomer
// @Accept */*
// @Produce json
// @Param key query string true "key"
// @Success 200 {object} model.ResponseData{data=[]entity.CustomerData}
// @Router /admin/customer-admin/search [get]
func (s *CustomerController) SearchAllCustomer(c *gin.Context) {
	key := c.Query("key")
	res, err := s.uc.SearchCustomer(c.Copy(), key)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, model.ResponseData{
		IsError:   false,
		ErrorCode: "00",
		Message:   "SUCCES",
		Data:      res,
	})
	return
}

// DeleteCustomer godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags DeleteCustomer
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{}
// @Router /admin/customer-admin/detail/:id [delete]
func (s *CustomerController) DeleteCustomer(c *gin.Context) {
	var person OrderDetailBindigObject
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, "ID TIDAK BOLEH KOSONG")
		return
	}
	err := s.uc.DeleteCustomerData(c.Copy(), person.ID)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, model.ResponseData{
		IsError:   false,
		ErrorCode: "00",
		Message:   "SUCCES",
	})
	return
}

// UpdateCustomer godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags UpdateCustomer
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{}
// @Param Payload body model.CustomerUpdateModel true "CustomerUpdateModel"
// @Router /admin/customer-admin/update [put]
func (s *CustomerController) UpdateCustomer(c *gin.Context) {
	var data *model.CustomerUpdateModel
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "INVALID OBJECT")
		return
	}
	err = s.uc.UpdateCustomerData(c.Copy(), data)
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
