package main

import "sort"

/**
 * LC 只出现一次的数字
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 * 输入: [2,2,1]
 * 输出: 1
 * 输入: [4,1,2,1,2]
 * 输出: 4
 */

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func singleNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums)-1; i += 2 {
		if nums[i] == nums[i+1] {
			return nums[i-1]
		}
	}
	return nums[len(nums)-1]
}

func main() {
	arr := []int{4, 1, 2, 1, 2, 3, 4}
	println(singleNumber(arr))
}
