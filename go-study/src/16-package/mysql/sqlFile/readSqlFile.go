package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadSqlFile 读取sql文件内容 -> 然后再使用gorm执行sql
func ReadSqlFile(path string) string {
	sqlFile, err := filepath.Abs(path)
	if err != nil {
		panic("Error: file path error !")
	}
	// file, err := ioutil.ReadFile(sqlFile)
	file, err := os.ReadFile(sqlFile)
	if err != nil {
		panic("Error: read sql file error !")
	}
	return string(file)
}

func main() {
	fmt.Println(ReadSqlFile("./src/package/mysql/sqlFile/user.sql"))
}