package DATABASE

import (
	Query "WISE_SOFTWARE/DB/db_query"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Conn *gorm.DB
var Q = Query.Q

func ConnectDatabase(dsn string) {
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	sqlDB, _ := database.DB()
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	Conn = database

	Query.SetDefault(Conn)
}
