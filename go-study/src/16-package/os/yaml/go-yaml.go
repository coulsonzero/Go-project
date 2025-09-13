package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// 初始化yml配置
	viper.SetConfigType("yml")
	viper.SetConfigFile("./src/conf/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// 获取yml配置项
	host := viper.GetString("mysql.host")
	fmt.Println(host)
}
