package main

/**
 * 27. 移除元素
 * 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
 * 要求：必须仅使用 O(1) 额外空间并 原地 修改输入数组。
 * 输入：nums = [3,2,2,3], val = 3
 * 输出：2, nums = [2,2]
 */

func removeElement(nums []int, val int) int {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != val {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
	return l
}


func removeElement2(nums []int, val int) int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == val {
			r--
			nums[l] = nums[r]
		} else {
			l++
		}
	}
	return l
}

func removeElement3(nums []int, val int) int {
	l := 0
	for _, v := range nums {
		if v != val {
			nums[l] = v
			l++
		}
	}
	return l
}