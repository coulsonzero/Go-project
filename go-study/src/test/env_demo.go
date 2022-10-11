package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg"
	"log"
)

func main() {
	envArr := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	DSN := "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	dsn, err := gopkg.ConfigEnv(".env", envArr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dsn)
	fmt.Println(dsn == DSN)

}
