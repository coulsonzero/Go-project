package main

/**
 * 在排序数组中查找数字 I
 * 难度：简单
 * 描述：统计一个数字在排序数组中出现的次数
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: 2
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: 0
 * 输入: nums = [1], target = 1
 * 输出: 1
 * 输入: nums = [1], target = 0
 * 输出: 0
 */

// 数组计数
func search(nums []int, target int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}

	return count
}

// 双指针
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		for l <= r && nums[l] < target {
			l++
		}
		for l <= r && nums[r] > target {
			r--
		}

		return r - l + 1
	}

	return 0
}

/**
// 二分查找，时间复杂度：O(logN)
class Solution {
    public int search(int[] nums, int target) {
        if (nums.length == 0) return 0;
        return binarySearch(nums, target + 1) - binarySearch(nums, target);
    }

    private int binarySearch(int[] nums, int target) {
        int left = 0, right = nums.length;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        return right;
    }
}
*/
