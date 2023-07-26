package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type MysqlConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadMysqlConf() (Conf *MysqlConf) {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/conf/" + "mysql.json")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	fmt.Printf("%v", LoadMysqlConf())
}
