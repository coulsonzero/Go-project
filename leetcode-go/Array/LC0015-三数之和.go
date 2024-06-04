package main

import (
	"fmt"
	"reflect"
	"sort"
)

/**
 * 问题：15. 三数之和
 * 描述：返回所有和为 0 且不重复的三元组。
 * 注意：答案中不可以包含重复的三元组。
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 * 输入：nums = [0,0,0,0]
 * 输出：[[0,0,0]]
 */

// 时间：752 m, 击败5.2%; 内存：9.2 MB, 击败5.5%.
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		// 排除重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if temp := twoSum(nums, -nums[i], i); temp != nil {
			res = append(res, temp...)
		}
	}

	// 除重
	if len(res) == 0 {
		return nil
	}
	arr := [][]int{res[0]}
	for i := 1; i < len(res); i++ {
		if reflect.DeepEqual(res[i], res[i-1]) {
			continue
		}
		arr = append(arr, res[i])
	}
	return arr
}

func threeSum(nums []int, target int, index int) [][]int {
	var ret [][]int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 排除target
		if i == index {
			continue
		}
		// 多个符合条件
		if j, ok := m[target-nums[i]]; ok && nums[index] <= nums[j] {
			ret = append(ret, []int{nums[index], nums[j], nums[i]})
		}
		m[nums[i]] = i
	}
	return ret
}

// 时间：40 ms, 击败20.6%; 内存：7.3 MB, 击败84.75%.
// 时间复杂度：O(N^2), 空间复杂度：O(logN)
func threeSum2(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	sort.Ints(nums)

	for first := 0; first < n; first++ {
		// 排除重复元素
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third, target := n-1, -nums[first]
		for second := first + 1; second < n; second++ {
			// 排除重复
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

// 时间：28 ms, 击败88.29%; 内存：7.3 MB, 击败84.75%.
func threeSum(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	sort.Ints(nums)

	for i, _ := range nums {
		// 除重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r, target := i+1, n-1, -nums[i]
		for l < r {
			if nums[l]+nums[r] == target {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				// 数组除重
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				for r > l && nums[l] == nums[l+1] {
					l--
				}
				r--
				l++
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}

func main() {
	fmt.Printf("res: %v\n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("res: %v\n", threeSum([]int{0, 1, 1}))
	fmt.Printf("res: %v\n", threeSum([]int{0, 0, 0}))
	fmt.Printf("res: %v\n", threeSum([]int{0, 0, 0, 0}))
	fmt.Printf("res: %v\n", threeSum([]int{-2, 0, 0, 2, 2}))
	fmt.Printf("res: %v\n", threeSum([]int{3, 0, -2, -1, 1, 2}))
	fmt.Printf("res: %v\n", threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
	fmt.Printf("res: %v\n", threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))

}
