package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var cnt int
var specifiedType []string = []string{"txt", "py", "js"}
var retainType []string = []string{"go"}

func main() {
	WriteFile()
}

// 创建文件
func CreateFile() {
	// 创建文件
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(file.Name())
	}
}

// 删除文件
func RemoveFile() {
	err := os.Remove("file.txt")
	if err != nil {
		fmt.Println(err)
	}
}

// 重命名文件
func RenameFile() {
	CreateFile()
	err := os.Rename("file.txt", "file.go")
	if err != nil {
		fmt.Println(err)
	}
}

// 读取文件
func ReadFile() {
	file, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(file))
	}
}

// 重写文件
func WriteFile() {
	s := `
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}

body {
    background-color: black;
}
`
	os.WriteFile("file.txt", []byte(s), os.ModePerm)
}

// 读取目录
func ReadDir() {
	dir, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, file := range dir {
			fmt.Println(file.Name())
		}
	}
}

// 创建目录
func CreateDir() {
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

// 创建多级目录
func CreateDirAll() {
	err := os.MkdirAll("test/demo", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

// 删除目录
func RemoveDir() {
	err := os.RemoveAll("test")
	if err != nil {
		log.Fatal(err)
	}
}

// 获取当前工作目录
func Getwd() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

// 切换当前工作目录
func Chdir(path string) {
	os.Chdir(path)
}

// 获取临时目录？
func TempDir() {
	dir := os.TempDir()
	fmt.Println("dir: ", dir)
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 获取当前目录下的所有文件或目录信息（包含当前工作目录）
func WalkDir() {
	pwd, _ := os.Getwd()
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {

		cnt++
		fmt.Printf("%-4d: %s\n", cnt, path)
		// fmt.Println(path) // 打印path信息
		// fmt.Println(info.IsDir())
		// fmt.Println(info.Name()) // 打印文件或目录名
		return nil
	})
}

func WalkDemo() {
	fmt.Println("/********** Start *********/")
	cur_dir, _ := os.Getwd()
	Walk(cur_dir)
	fmt.Printf("Total count of files: %d.", cnt)
}

// 获取当前目录下的所有文件或目录信息
func Walk(path string) {
	dir, _ := os.ReadDir(path)
	for _, file := range dir {
		file_path := filepath.Join(path, file.Name())
		cnt++
		fmt.Printf("%-4d: %s\n", cnt, file_path)
		if file.IsDir() {
			Walk(file_path)
		}
	}
}

// 遍历目录下多有文件(仅文件，不包含目录)
func WalkFiles(path string) {
	dir, _ := os.ReadDir(path)
	for _, file := range dir {
		file_path := filepath.Join(path, file.Name())
		if file.IsDir() {
			WalkFiles(file_path)
		} else {
			cnt++
			fmt.Printf("%-4d: %s\n", cnt, file_path)
		}
	}
}

func WalkRemoveFiles(path string) {
	dir, _ := os.ReadDir(path)
	for _, file := range dir {
		file_path := filepath.Join(path, file.Name())
		if file.IsDir() {
			WalkRemoveFiles(file_path)
		} else {
			if strings.HasSuffix(file_path, "txt") {
				cnt++
				fmt.Printf("%-4d: %s\n", cnt, file_path)
				// err := os.Remove(file_path)
				// if err != nil {
				// 	return
				// }
			}
		}
	}
}

func WalkRemoveFiles2(path string) {

	ContainsStr := func(arr []string, value string) bool {
		for _, v := range arr {
			if v == value {
				return true
			}
		}
		return false
	}

	print := func(file_path string) {
		cnt++
		fmt.Printf("%-4d: %s\n", cnt, file_path)
		// os.Remove(file_path)
	}
	dir, _ := os.ReadDir(path)
	for _, file := range dir {
		file_path := filepath.Join(path, file.Name())
		if file.IsDir() {
			WalkRemoveFiles2(file_path)
		} else {
			arr := strings.Split(file_path, ".")
			file_type := arr[len(arr)-1]
			switch {
			// 仅保留指定类型文件，删除类型为空
			case len(specifiedType) == 0 && len(retainType) != 0 && !ContainsStr(retainType, file_type):
				print(file_path)
			// 仅删除指定类型文件，保留类型为空
			case len(retainType) == 0 && len(specifiedType) != 0 && ContainsStr(specifiedType, file_type):
				print(file_path)
			// 保留指定类型文件的同时, 且删除指定类型文件
			case len(retainType) != 0 && len(specifiedType) != 0 && (!ContainsStr(retainType, file_type) && ContainsStr(specifiedType, file_type)):
				print(file_path)
			default:
				// print(file_path)
			}
		}
	}
}
