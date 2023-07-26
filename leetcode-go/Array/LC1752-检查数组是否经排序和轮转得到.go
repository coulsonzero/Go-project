package main

/**
 * 给你一个数组 nums 。nums 的源数组中，所有元素与 nums 相同，但按非递减顺序排列。
 * 如果 nums 能够由源数组轮转若干位置（包括 0 个位置）得到，则返回 true ；否则，返回 false 。
 * 源数组中可能存在 重复项 。
 * 注意：我们称数组 A 在轮转 x 个位置后得到长度相同的数组 B ，当它们满足 A[i] == B[(i+x) % A.length] ，其中 % 为取余运算
 */

// 个非递减序列无论怎么轮转，整个循环数组里最多只会出现一次逆序（num[i] > nums[(i + 1) % n），即发生在两个有序数列的交界处
// 那么只需要统计循环数组中出现的逆序的次数，如果这个次数<=1，那么就是符合题意的
func check(nums []int) bool {
	n := len(nums)
	cnt := 0
	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			cnt++
		}
	}
	return cnt <= 1
}

// 出现两次以上递减直接返回false；
// 不出现递减，返回true；
// 出现一次递减，就看nums的最后一个元素是否不大于第一个元素

// 在遍历的过程中，元素只能下降一次，否则返回false
func check2(nums []int) bool {
	f := false
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			if f {
				return false
			}
			f = true
		}
	}
	return true
}
