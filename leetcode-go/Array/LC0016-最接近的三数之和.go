package main

import (
	"math"
	"sort"
)

/**
 * 请你从 nums 中选出三个整数，使它们的和与 target 最接近。假定每组输入只存在恰好一个解。
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 * 3 <= nums.length <= 1000
 */

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans, n := nums[0]+nums[1]+nums[2], len(nums)
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			} else {
				if math.Abs(float64(target-sum)) < math.Abs(float64(target-ans)) {
					ans = sum
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	println(threeSumClosest(nums, target))
}
