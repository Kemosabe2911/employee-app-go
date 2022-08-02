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
	config.LoadConfig()
	config := config.GetConfig()

	db, dbErr := database.GetDBConnection()

	if dbErr != nil {
		panic(dbErr)
	}

	database.InitialMigration(db)

	//Initialize Logger
	_, err := logger.InitLogger(config.Env)
	if err != nil {
		logger.Errorf("Error in initializing logger", "error", err)
	}

	// //Google Login
	// mux := http.NewServeMux()

	// mux.HandleFunc("/google/login", controller.GoogleLogin)
	// mux.HandleFunc("/google/callback", controller.GoogleCallback)
	// // run server
	// log.Println("started server on :: http://localhost:8080/")
	// if oops := http.ListenAndServe(":"+config.Port, mux); oops != nil {
	// 	log.Fatal(oops)
	// }

	role := &controller.RoleController{
		RoleService: service.CreateRoleService(db),
	}

	employee := &controller.EmployeeController{
		EmployeeService: service.CreateEmployeeService(db),
	}

	department := &controller.DepartmentController{
		DepartmentService: service.CreateDepartmentService(db),
	}

	user := &controller.UserController{
		UserService: service.CreateUserService(db),
	}

	router := ApplicationRouter(employee, role, department, user)
	// http.HandleFunc("/google/callback", controller.GoogleCallback)

	logger.Infof("Starting the Server at Port %s", config.Port)
	errServerStart := router.Run(":" + config.Port)
	if errServerStart != nil {
		logger.Fatalf("Error in Starting the HTTP Server, Err: %s", errServerStart.Error())
	}
}
