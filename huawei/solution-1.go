package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		arr = append(arr, x)
	}

	// arr := []int{1, 2, 3, 1, 4}
	// arr := []int{4, 3, 1, 2, 3, 1, 4, 2, 1}
	fmt.Println(getMaxLength(arr))

}

func getMaxLength(arr []int) int {
	var length int
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; i < j; {
			if arr[j] == arr[i] {
				length = max(j-i, length)
			}
			j--
		}
	}
	if length == 0 {
		return -1
	} else {
		return length
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
