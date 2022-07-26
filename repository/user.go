package repository

import (
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(model.User) (model.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func CreateUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(user model.User) (model.User, error) {
	logger.Info("Started CreateUser in Repo")
	err := ur.DB.Create(&user).Error
	logger.Info("Ended CreateUser in Repo")
	return user, err
}
