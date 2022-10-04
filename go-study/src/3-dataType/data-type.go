package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

func main() {
	// dataType()
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(Leng(nums))

	var num int = 7
	var pi float32 = 3.141
	var c complex64 = 1.10 + 10i
	fmt.Println(num, pi, c)
	fmt.Println(unsafe.Sizeof(num))  // 8
	fmt.Println(unsafe.Sizeof(pi))   // 4
	fmt.Println(unsafe.Sizeof(c))    // 8
	fmt.Println(reflect.TypeOf(num)) // int
	fmt.Println(reflect.TypeOf(pi))  // float32
	fmt.Println(reflect.TypeOf(c))   // complex64
	fmt.Printf("int: -%d ~ %d\n", math.MaxInt, math.MaxInt)

	var r rune = '你'
	fmt.Println(string(r))
	var b byte = 'v'
	fmt.Println(string(b))

}

func dataType() {
	var (
		name     string
		age      int
		gender   bool
		score    float64
		skills   []string
		scoreMap map[string]int
	)
	fmt.Printf("name =  %s, age = %d, gender = %t, score = %f \n", name, age, gender, score)
	// name = "", age = 0, gender = false, score = 0.000000
	fmt.Printf("%T, %T, %T, %T \n", name, age, gender, score)
	// string, int, bool, float64
	fmt.Printf("%v %T \n", skills, skills)
	// [] []string
	fmt.Printf("%v, %T \n", scoreMap, scoreMap)
	// map[], map[string]int
}

func Leng(arr interface{}) (length int) {
	if arr == nil {
		return
	}
	switch arr.(type) {
	case []int:
		length = len(arr.([]int))
	case []string:
		length = len(arr.([]string))
	case []float32:
		length = len(arr.([]float32))
	default:
		return
	}
	return
}
