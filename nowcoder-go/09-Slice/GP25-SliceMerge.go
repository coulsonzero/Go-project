package main

import "sort"

/**
 * question: GP25 合并有序数组
 * description: 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
 *              请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列
 * input : [1,2,3,0,0,0],3,[2,5,6],3
 * output: [1,2,2,3,5,6]
 */
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// write code here
	for i, _ := range nums2 {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
	return nums1
}
