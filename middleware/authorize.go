package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuthorized(adminAccessRequired bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		logger.Info(authHeader)
		if authHeader == "" {
			c.JSON(401, "Token not found")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		signingKey := []byte(config.GetConfig().JwtSecretKey)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				return nil, fmt.Errorf("invalid token %+v", token.Header["alg"])
			}
			return signingKey, nil
		})
		logger.Info(token)
		if err != nil || !token.Valid {
			c.JSON(401, "Invalid token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if adminAccessRequired {
			if claims["isAdmin"] != true {
				c.JSON(401, "Admin access required")
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}

	}
}
