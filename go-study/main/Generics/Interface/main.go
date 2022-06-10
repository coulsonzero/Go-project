package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	ArrayToString([]int{1, 2, 3, 4})
}

// 数组转字符串 (泛型出来前准备)
func ArrayToString(arr interface{}) {
	// 反射获取数据
	array := reflect.ValueOf(arr)
	for i := 0; i < array.Len(); i++ {
		fmt.Println(array.Index(i))
	}
}
