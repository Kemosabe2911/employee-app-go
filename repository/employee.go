package repository

import (
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(model.Employee) (model.Employee, error)
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
