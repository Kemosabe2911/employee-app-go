package repository

import (
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role model.Role, tx *gorm.DB) (model.Role, error)
	GetAllRoles() ([]model.Role, error)
	GetRoleById(id string) (model.Role, error)
}

type roleRepository struct {
	DB *gorm.DB
}

func CreateRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{
		DB: db,
	}
}

func (rr *roleRepository) CreateRole(role model.Role, tx *gorm.DB) (model.Role, error) {
	logger.Infof("Start CreateRole %+v ", role)
	err := tx.Create(&role).Error
	logger.Info("End CreateRole")
	return role, err
}

func (rr *roleRepository) GetAllRoles() ([]model.Role, error) {
	var roles []model.Role
	res := rr.DB.Find(&roles)
	if res.Error != nil {
		msg := res.Error
		return nil, msg
	}
	return roles, res.Error
}

func (rr *roleRepository) GetRoleById(id string) (model.Role, error) {
	logger.Info("Start GetRoleById")
	var role model.Role
	response := rr.DB.Where("id =?", id).First(&role)
	logger.Info(response.Error)
	if response.Error != nil {
		logger.Error("Error while fetching from role repo", response.Error.Error())
	}
	logger.Infof("End GetRoleById")
	return role, response.Error
}
