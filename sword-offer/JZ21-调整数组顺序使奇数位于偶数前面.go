package main

import "sort"

/**
 * 输入：nums = [1,2,3,4]
 * 输出：[1,3,2,4]
 * 注：[3,1,2,4] 也是正确的答案之一。
 */

// 空间复杂度：O(N)
func exchange(nums []int) []int {
	var arr_q, arr_o []int
	for _, v := range nums {
		if v%2 == 1 {
			arr_q = append(arr_q, v)
		} else {
			arr_o = append(arr_o, v)
		}
	}

	arr_q = append(arr_q, arr_o...)

	return arr_q
}

// 排序：空间复杂度O(1)
func exchange2(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]%2 > nums[j]%2
	})

	return nums
}

// 快速排序: 空间复杂度O(1)
func exchange3(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[r]%2 == 0 {
			r--
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
	return nums
}

func exchange4(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		for l < r && nums[l]%2 == 1 {
			l++
		}
		for l < r && nums[r]%2 == 0 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return nums
}

func exchange5(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[r]%2 == 0 {
			r--
		} else if nums[l]%2 == 1 {
			l++
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return nums
}
