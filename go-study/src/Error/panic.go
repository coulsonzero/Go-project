package main

import (
	"fmt"
	"strconv"
)

// panic: 返回相应的错误信息，终止程序运行

func main() {
	num, err := strconv.Atoi("a123")
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}
