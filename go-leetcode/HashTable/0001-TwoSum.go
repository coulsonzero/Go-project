package Array

import "fmt"

/******************************************************************
 * Input: nums = [2,7,11,15], target = 9                          *
 * Output: [0,1]                                                  *
 * Explanation: Because nums[0] + nums[1] == 9, we return [0, 1]. *
 ******************************************************************/

// HashMap
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, x := range nums {
		if j, ok := dict[target-x]; ok {
			return []int{j, i}
		}
		dict[x] = i
	}
	return nil
}

// 暴力解法
func twoSum2(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums, target := []int{2, 7, 11, 15}, 9
	res := twoSum(nums, target)
	fmt.Println(res) // [0, 1]
}
