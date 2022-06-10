package main

import "fmt"

/**
 * 长度自定义
 * 可增删改
 */

func main() {

}

func test() {
	list := make([]int, 3) // 初始默认值为0
	fmt.Println(list)      // [0, 0, 0]
	list = append(list, 1, 3)
	fmt.Println(list) // [0 0 0 1 3]

	iter(list...)

	arr := []int{1, 3, 5, 4, 2}
	fmt.Println(arr)
	list = append(arr, 8)
	fmt.Println(list) // [1 3 5 4 2 8]

	sum(list...) // 23
}

func iter(nums ...int) {
	for i, v := range nums {
		fmt.Println(i, v)
	}
}

func sum(nums ...int) {
	res := 0
	for _, v := range nums {
		res += v
	}
	fmt.Println(res)
}
