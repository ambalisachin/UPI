package config

import (
	"BBT/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Credentials struct can be used to store credentials in a single data type.
type Credentials struct {
	Username string
	Password string
	Server   string
	Dbname   string
}

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/bbtt?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate will ensure the tables exist and are up-to-date.

	db.AutoMigrate(&models.FetchBillRequest{})

	return db
}
