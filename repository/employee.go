package repository

import (
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(model.Employee) (model.Employee, error)
	GetAllEmployees() ([]model.Employee, error)
}

type employeeRepository struct {
	DB *gorm.DB
}

func CreateEmployeeRepository(db *gorm.DB) *employeeRepository {
	return &employeeRepository{
		DB: db,
	}
}

func (er *employeeRepository) CreateEmployee(employee model.Employee) (model.Employee, error) {
	logger.Info("Start CreateEmployee in Repo")
	err := er.DB.Create(&employee).Error
	logger.Info("End CreateEmployee in Repo")
	return employee, err
}

func (er *employeeRepository) GetAllEmployees() ([]model.Employee, error) {
	logger.Info("Start GetAllEmployees in Repo")
	var employee []model.Employee
	err := er.DB.Find(&employee).Error
	logger.Info("End GetAllEmployees in Repo")
	return employee, err
}
