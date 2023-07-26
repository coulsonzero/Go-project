package main

func findDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		temp := nums[0]
		if nums[0] == nums[temp] {
			return nums[0]
		}

		nums[0], nums[temp] = nums[temp], nums[0]
	}

	return -1
}
