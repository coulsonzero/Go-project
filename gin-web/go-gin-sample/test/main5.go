package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	// dir, _ := os.Getwd()
	sqlFile, _ := filepath.Abs("./test.sql")
	file, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		log.Fatal("read file error !")
	}
	fmt.Println(string(file))
}
