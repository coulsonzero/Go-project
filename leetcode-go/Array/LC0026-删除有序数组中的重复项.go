package main

import "fmt"

/**
 * Question: LC26. 删除有序数组中的重复项
 *
 * Input   : nums = [0,0,1,1,1,2,2,3,3,4]
 * Output  : 5, nums = [0,1,2,3,4,_,_,_,_,_]
 * Explain : 不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
 *
 * @Date: 2022-06-14
 * @Author: CoulsonZero
 */

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针法
	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			l++
			nums[l] = nums[r]
		}
	}

	return l + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums)) // len: 5
}
