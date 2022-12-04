package router

import (
	"dbo_backend_test/app/controller"
	"github.com/gin-gonic/gin"
)

func ProductRouter(c *gin.RouterGroup, authservice *controller.ProductService) {
	c.GET("/", authservice.GetAllProduct)
	c.POST("/produk/add", authservice.AddProduct)
}
