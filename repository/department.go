package repository

import (
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CreateDepartment(department model.Department, tx *gorm.DB) (model.Department, error)
	CreateDepartmentDetails(departmentDetails model.DepartmentDetails, tx *gorm.DB) (model.DepartmentDetails, error)
	GetAllDepartments() ([]model.Department, error)
	GetDepartmentById(id string) (model.Department, error)
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

func (dr *departmentRepository) GetAllDepartments() ([]model.Department, error) {
	var departments []model.Department
	res := dr.DB.Preload("Department").Find(&departments)
	if res.Error != nil {
		msg := res.Error
		return nil, msg
	}
	return departments, res.Error
}

func (dr *departmentRepository) GetDepartmentById(id string) (model.Department, error) {
	logger.Info("Start GetDepartmentById")
	var department model.Department
	response := dr.DB.Where("id =?", id).Preload("Department").First(&department)
	logger.Info(response.Error)
	if response.Error != nil {
		logger.Error("Error while fetching from department repo", response.Error.Error())
	}
	logger.Infof("End GetDepartmentById")
	return department, response.Error
}
