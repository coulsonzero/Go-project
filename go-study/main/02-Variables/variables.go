package main

import "fmt"

// 变量声明赋值后仍需调用，否则报错

// 全局变量(函数外声明的变量)
var (
	id       int
	username string
	password string
)

var (
	Host = "127.0.0.1"
	Port = 3306
)

func main() {
	// 局部变量(函数内声明的变量)
	name, age := "Tom", 20

	fmt.Println(name, age)
	fmt.Printf("result: %d, %s, %s", id, username, password)
	fmt.Println(Host, Port)

}

// single variable
func createSingleVar() {
	var num int
	fmt.Println(num)

	var str = "hello"
	fmt.Println(str)

	name := "coulson"
	fmt.Println(name)
}

// multiple variable
func createMultipleVar() {
	var name, age = "john", 20
	fmt.Println(name, age)

	num, str := 12, "hello"
	fmt.Println(num, str)

	var (
		host = "127.0.0.1"
		port = "3306"
	)
	fmt.Println(host, port)
}
