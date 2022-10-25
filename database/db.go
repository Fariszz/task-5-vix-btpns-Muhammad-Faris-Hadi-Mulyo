package database

import (
	"GOlangRakamin/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", configuration.DB_HOST, configuration.DB_PORT, configuration.DB_USERNAME, configuration.DB_NAME, configuration.DB_PASSWORD)
	db, err = gorm.Open(postgres.Open(connect_string), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
}

func DbManager() *gorm.DB {
	return db
}
