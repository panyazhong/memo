package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Id uint `json:"user_id"`
	jwt.StandardClaims
}

const TokenExpiresDuration = time.Hour * 2
var CustomSecret = []byte("memo")

func GenerToken(id uint) (string, error) {
	c := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpiresDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(CustomSecret)
}

func ParseToken(tokenString string) (*Claims, error){
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(tokenSting *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})

	if (err != nil) {
		return nil, err
	}

	if c, ok := token.Claims.(*Claims); ok && token.Valid {
		return c, nil
	}

	return nil, errors.New("invalid token")
}

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	claims, err := ParseToken(token)

	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("user_id", claims.Id)
	c.Next()
}