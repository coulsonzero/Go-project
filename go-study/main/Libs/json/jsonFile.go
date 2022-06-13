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
	decoder := json.NewDecoder(f)
	decoder.Decode(&bar)
	// fmt.Println(bar)

	data, _ := json.Marshal(bar)
	fmt.Printf("%T\n", bar)
	fmt.Println(string(data))
}
