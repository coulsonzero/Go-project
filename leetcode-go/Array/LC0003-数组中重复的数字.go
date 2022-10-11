package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums)) // 2
}

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if m[v] != 0 {
			return v
		}
		m[v]++
	}
	return -1
}
