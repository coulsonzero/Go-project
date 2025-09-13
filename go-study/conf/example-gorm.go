package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置全局表名禁用复数
		},
	})
	if err != nil {
		log.Fatal("Failed to create a connection to database")
	}
	fmt.Println("连接mysql数据库成功")

	// 迁移 schema
	db.AutoMigrate(&Product{})

	var product []Product

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})
	// db.Exec(`insert into product(code, price, created_at) values("D36", 80, ?)`, time.Now().Format(time.DateTime))
	// db.Raw(`insert into product(code, price) values("C45", 100)`).Scan(&product)

	// Read
	db.Raw("select * from product").Scan(&product)
	fmt.Println(product)

	// Change
	// db.Exec(`update product set code = "B70", price = 80, updated_at = ? where id = 7`, time.Now().Format(time.DateTime))

	// Delete
	// db.Exec("delete from product")
}
