package main

import "sort"

/**
 * TODO 多数元素
 * 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素
 * TODO 输入：nums = [2,2,1,1,1,2,2]
 * TODO 输出：2
 */

func majorityElement(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}

// hashmap 哈希表计数
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	n := len(nums)
	for _, v := range nums {
		m[v]++
		if m[v] > n/2 {
			return v
		}
	}
	return -1
}
