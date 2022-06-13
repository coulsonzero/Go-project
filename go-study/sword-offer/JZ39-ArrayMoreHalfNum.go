package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(MoreThanHalfNum_Solution2(nums))
}

func MoreThanHalfNum_Solution2(numbers []int) int {
	cntMap := make(map[int]int, len(numbers))
	for _, v := range numbers {
		cntMap[v]++
		if cntMap[v] > len(numbers)/2 {
			return v
		}
	}
	return -1
}

func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	cntMap := make(map[int]int, len(numbers))
	for _, v := range numbers {
		cntMap[v]++
	}

	for i, v := range cntMap {
		if v > len(numbers)/2 {
			return i
		}
	}
	return -1
}
