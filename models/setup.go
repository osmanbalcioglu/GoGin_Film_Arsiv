package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("films.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = database.AutoMigrate(&Film{})
	if err != nil {
		panic(err.Error())
	}

	DB = database
}
