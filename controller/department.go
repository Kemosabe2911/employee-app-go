package controller

import (
	"net/http"

	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/service"
	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	DepartmentService service.DepartmentService
}

func (dc *DepartmentController) CreateDepartment(c *gin.Context) {
	logger.Info("Start CreateDepartment in Controller")
	var createDepartmentDto dto.CreateDepartment
	if err := c.BindJSON(&createDepartmentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := dc.DepartmentService.CreateDepartment(createDepartmentDto)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End CreateDepartment in Controller")
}
