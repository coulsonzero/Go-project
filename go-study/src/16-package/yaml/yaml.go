package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// 初始化yml配置
	InitConfigYml("./config.yaml")

	// 获取yml配置项
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	dbname := viper.GetString("mysql.dbname")
	fmt.Println(host, port, username, password, dbname)
}

func InitConfigYml(fileName string) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
