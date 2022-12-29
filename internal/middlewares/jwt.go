package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Jwt(key []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := validate(key, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func validate(key []byte, c *gin.Context) error {
	t, err := c.Cookie("token")
	if err != nil || t == "" {
		return errors.New("empty token")
	}

	_, err = jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return errors.New("token expired")
		default:
			return errors.New("invalid token")
		}
	}

	return nil
}
