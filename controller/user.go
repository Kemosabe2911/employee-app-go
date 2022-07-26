package controller

import (
	"github.com/Kemosabe2911/employee-app-go/dto"
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
