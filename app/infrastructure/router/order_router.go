package router

import (
	"dbo_backend_test/app/controller"
	"github.com/gin-gonic/gin"
)

func OrderRouter(c *gin.RouterGroup, authservice *controller.OrderController) {
	c.GET("/detail/:id", authservice.GetDetailOrder)
	c.GET("/", authservice.GetAllOrder)
	c.GET("/paging", authservice.GetAllOrderPagination)
	c.GET("/search", authservice.SearchAllOrder)
	c.PUT("/update", authservice.UpdateOrder)
	c.POST("/add", authservice.InsertOrder)
	c.DELETE("/detail/:id", authservice.DeleteOrder)
}
