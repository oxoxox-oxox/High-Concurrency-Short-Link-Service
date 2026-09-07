package model

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fail to connect to database: %v", err)
	}

	DB.AutoMigrate(&Link{})
	log.Println("Success to connect to the database, the structure of table is syncronised now!")
}