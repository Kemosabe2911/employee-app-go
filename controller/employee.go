package controller

import (
	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/helpers"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/service"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	logger.Info("Start CreateEmployee in Controller")
	var employeeData dto.CreateEmployeeRequest
	if err := c.BindJSON(&employeeData); err != nil {
		c.JSON(400, helpers.InvalidRequestError)
		return
	}

	resp := ec.EmployeeService.CreateEmployee(employeeData)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End CreateEmployee in Controller")
}
