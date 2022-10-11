package main

import "fmt"

/**
 * Q: 2295. 替换数组中的元素
 *
 * input : nums = [1,2,4,6], operations = [[1,3],[4,7],[6,1]]
 * output: [3,2,7,1]
 * Explanation: We perform the following operations on nums:
 * - Replace the number 1 with 3. nums becomes [3,2,4,6].
 * - Replace the number 4 with 7. nums becomes [3,2,7,6].
 * - Replace the number 6 with 1. nums becomes [3,2,7,1].
 * We return the final array [3,2,7,1].
 */

func arrayChange(nums []int, operations [][]int) []int {
	m := make(map[int]int, len(nums))

	for i := range nums {
		m[nums[i]] = i
	}
	for i := range operations {
		a, b := operations[i][0], operations[i][1]
		m[b] = m[a]
		delete(m, a)
	}
	fmt.Println(m)
	for k, v := range m {
		nums[v] = k
	}
	return nums
}

func main() {
	nums := []int{1, 2, 4, 6}
	operations := [][]int{{1, 3}, {4, 7}, {6, 1}}

	fmt.Println(arrayChange(nums, operations)) // Output: [3 2 7 1]
}
