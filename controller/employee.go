package controller

import (
	"net/http"
	"path/filepath"

	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/helpers"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	logger.Info("Start CreateEmployee in Controller")
	var employeeData dto.CreateEmployeeRequest
	if err := c.BindJSON(&employeeData); err != nil {
		logger.Error(err)
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

func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	logger.Info("Start UpdateEmployee - COntroller")
	id := c.Param("id")
	var employeeData dto.UpdateEmployeeRequest
	if err := c.BindJSON(&employeeData); err != nil {
		c.JSON(400, helpers.InvalidRequestError)
		return
	}
	resp := ec.EmployeeService.UpdateEmployee(id, employeeData)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End UpdateEmployee - Controller")
}

func (ec *EmployeeController) UploadIdProof(c *gin.Context) {
	logger.Info("Start UploadIdProof in Controller")
	id := c.Param("id")
	file, err := c.FormFile("file")

	logger.Info(file.Filename)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "./assets/"+newFileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		logger.Info(err)
		return
	}

	resp := ec.EmployeeService.UploadIdProof(id, newFileName)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End UploadIdProof in Controller")
}
