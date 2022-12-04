package router

import (
	"dbo_backend_test/app/controller"
	"github.com/gin-gonic/gin"
)

func AuthenticationTouter(c *gin.RouterGroup, authservice *controller.AuthService) {
	c.POST("/pendaftaran", authservice.Pendaftaran)
	c.POST("/login", authservice.Login)
}
