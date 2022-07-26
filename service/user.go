package service

import (
	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"github.com/Kemosabe2911/employee-app-go/repository"
	"github.com/Kemosabe2911/employee-app-go/utils"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(dto.UserSignUpRequest) *model.APIResponse
}

type userService struct {
	userRepository repository.UserRepository
	DB             *gorm.DB
}

func CreateUserService(db *gorm.DB) *userService {
	return &userService{
		userRepository: repository.CreateUserRepository(db),
		DB:             db,
	}
}

func (us *userService) CreateUser(userData dto.UserSignUpRequest) *model.APIResponse {
	if ok := utils.VerfityPassword(userData.Password, userData.ConfirmPassword); !ok {
		return &model.APIResponse{
			StatusCode: 400,
			Data:       "Passwords doesn't match",
		}
	}

	hashedPassword, ok := utils.HashPassword(userData.Password)
	if ok != nil {
		return &model.APIResponse{
			StatusCode: 500,
			Data:       "Error while hashing password",
		}
	}

	if _, ok := utils.ValidMailAddress(userData.Email); !ok {
		return &model.APIResponse{
			StatusCode: 400,
			Data:       "Invalid Email",
		}
	}

	user := model.User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Password:  hashedPassword,
	}
	logger.Info(user)

	user, err := us.userRepository.CreateUser(user)
	if err != nil {
		logger.Error("Error while creating user")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot save user",
			},
		}
	}
	logger.Info("Saved user")
	return &model.APIResponse{
		StatusCode: 201,
		Data:       user,
	}
}
