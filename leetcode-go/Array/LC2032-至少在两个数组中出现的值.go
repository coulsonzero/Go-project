package main

import "sort"

/**
 * 题目：2032. 至少在两个数组中出现的值
 * 难度：简单
 * 描述：给你三个整数数组nums1、nums2和nums3，请你构造并返回一个元素各不相同的数组，且由至少在两个数组中出现的所有值组成。数组中的元素可以按任意顺序排列
 *
 * 输入：nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
 * 输出：[3,2]
 * 输入：nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
 * 输出：[2,3,1]
 */

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m := make(map[int]int)
	addToMap(nums1, m)
	addToMap(nums2, m)
	addToMap(nums3, m)
	// fmt.Printf("%v\n", m)

	ans := make([]int, 0)
	for k, v := range m {
		if v >= 2 {
			ans = append(ans, k)
		}
	}

	return ans
}

func addToMap(nums []int, m map[int]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		m[nums[i]]++
	}
	m[nums[len(nums)-1]]++
}
