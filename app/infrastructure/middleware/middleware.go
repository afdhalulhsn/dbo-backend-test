package middleware

import (
	"dbo_backend_test/app/helper/jwt_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt_token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtAuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt_token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "ANDA TIDAK DIIJINKAN UNTUK MENGKAKSES API INI,PASTIKAN TOKEN ANDA BENAR")
			c.Abort()
			return
		}
		user_id, err := jwt_token.ExtractTokenID(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if user_id.Role != "ADMIN" {
			c.String(http.StatusUnauthorized, "MENU INI HANYA UNTUK ADMIN")
			c.Abort()
			return
		}
		c.Next()
	}
}
