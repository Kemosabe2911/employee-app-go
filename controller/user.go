package controller

import (
	"github.com/Kemosabe2911/employee-app-go/auth"
	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"github.com/Kemosabe2911/employee-app-go/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userData dto.UserSignUpRequest
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(400, "Error while binding")
		return
	}
	resp := uc.UserService.CreateUser(userData)
	c.JSON(resp.StatusCode, resp.Data)
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var userData dto.UserLoginRequest
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(400, "Error while binding")
		return
	}
	resp := uc.UserService.UserLogin(userData)
	logger.Info(resp.Data)
	email_id := resp.Data.(model.User).Email
	access_token, err := auth.GenerateAccessToken(email_id)
	if err != nil {
		logger.Error("Error while creating Access Token")
		c.JSON(500, "Access Token Failed")
	}
	logger.Info(access_token)

	refresh_token, err := auth.GenerateRefreshToken(email_id)
	if err != nil {
		logger.Error("Error while creating Refresh Token")
		c.JSON(500, "Refresh Token Failed")
	}
	logger.Info(refresh_token)

	c.SetCookie("access", access_token, 60*60*24, "/", "localhost", false, true)
	c.SetCookie("refresh", refresh_token, 60*60*24, "/", "localhost", false, true)

	c.JSON(resp.StatusCode, resp.Data)
}
