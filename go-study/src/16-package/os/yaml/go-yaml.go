package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// 初始化yml配置
	InitConfigYml("./src/conf/config.yml")

	// 获取yml配置项
	host := viper.GetString("mysql.host")
	fmt.Println(host)
}

func InitConfigYml(fileName string) {
	viper.SetConfigType("yml")
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
