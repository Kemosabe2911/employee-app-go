package middleware

import (
	"net/http"

	"github.com/Kemosabe2911/employee-app-go/auth"
	"github.com/Kemosabe2911/employee-app-go/logger"

	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {

		accessToken, err := c.Cookie("access")
		if err != nil {
			c.JSON(401, "access token not found")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		_, err = auth.ValidateToken(accessToken)

		if err != nil {

			refreshToken, err1 := c.Cookie("refresh")
			if err1 != nil {
				c.JSON(401, "refresh token not found")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			logger.Info("got refresh token")

			email, err1 := auth.ValidateToken(refreshToken)
			if err1 != nil {
				c.JSON(401, gin.H{"error": err1.Error()})
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			logger.Info("valid refresh token")

			newAccessToken, err1 := auth.GenerateAccessToken(email)
			if err1 != nil {
				c.JSON(500, "error while creating new access token")
				c.AbortWithStatus(500)
				return
			}
			logger.Info("new acces token generated")

			c.SetCookie("access", newAccessToken, 60*60*24*90, "/", "localhost", false, true)
		}

	}
}
