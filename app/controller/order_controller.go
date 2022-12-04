package controller

import (
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/usecase/order_uc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderController struct {
	uc order_uc.OrderPort
}

type OrderDetailBindigObject struct {
	ID string `uri:"id" binding:"required"`
}

func NewOrderController(uc order_uc.OrderPort) *OrderController {
	return &OrderController{
		uc: uc,
	}
}

// GetDetailOrder godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetDetailOrder
// @Accept */*
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.ResponseData{data=entity.DataOrder}
// @Router /api/v1/order/detail/:id [get]
func (s *OrderController) GetDetailOrder(c *gin.Context) {
	var person OrderDetailBindigObject
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, "ID TIDAK BOLEH KOSONG")
		return
	}
	res, err := s.uc.GetDetailOrder(c.Copy(), person.ID)
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

// GetAllOrder godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetAllOrder
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{data=[]entity.DataOrder}
// @Router /api/v1/order [get]
func (s *OrderController) GetAllOrder(c *gin.Context) {
	res, err := s.uc.GetAllOrder(c.Copy())
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

// GetAllOrderPagination godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetAllOrderPagination
// @Accept */*
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} model.ResponseData{data=[]entity.DataOrder}
// @Router /api/v1/order/paging [get]
func (s *OrderController) GetAllOrderPagination(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	res, err := s.uc.GetAllOrderWithPaging(c.Copy(), intPage, intLimit)
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

// SearchAllOrder godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags SearchAllOrder
// @Accept */*
// @Produce json
// @Param key query string true "key"
// @Success 200 {object} model.ResponseData{data=[]entity.DataOrder}
// @Router /api/v1/order/search [get]
func (s *OrderController) SearchAllOrder(c *gin.Context) {
	key := c.Query("key")
	res, err := s.uc.SearchDataOrder(c.Copy(), key)
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

// UpdateOrder godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags UpdateOrder
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{}
// @Param Payload body model.UpdateOrder true "UpdateOrder"
// @Router /api/v1/order/update [put]
func (s *OrderController) UpdateOrder(c *gin.Context) {
	var data *model.UpdateOrder
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "INVALID OBJECT")
		return
	}
	err = s.uc.UpdateOrderData(c.Copy(), data)
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

// InsertOrder godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags InsertOrder
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{}
// @Param Payload body model.InsertOrder true "InsertOrder"
// @Router /api/v1/order/add [post]
func (s *OrderController) InsertOrder(c *gin.Context) {
	var data *model.InsertOrder
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "INVALID OBJECT")
		return
	}
	err = s.uc.InsertDataOrder(c.Copy(), data)
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

// DeleteOrder godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags DeleteOrder
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{}
// @Router /api/v1/order/detail/:id [delete]
func (s *OrderController) DeleteOrder(c *gin.Context) {
	var person OrderDetailBindigObject
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, "ID TIDAK BOLEH KOSONG")
		return
	}
	err := s.uc.DeleteDataOrder(c.Copy(), person.ID)
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
