package main

import (
	// Import the routers package
	DB "WISE_SOFTWARE/DB"
	fun "WISE_SOFTWARE/functions"
	"WISE_SOFTWARE/routers"
	custom_validation "WISE_SOFTWARE/validator"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)
	custom_validation.HcValidate()

	// Create a CONFIG.json file if it does not exist in the user's home directory
	fun.CreateConfigFile()

	// Initialize the cache
	fun.InitCache()

	// Connect to the database

	dsn := "Server=192.168.43.73\\SQLEXPRESS;Database=seatDB7;User Id=root;Password="
	DB.ConnectDatabase(dsn)
	if DB.Conn != nil {
		routers.NewRoutes().Run("127.0.0.1:8080")
	} else {
		fmt.Println("We Could Not Connect To Database")
	}
}
