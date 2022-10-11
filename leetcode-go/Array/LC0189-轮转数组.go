package main

import "fmt"

/**
 * LC 189. 轮转数组
 * 原地修改
 * Input: nums = [1,2,3,4,5,6,7], k = 3
 * Output: [5,6,7,1,2,3,4]
 *
 * Input: nums = [-1,-100,3,99], k = 2
 * Output: [3,99,-1,-100]
 */

// 空间复杂度O(n)
func rotate(nums []int, k int) {
	k = k % len(nums)
	m := len(nums) - k
	ret := make([]int, len(nums[:m]))
	copy(ret, nums[:m])
	copy(nums, nums[m:])
	copy(nums[k:], ret)
}

// 空间复杂度O(n)
func rotate2(nums []int, k int) {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[(i+k)%len(nums)] = v
	}
	copy(nums, res)
}

// 空间复杂度O(1): 推荐
func rotate3(nums []int, k int) {
	reverse := func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}

	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)

	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	rotate(nums2, k2)
	fmt.Println(nums2)

	nums3 := []int{-1}
	k3 := 2
	rotate(nums3, k3)
	fmt.Println(nums3)
}

/*
// python3
class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        k %= len(nums)
        nums[:] = nums[-k:]+nums[:-k]
*/
