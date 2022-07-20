package database

import (
	"fmt"

	"github.com/Kemosabe2911/employee-app-go/model"

	"github.com/Kemosabe2911/employee-app-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDBConnection() (*gorm.DB, error) {
	config := config.GetConfig()

	DBurl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode = disable password= %s",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUsername,
		config.PostgresDB,
		config.PostgresPassword,
	)

	db, err := gorm.Open(postgres.Open(DBurl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.PostgresSchema + ".", // schema name
			SingularTable: false,
		},
	})
	if err != nil {
		fmt.Println("Error in Connecting to DB  %w", err)
		return nil, err
	}
	fmt.Println("DB Connection successfull")
	return db, nil
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&model.Employee{}, &model.Address{}, &model.Department{}, &model.DepartmentDetails{}, &model.Project{}, &model.Role{})
}
