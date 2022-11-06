package main

import "fmt"

// 空间复杂度: O(1)
func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	// 移动零
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}

	return nums
}

// 空间复杂度: O(n)
func applyOperations2(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	var res []int
	for _, v := range nums {
		if v != 0 {
			res = append(res, v)
		}
	}

	temp := make([]int, n)
	copy(temp, res)
	return temp
}

func main() {
	nums := []int{1, 2, 2, 1, 1, 0}
	fmt.Printf("%v", applyOperations(nums))
}
