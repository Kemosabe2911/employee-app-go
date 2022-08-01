package service

import (
	"errors"

	"github.com/Kemosabe2911/employee-app-go/auth"
	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"github.com/Kemosabe2911/employee-app-go/repository"
	"github.com/Kemosabe2911/employee-app-go/utils"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(dto.UserSignUpRequest) *model.APIResponse
	UserLogin(dto.UserLoginRequest) *model.APIResponseWithError
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

	if ok := utils.ValidMailAddress(userData.Email); !ok {
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
			StatusCode: 400,
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

func (us *userService) UserLogin(loginData dto.UserLoginRequest) *model.APIResponseWithError {
	if ok := utils.ValidMailAddress(loginData.Email); !ok {
		return &model.APIResponseWithError{
			StatusCode: 400,
			Data:       "Invalid Email",
			Error:      errors.New("invalid email"),
		}
	}

	user, err := us.userRepository.GetUserByEmail(loginData.Email)
	if err != nil {
		logger.Error("Error while getting user")
		return &model.APIResponseWithError{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot get user",
			},
			Error: err,
		}
	}

	logger.Info(user)
	if ok := utils.CheckPasswordHash(loginData.Password, user.Password); !ok {
		return &model.APIResponseWithError{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Password doesn't match",
			},
			Error: errors.New("password not correct"),
		}
	}

	access_token, err := auth.GenerateAccessToken(loginData.Email)
	if err != nil {
		logger.Error("Error while creating Access Token")
		return &model.APIResponseWithError{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Access Token Failed",
			},
			Error: err,
		}
	}
	logger.Info(access_token)

	refresh_token, err := auth.GenerateRefreshToken(loginData.Email)
	if err != nil {
		logger.Error("Error while creating Refresh Token")
		return &model.APIResponseWithError{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Refresh Token Failed",
			},
			Error: err,
		}
	}
	logger.Info(refresh_token)

	tokens := auth.TokenStruct{
		Access:  access_token,
		Refresh: refresh_token,
	}

	logger.Info("Login user")
	return &model.APIResponseWithError{
		StatusCode: 200,
		Data:       tokens,
		Error:      nil,
	}
}
