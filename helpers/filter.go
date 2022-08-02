package helpers

import (
	"github.com/gin-gonic/gin"
)

func GetFilterValue(c *gin.Context) string {
	search := c.Query("search")
	return search
}
