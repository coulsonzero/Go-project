package main

import "fmt"

// 常量可以定义后暂不调用
const str string = "constant"

const (
	pi   float64 = 3.14
	host string  = "127.0.0.1"
	port int     = 3306
)

const (
	a = iota * 2 // 0
	b            // 2
	c            // 4
)

func main() {
	fmt.Println(pi, host, port, str)
	fmt.Println(a, b, c)
}
