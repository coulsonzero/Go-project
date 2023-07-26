package hash

import "sort"

// sort
func minNumberDisappeared(nums []int) int {
	// write code here
	sort.Ints(nums)
	t := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == t {
			t++
		} else if nums[i] > t {
			return t
		}
	}
	return t
}

// hash
func minNumberDisappeared2(nums []int) int {
	// write code here
	m := make(map[int]bool)
	for _, v := range nums {
		if v > 0 {
			m[v] = true
		}
	}
	for k := 1; k < len(m); k++ {
		if !m[k] {
			return k
		}
	}
	return len(m) + 1
}

// []bool
func minNumberDisappeared3(nums []int) int {
	// write code here
	flag := make([]bool, len(nums))
	for _, v := range nums {
		if v > 0 && v <= len(nums) {
			flag[v-1] = true
		}
	}
	for i, v := range flag {
		if !v {
			return i + 1
		}
	}
	return len(nums) + 1
}
