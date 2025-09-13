package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func SetupDB() *gorm.DB {
	// load ini file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)

	// Gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置全局表名禁用复数
		},
	})

	if err != nil {
		log.Fatal("Failed to create a connection to database")
	}
	fmt.Println("连接mysql数据库成功")
	DB = db
	DB.AutoMigrate(&User{})
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()

	dbSQL.SetMaxIdleConns(10)                // 设置空闲连接池中连接的最大数量
	dbSQL.SetMaxOpenConns(100)               // 设置打开数据库连接的最大数量
	dbSQL.SetConnMaxLifetime(time.Hour * 24) // 设置了连接可复用的最大时间

	if err != nil {
		log.Fatal("Failed to close connection from database")
	}
	dbSQL.Close()
}
