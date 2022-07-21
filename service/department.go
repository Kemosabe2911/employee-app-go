package service

import (
	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"github.com/Kemosabe2911/employee-app-go/repository"
	"gorm.io/gorm"
)

type DepartmentService interface {
	CreateDepartment(createDepartmentDto dto.CreateDepartment) *model.APIResponse
}

type departmentService struct {
	departmentRepository repository.DepartmentRepository
	DB                   *gorm.DB
}

func CreateDepartmentService(db *gorm.DB) *departmentService {
	return &departmentService{
		departmentRepository: repository.CreateDepartmentRepository(db),
		DB:                   db,
	}
}

func (ds *departmentService) CreateDepartment(createDepartmentDto dto.CreateDepartment) *model.APIResponse {
	logger.Infof("Start CreateDepartment %+v", createDepartmentDto)

	tx := ds.DB.Begin()

	departmentDetails := model.DepartmentDetails{
		DepartmentRoom: createDepartmentDto.DepartmentRoom,
		DepartmentCode: createDepartmentDto.DepartmentCode,
		Website:        createDepartmentDto.Website,
	}

	departmentDetails, err := ds.departmentRepository.CreateDepartmentDetails(departmentDetails, tx)
	if err != nil {
		logger.Error("Error while inserting department details")
		tx.Rollback()
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot save department details",
			},
		}
	}

	department := model.Department{
		Name:       createDepartmentDto.Name,
		Department: departmentDetails,
	}

	department, err = ds.departmentRepository.CreateDepartment(department, tx)
	if err != nil {
		logger.Error("Error while creating department", err.Error())
		tx.Rollback()
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Unable to create department",
			},
		}
	}
	tx.Commit()

	logger.Infof("End CreateDepartment %+v", department)
	return &model.APIResponse{
		StatusCode: 201,
		Data:       department,
	}
}
