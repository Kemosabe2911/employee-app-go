package controller

import (
	"net/http"

	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func (controller *HealthController) GetHealth(c *gin.Context) {
	logger.Infof("Calling GetHealth")
	message := "Health is UP"
	c.String(http.StatusOK, message)
	logger.Infof("Called GetHealth")
}
