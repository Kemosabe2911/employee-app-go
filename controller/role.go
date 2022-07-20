package controller

import (
	"net/http"

	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/service"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	RoleService service.RoleService
}

func (rc *RoleController) CreateRole(c *gin.Context) {
	logger.Info("Start CreateRole in Controller")
	var createRoleDto dto.CreateRole
	if err := c.BindJSON(&createRoleDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := rc.RoleService.CreateRole(createRoleDto)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End CreateRole in Controller")
}

func (rc *RoleController) GetAllRoles(c *gin.Context) {
	logger.Info("Start getAllRoles in Controller")
	resp := rc.RoleService.GetAllRoles()
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End getAllRoles in Controller")
}

func (rc *RoleController) GetRoleById(c *gin.Context) {
	id := c.Param("id")
	logger.Info("Start GetRoleById in Controller")
	resp := rc.RoleService.GetRoleById(id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetRoleById in Controller")
}
