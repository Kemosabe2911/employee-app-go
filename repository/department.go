package repository

import (
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CreateDepartment(department model.Department, tx *gorm.DB) (model.Department, error)
	CreateDepartmentDetails(departmentDetails model.DepartmentDetails, tx *gorm.DB) (model.DepartmentDetails, error)
}

type departmentRepository struct {
	DB *gorm.DB
}

func CreateDepartmentRepository(db *gorm.DB) *departmentRepository {
	return &departmentRepository{
		DB: db,
	}
}

func (dr *departmentRepository) CreateDepartment(department model.Department, tx *gorm.DB) (model.Department, error) {
	logger.Infof("Start CreateDepartment %+v ", department)
	err := tx.Create(&department).Error
	logger.Info("End CreateDepartment")
	return department, err
}

func (dr *departmentRepository) CreateDepartmentDetails(departmentDetails model.DepartmentDetails, tx *gorm.DB) (model.DepartmentDetails, error) {
	logger.Infof("Start CreateDepartmentDetails %+v ", departmentDetails)
	err := tx.Create(&departmentDetails).Error
	logger.Info("End CreateDepartmentDetails")
	return departmentDetails, err
}
