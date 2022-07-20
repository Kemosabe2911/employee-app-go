package server

import (
	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/Kemosabe2911/employee-app-go/constant"
	"github.com/Kemosabe2911/employee-app-go/controller"
	"github.com/Kemosabe2911/employee-app-go/middleware"
	"github.com/gin-gonic/gin"
)

//ApplicationRouter function to setup a new router object with the routes to be exposed and return it
func ApplicationRouter(employeeController *controller.EmployeeController, roleController *controller.RoleController) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		gin.Recovery(),
	)
	config := config.GetConfig()
	if config.Env != constant.PRD {
		// Swagger UI.
		router.Static("/swaggerui/", "./swagger-ui")
	}
	router.Use(middleware.CORSMiddleware())
	health := new(controller.HealthController)
	router.GET("/health", health.GetHealth)

	v1 := router.Group("v1")
	{
		{
			v1.POST("/role", roleController.CreateRole)
			v1.GET("/role", roleController.GetAllRoles)
			v1.GET("/role/:id", roleController.GetRoleById)
			v1.POST("/employee", employeeController.CreateEmployee)
		}
	}

	return router
}
