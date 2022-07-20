package service

import (
	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"github.com/Kemosabe2911/employee-app-go/repository"
	"gorm.io/gorm"
)

type EmployeeService interface {
	CreateEmployee(employeeRequest dto.CreateEmployeeRequest) *model.APIResponse
	GetAllEmployees() *model.APIResponse
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
	DB                 *gorm.DB
}

func CreateEmployeeService(db *gorm.DB) *employeeService {
	return &employeeService{
		employeeRepository: repository.CreateEmployeeRepository(db),
		DB:                 db,
	}
}

func (es *employeeService) CreateEmployee(employeeRequest dto.CreateEmployeeRequest) *model.APIResponse {
	logger.Info("Start CreateEmployee in Service")
	employee := model.Employee{
		Name:         employeeRequest.Name,
		Username:     employeeRequest.Username,
		Password:     employeeRequest.Password,
		Age:          employeeRequest.Age,
		IsActive:     true,
		DepartmentID: employeeRequest.DepartmentID,
		RoleID:       employeeRequest.RoleID,
		AddressID:    employeeRequest.AddressID,
	}
	logger.Info(employee)

	employee, err := es.employeeRepository.CreateEmployee(employee)
	if err != nil {
		logger.Error("Error while saving product")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot save employee",
			},
		}
	}
	logger.Info("Saved employee")
	return &model.APIResponse{
		StatusCode: 201,
		Data:       employee,
	}
}

func (es *employeeService) GetAllEmployees() *model.APIResponse {
	logger.Info("Start GetAllEmployees in Service")
	employee, err := es.employeeRepository.GetAllEmployees()
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employees not found",
			},
		}
	}
	logger.Info("End GetAllEmployees in Service")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       employee,
	}
}
