package router

import (
	"dbo_backend_test/app/controller"
	"github.com/gin-gonic/gin"
)

func AdminRouter(c *gin.RouterGroup, authservice *controller.AuthService) {
	c.GET("/login-data", authservice.GetLoginData)
}
