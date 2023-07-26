package main

import (
	"fmt"
	"reflect"
)

func Kind() {
	var num int
	var pi float64
	var name string
	var array [5]string
	var slice []string
	var m map[string]int
	var student struct{}
	var person interface{}

	fmt.Printf("%v\n", reflect.ValueOf(num).Kind())  // int: reflect.Int
	fmt.Printf("%v\n", reflect.ValueOf(pi).Kind())   // float64: reflect.Float64
	fmt.Printf("%v\n", reflect.ValueOf(name).Kind()) // string: reflect.String

	fmt.Printf("%v\n", reflect.ValueOf(slice).Kind())   // slice: reflect.Slice
	fmt.Printf("%v\n", reflect.ValueOf(array).Kind())   // array: reflect.Array
	fmt.Printf("%v\n", reflect.ValueOf(m).Kind())       // map: reflect.Map
	fmt.Printf("%v\n", reflect.ValueOf(student).Kind()) // struct: reflect.Struct

	fmt.Printf("%v\n", reflect.ValueOf(person).Kind()) // invalid: reflect.Interface

	fmt.Println(reflect.ValueOf(slice).Kind() == reflect.Slice)      // true
	fmt.Println(reflect.ValueOf(student).Kind() == reflect.Struct)   // true
	fmt.Println(reflect.ValueOf(person).Kind() == reflect.Interface) // false

}

func demo2() {
	slice_int := []int{1, 2, 3, 4, 5}
	// slice_string := []string{"a", "b", "c", "d", "a", "b", "c", "d"}
	// slice_float := []float64{1.11, 2.22, 3.33, 4.44, 1.11, 2.22, 3.33, 4.44}

	// fmt.Println(reflect.TypeOf(slice_int))                         // []int
	// fmt.Println(reflect.TypeOf(slice_int).String())                // []int
	// fmt.Println(reflect.TypeOf(slice_int).Kind())                  // slice
	// fmt.Println(reflect.TypeOf(slice_int).Kind() == reflect.Slice) // true

	fmt.Println(reflect.TypeOf(12).Name())        // int, 仅适用于基础类型
	fmt.Println(reflect.TypeOf(slice_int).Elem()) // int

	// fmt.Println(reflect.ValueOf(slice_int))                         // [1 2 3 4 5]
	// fmt.Println(reflect.ValueOf(slice_int).String())                // <[]int Value>
	// fmt.Println(reflect.ValueOf(slice_int).Kind())                  // slice
	// fmt.Println(reflect.ValueOf(slice_int).Kind() == reflect.Slice) // true

	fmt.Println(reflect.ValueOf(slice_int).Type())        // []int
	fmt.Println(reflect.ValueOf(slice_int).Len())         // 5
	fmt.Println(reflect.ValueOf(slice_int).Slice(1, 3))   // [2, 3]
	println(reflect.ValueOf(slice_int).CanSet())          // false
	println(reflect.ValueOf(slice_int).Index(0).CanSet()) // true
	reflect.ValueOf(slice_int).Index(0).SetInt(7)
	fmt.Println(reflect.ValueOf(slice_int)) // [7, 2, 3, 4, 5]
	// reflect.DeepEqual()

}

func demo() {
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
