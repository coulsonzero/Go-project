package main

func maximumSubarraySum2(nums []int, k int) int64 {
	ans := 0
	sum := 0
	n := len(nums)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		sum += nums[i]
		cnt[nums[i]]++
		if i >= k {
			sum -= nums[i-k]
			cnt[nums[i-k]]--
			if cnt[nums[i-k]] == 0 {
				delete(cnt, nums[i-k])
			}
		}
		if len(cnt) == k {
			ans = max(ans, sum)
		}
	}
	return int64(ans)
}

/*
	func maximumSubarraySum(nums []int, k int) int64 {
		var temp int
		for i := 0; i <= len(nums)-k; i++ {
			if !containsDuplicate(nums[i : i+k]) {
				temp = max(temp, sum(nums[i:i+k]))
			}
		}
		return int64(temp)
	}

	func containsDuplicate(nums []int) bool {
		m := make(map[int]int)
		for _, v := range nums {
			if _, ok := m[v]; ok {
				return true
			}
			m[v]++
		}
		return false
	}

	func sum(arr []int) int {
		var res int
		for _, v := range arr {
			res += v
		}
		return res
	}
*/
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func main() {
	nums := []int{1, 5, 4, 2, 9, 9, 9}
	k := 3
	println(maximumSubarraySum2(nums, k))                    // 15
	println(maximumSubarraySum2([]int{1, 1, 1, 7, 8, 9}, 3)) // 24
}
