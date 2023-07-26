package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var ans [][]int
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			val := target - nums[i] - nums[j]
			l := j + 1
			r := n - 1
			for l < r {
				cur := nums[l] + nums[r]
				if cur == val {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if cur > val {
					r--
				} else {
					l++
				}
			}
		}
	}

	return ans
}
