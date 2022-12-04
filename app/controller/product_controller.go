package controller

import (
	"dbo_backend_test/app/domain/model"
	"dbo_backend_test/app/usecase/product_uc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductService struct {
	uc product_uc.ProductRepo
}

func NewProductUc(uc product_uc.ProductRepo) *ProductService {
	return &ProductService{
		uc: uc,
	}
}

// GetAllProduct godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetAllProduct
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{data=[]entity.ProductData}
// @Router /admin/produk [get]
func (s *ProductService) GetAllProduct(c *gin.Context) {
	res, err := s.uc.GetProduct(c.Copy())
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

// AddProduct godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags AddProduct
// @Accept */*
// @Produce json
// @Success 200 {object} model.ResponseData{}
// @Param Payload body model.AddProduct true "AddProduct"
// @Router /admin/produk/add [post]
func (s *ProductService) AddProduct(c *gin.Context) {
	var data *model.AddProduct
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "INVALID OBJECT")
		return
	}
	err = s.uc.AddProduct(c.Copy(), data)
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
