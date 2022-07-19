package main

import (
	"fmt"

	"github.com/Kemosabe2911/employee-app-go/database"
)

func main() {
	// config := config.GetConfig()
	db, dbErr := database.GetDBConnection()
	if dbErr != nil {
		panic(dbErr)
	} else {
		fmt.Println(db)
	}
}
