package database

import (
	"main/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root@/dimp_db"), &gorm.Config{})

	if err != nil {
		panic("Could not establish connection to MySQL database")
	}

	connection.AutoMigrate(&model.User{})

	DB = connection
}
