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

func (ec *EmployeeController) GetAllEmployees(c *gin.Context) {
	logger.Info("Start GetAllEmployees in Controller")
	resp := ec.EmployeeService.GetAllEmployees()
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetAllEmployees in Controller")
}

func (ec *EmployeeController) GetEmployeeById(c *gin.Context) {
	logger.Info("Start GetEmployeeById - Controller")
	id := c.Param("id")
	resp := ec.EmployeeService.GetEmployeeById(id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetEmployeeById - Controller")
}

func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {
	logger.Info("Start DeleteEmployee - Controller")
	id := c.Param("id")
	resp := ec.EmployeeService.DeleteEmployee(id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End DeleteEmployee - Controller")
}
