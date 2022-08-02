package server

import (
	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/Kemosabe2911/employee-app-go/constant"
	"github.com/Kemosabe2911/employee-app-go/controller"
	"github.com/Kemosabe2911/employee-app-go/middleware"
	"github.com/gin-gonic/gin"
)

//ApplicationRouter function to setup a new router object with the routes to be exposed and return it
func ApplicationRouter(employeeController *controller.EmployeeController, roleController *controller.RoleController, departmentController *controller.DepartmentController, UserController *controller.UserController) *gin.Engine {
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
			v1.GET("/employee", middleware.IsAuthorized(), employeeController.GetAllEmployees)
			v1.GET("/employee/:id", employeeController.GetEmployeeById)
			v1.POST("/department", departmentController.CreateDepartment)
			v1.GET("/department", departmentController.GetAllDepartments)
			v1.GET("/department/:id", departmentController.GetDepartmentById)
			v1.DELETE("/employee/:id", employeeController.DeleteEmployee)
			v1.PUT("/department/:id", departmentController.UpdateDepartment)
			v1.PUT("/employee/:id", employeeController.UpdateEmployee)
			v1.POST("/signup", UserController.CreateUser)
			v1.POST("/login", UserController.LoginUser)
			v1.GET("/logout", UserController.LogoutUser)
			// v1.GET("/google/login", controller.GoogleLogin)
			// v1.GET("/google/callback", controller.GoogleCallback)
			v1.PATCH("/employee/id-proof/:id", employeeController.UploadIdProof)
		}
	}

	return router
}
