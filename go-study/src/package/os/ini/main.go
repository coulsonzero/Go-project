package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	file, err := ini.Load("config.ini")
	if err != nil {
		panic("Failed to load ini file")
	}

	Db := file.Section("mysql").Key("Db").String()
	DbHost := file.Section("mysql").Key("DbHost").String()

	fmt.Printf("database: %s, host: %s \n", Db, DbHost)
}
