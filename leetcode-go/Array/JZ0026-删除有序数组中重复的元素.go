package main

import "fmt"

func main() {
	nums := []int {1, 6, 6, 8}
	res := remove_duplicates(nums)
	fmt.Println(res)
	// [1, 6, 8]
}


func remove_duplicates(nums []int) []int {
	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
	}
	return nums[:l+1]
}