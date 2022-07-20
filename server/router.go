package server

import (
	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/Kemosabe2911/employee-app-go/constant"
	"github.com/Kemosabe2911/employee-app-go/controller"
	"github.com/Kemosabe2911/employee-app-go/middleware"
	"github.com/gin-gonic/gin"
)

//ApplicationRouter function to setup a new router object with the routes to be exposed and return it
func ApplicationRouter() *gin.Engine {
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

	return router
}
