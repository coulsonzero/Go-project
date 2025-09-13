package model

import (
	"go-gin-todolist/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func SetupDB() *gorm.DB {
	dsn := conf.LoadConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to create a connection to database")
	}
	DB = db

	DB.AutoMigrate(&User{})
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close connection from database")
	}
	dbSQL.Close()
}
