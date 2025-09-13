package main

import (
	"log"

	"gin-grom-mysql/models"
	"gin-grom-mysql/router"
)

func main() {
	db := models.SetupDB()
	defer models.CloseDB(db)

	r := router.SetupRouter()
	// if err := r.Run(":8090"); err != nil {
	// 	log.Fatal(err)
	// }
	log.Fatal(r.Run(":8080"))
}
