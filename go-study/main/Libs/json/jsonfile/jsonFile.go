package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bar struct {
	xAxis  map[string]interface{}
	yAxis  map[string]interface{}
	series map[string]interface{}
}

func main() {
	ReadJsonFile()
}

func ReadJsonFile() {
	f, _ := os.Open("bar.json")
	defer f.Close()
	var bar map[string]interface{}
	// json File -> object
	json.NewDecoder(f).Decode(&bar)
	fmt.Println(bar)

	// object -> jsonStr
	data, _ := json.Marshal(bar)
	fmt.Println(string(data))
}

// ReadJsonFile2 json file -> object -> json String
func ReadJsonFile2(jsonfile string, obj map[string]interface{}) string {
	// 打开json文件
	f, _ := os.Open(jsonfile)
	defer f.Close()

	// json文件 -> struct
	// var bar map[string]interface{}
	json.NewDecoder(f).Decode(&obj)
	// fmt.Println(bar)

	// struct -> json
	data, _ := json.Marshal(obj)
	return string(data)
}
