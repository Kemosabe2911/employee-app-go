package controller

import (
	"net/http"

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
	logger.Info(resp.Data)
	if resp.Error != nil {
		c.JSON(resp.StatusCode, resp.Data)
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("access", resp.Data.(service.ReturnData).Token.Access, 60*60*24, "/", "http://localhost:8080", true, true)
	c.SetCookie("refresh", resp.Data.(service.ReturnData).Token.Refresh, 60*60*24, "/", "http://localhost:8080", true, true)

	logger.Info("Successfully Signed Up")
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
	if resp.Error != nil {
		c.JSON(resp.StatusCode, resp.Data)
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("access", resp.Data.(service.ReturnData).Token.Access, 60*60*24, "/", "http://localhost:8080", true, true)
	c.SetCookie("refresh", resp.Data.(service.ReturnData).Token.Refresh, 60*60*24, "/", "http://localhost:8080", true, true)

	logger.Info("Successfully Logged In")

	c.JSON(resp.StatusCode, resp.Data)
}

func (uc *UserController) LogoutUser(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("access", "", -1, "/", "http://localhost:8080", true, true)
	c.SetCookie("refresh", "", -1, "/", "http://localhost:8080", true, true)
	c.JSON(200, "Successfully Logged Out")
}
