package main

import "sort"

/**
 * 349. 两个数组的交集
 * 难度：简单
 * 描述：给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序。
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2]
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[9,4]
 * 解释：[4,9] 也是可通过的
 */

// hashmap
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = true
	}

	m2 := make(map[int]bool)
	for _, v := range nums2 {
		if m[v] {
			m2[v] = true
		}
	}
	// 除重
	var ans []int
	for k := range m2 {
		ans = append(ans, k)
	}

	return ans
}

// sort
func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	l := 0
	r := 0
	var ans []int
	for l < len(nums1) && r < len(nums2) {
		if nums1[l] < nums2[r] {
			l++
		} else if nums1[l] > nums2[r] {
			r++
		} else {
			// 除重
			if len(ans) == 0 || nums1[l] > ans[len(ans)-1] {
				ans = append(ans, nums1[l])
			}
			l++
			r++
		}
	}

	return ans
}
