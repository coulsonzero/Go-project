package main

import "fmt"

func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	arr := make([]int, len(nums)/2)
	for i := 0; i < len(nums); i += 2 {
		if i&1 == 0 {
			arr[i] = min(nums[i], nums[i+1])
		} else {
			arr[i] = max(nums[i], nums[i+1])
		}
	}

	nums = arr
	return nums[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	nums := []int{1, 3, 5, 2, 4, 8, 2, 2}
	fmt.Println(minMaxGame(nums)) // Output: 1

	nums2 := []int{70, 38, 21, 22}
	fmt.Println(minMaxGame(nums2)) // Output: 22
}
