package main

import "fmt"

/**
 * 长度固定
 * 无法增删改
 */

func main() {
	arr := [5]int{1, 3, 5, 4, 2}

	fmt.Println(arr)      // [1, 3, 5, 4, 2]
	fmt.Println(len(arr)) // 5
	fmt.Println(arr[0])   // 1
	fmt.Println(arr[1:3]) // [3, 5]
	for i := range arr {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	// fmt.Println(sum(arr...)) // 15
}


func sum(nums ...int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}



