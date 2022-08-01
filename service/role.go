package service

import (
	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/helpers"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"github.com/Kemosabe2911/employee-app-go/repository"
	"gorm.io/gorm"
)

type RoleService interface {
	CreateRole(createRoleDto dto.CreateRole) *model.APIResponse
	GetAllRoles() *model.APIResponse
	GetRoleById(id string) *model.APIResponse
}

type roleService struct {
	roleRepository repository.RoleRepository
	DB             *gorm.DB
}

func CreateRoleService(db *gorm.DB) *roleService {
	return &roleService{
		roleRepository: repository.CreateRoleRepository(db),
		DB:             db,
	}
}

func (rs *roleService) CreateRole(createRoleDto dto.CreateRole) *model.APIResponse {
	logger.Infof("Start CreateRole %+v", createRoleDto)
	role := model.Role{
		Role: createRoleDto.Role,
	}

	tx := rs.DB.Begin()

	role, err := rs.roleRepository.CreateRole(role, tx)
	if err != nil {
		logger.Error("Error while creating role", err.Error())
		tx.Rollback()
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Unable to create role",
			},
		}
	}
	tx.Commit()

	logger.Infof("End CreateRole %+v", role)
	return &model.APIResponse{
		StatusCode: 201,
		Data:       role,
	}
}

func (rs *roleService) GetAllRoles() *model.APIResponse {
	logger.Info("Start GetAllRoles")
	roles, err := rs.roleRepository.GetAllRoles()
	if err != nil {
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: helpers.ErrRoleNotFoundError.Error(),
			},
		}
	}
	logger.Infof("End GetAllRoles count %d", len(roles))
	return &model.APIResponse{
		StatusCode: 200,
		Data:       roles,
	}
}

func (rs *roleService) GetRoleById(id string) *model.APIResponse {
	logger.Info("Start GetRoleById")
	role, err := rs.roleRepository.GetRoleById(id)
	if err != nil {
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: helpers.ErrRoleNotFoundError.Error(),
			},
		}
	}
	logger.Infof("End GetRoleById")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       role,
	}
}
