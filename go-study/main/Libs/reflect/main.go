package main

import (
	"fmt"
	"reflect"
)

func main() {
	// x := 12.6
	// fmt.Println("type:", reflect.TypeOf(x))
	// fmt.Println("value:", reflect.ValueOf(x))
	//
	// slice := []int{1, 2, 3}
	// fmt.Println("type:", reflect.TypeOf(slice))   // []int
	// fmt.Println("value:", reflect.ValueOf(slice)) // [1, 2, 3]
	// fmt.Println(reflect.ValueOf(slice).Kind())    // slice

	// ArrayToString([]int{1, 2, 3, 4})

}

func ArrayToString(arr interface{}) {
	// 反射获取数据
	array := reflect.ValueOf(arr)
	for i := 0; i < array.Len(); i++ {
		fmt.Println(array.Index(i))
	}
}

func Contains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
