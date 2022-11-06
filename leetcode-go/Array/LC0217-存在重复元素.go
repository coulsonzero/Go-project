package main

import (
	"fmt"
	"sort"
)

/**
 *
 * 输入：nums = [1,2,3,1]
 * 输出：true
 * 输入：nums = [1,2,3,4]
 * 输出：false
 */

// 哈希表法
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v]++
	}
	return false
}

// 排序法, 前后比较
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums)) // true

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2)) // false

	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums3)) // true
}
