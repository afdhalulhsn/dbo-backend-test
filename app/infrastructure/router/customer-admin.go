package router

import (
	"dbo_backend_test/app/controller"
	"github.com/gin-gonic/gin"
)

func CustomerAdmin(c *gin.RouterGroup, authservice *controller.CustomerController) {
	c.GET("/detail/:id", authservice.GetDetailCustomer)
	c.GET("/", authservice.GetAllCustomer)
	c.GET("/paging", authservice.GetAllCustomerPagination)
	c.GET("/search", authservice.SearchAllCustomer)
	c.DELETE("/detail/:id", authservice.DeleteCustomer)

}

func Customer(c *gin.RouterGroup, authservice *controller.CustomerController) {
	c.GET("/detail/:id", authservice.GetDetailCustomer)
	c.GET("/update", authservice.UpdateCustomer)
}
