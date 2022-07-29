package controller

import (
	"github.com/Kemosabe2911/employee-app-go/auth"
	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
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
	// if resp.StatusCode == 404 || resp.StatusCode == 400 {
	// 	c.JSON(resp.StatusCode, resp.Data)
	// 	return
	// }
	c.SetCookie("access", resp.Data.(auth.TokenStruct).Access, 60*60*24, "/", "localhost", false, true)
	c.SetCookie("refresh", resp.Data.(auth.TokenStruct).Refresh, 60*60*24, "/", "localhost", false, true)

	logger.Info("Successfully Logged In")

	c.JSON(resp.StatusCode, resp.Data)
}
