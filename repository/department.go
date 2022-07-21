package repository

import (
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CreateDepartment(department model.Department, tx *gorm.DB) (model.Department, error)
	GetAllDepartments() ([]model.Department, error)
	GetDepartmentById(id string) (model.Department, error)
	UpdateDepartment(updatedDepartment model.Department, id string, tx *gorm.DB) (model.Department, error)
	UpdateDepartmentDetails(updatedDepartmentDetails model.DepartmentDetails, id string, tx *gorm.DB) (model.DepartmentDetails, error)
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

func (dr *departmentRepository) UpdateDepartment(updatedDepartment model.Department, id string, tx *gorm.DB) (model.Department, error) {
	logger.Infof("Start UpdateDepartment %+v ", updatedDepartment)
	var department model.Department
	if err := dr.DB.Where("id = ?", id).Preload("Department").First(&department).Error; err != nil {
		logger.Error("Error while fetching from department repo", err.Error())
		return model.Department{}, err
	}
	err := tx.Model(&department).Updates(&updatedDepartment).Preload("Department").First(&department).Error
	logger.Info("End UpdateDepartment")
	return department, err
}

func (dr *departmentRepository) UpdateDepartmentDetails(updatedDepartmentDetails model.DepartmentDetails, id string, tx *gorm.DB) (model.DepartmentDetails, error) {
	logger.Infof("Start UpdateDepartmentDetails %+v ", updatedDepartmentDetails)
	var departmentDetails model.DepartmentDetails
	if err := dr.DB.Where("id = ?", id).First(&departmentDetails).Error; err != nil {
		logger.Error("Error while fetching from department details repo", err.Error())
		return model.DepartmentDetails{}, err
	}
	err := tx.Model(&departmentDetails).Updates(&updatedDepartmentDetails).Error
	logger.Info("End UpdateDepartmentDetails")
	return departmentDetails, err
}
