package jwt_token

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string
	Role     string
}

type TokenResponse struct {
	Token string
}
type ClaimTokenResponnse struct {
	Username string
	Role     string
	Expired  string
}

var SecretKey = []byte("backend-testing-dbo-jwt")

func GenerateToken(request TokenRequest) (*TokenResponse, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_name"] = request.Username
	claims["role"] = request.Role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sec, err := token.SignedString(SecretKey)
	if err != nil {
		return nil, err
	}
	return &TokenResponse{Token: sec}, nil

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (*ClaimTokenResponnse, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		out := &ClaimTokenResponnse{
			Username: fmt.Sprintf("%v", claims["user_name"]),
			Role:     fmt.Sprintf("%v", claims["role"]),
		}
		return out, nil
	}
	return nil, nil
}
