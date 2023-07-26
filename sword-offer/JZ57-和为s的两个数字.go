package main

/**
 * 注意：此题为升序数组，故可用二分查找；两数之和的数组为随机顺序，不可用此二分查找法，比较繁琐
 */

// hashmap 哈希表法
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if j, ok := m[target-nums[i]]; ok {
			return []int{nums[i], nums[j]}
		}
		m[nums[i]] = i
	}

	return nil
}

// 二分查找：此题为升序数组
func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		sum := nums[l] + nums[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			return []int{nums[l], nums[r]}
		}
	}

	return nil
}
