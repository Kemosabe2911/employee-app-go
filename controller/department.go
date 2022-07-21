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

func (dc *DepartmentController) GetAllDepartments(c *gin.Context) {
	logger.Info("Start GetAllDepartments in Controller")
	resp := dc.DepartmentService.GetAllDepartments()
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetAllDepartments in Controller")
}

func (dc *DepartmentController) GetDepartmentById(c *gin.Context) {
	id := c.Param("id")
	logger.Info("Start GetDepartmentById in Controller")
	resp := dc.DepartmentService.GetDepartmentById(id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetDepartmentById in Controller")
}

func (dc *DepartmentController) UpdateDepartment(c *gin.Context) {
	logger.Info("Start UpdateDepartment in Controller")
	var UpdateDepartmentDto dto.UpdateDepartment
	if err := c.BindJSON(&UpdateDepartmentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	resp := dc.DepartmentService.UpdateDepartment(UpdateDepartmentDto, id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End UpdateDepartment in Controller")
}
