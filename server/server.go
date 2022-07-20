package server

import (
	"github.com/Kemosabe2911/employee-app-go/controller"
	"github.com/Kemosabe2911/employee-app-go/service"

	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/Kemosabe2911/employee-app-go/database"
	"github.com/Kemosabe2911/employee-app-go/logger"
)

//Start function to initialize the server and begins to listen at the configured port
func Start() {
	config := config.GetConfig()
	db, dbErr := database.GetDBConnection()

	if dbErr != nil {
		panic(dbErr)
	}

	// database.InitialMigration(db)

	//Initialize Logger
	_, err := logger.InitLogger(config.Env)
	if err != nil {
		logger.Errorf("Error in initializing logger", "error", err)
	}

	role := &controller.RoleController{
		RoleService: service.CreateRoleService(db),
	}

	employee := &controller.EmployeeController{
		EmployeeService: service.CreateEmployeeService(db),
	}

	router := ApplicationRouter(employee, role)

	logger.Infof("Starting the Server at Port %s", config.Port)
	errServerStart := router.Run(":" + config.Port)
	if errServerStart != nil {
		logger.Fatalf("Error in Starting the HTTP Server, Err: %s", errServerStart.Error())
	}
}
