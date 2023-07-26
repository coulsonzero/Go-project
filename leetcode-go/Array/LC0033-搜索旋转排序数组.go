package main

/*
 * question: 33. 搜索旋转排序数组
 * descript: 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
 * 要求：你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 */

func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
