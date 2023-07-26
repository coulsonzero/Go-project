package 二分查找_排序

import "fmt"

/**
 * 二分查找
 * 输入：[-1,0,3,4,6,10,13,14],13
 * 输出：6
 */

func search(nums []int, target int) int {
	// write code here
	l, r := 0, len(nums)-1
	for l <= r {
		// m := (l + r) / 2  	// not best
		m := l + (r-l)/2
		// m := l + (r - l) >> 1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

func notice() {
	l := -4
	r := -1
	// 建议使用后者: mid = l + (r-l) / 2
	// 当L和R为负数时: 编译器计算整数除法时是按照绝对值向下取整的规则计算的
	fmt.Printf("m = (L+R)/2: %v \n", (l+r)/2)         // -2
	fmt.Printf("m = L + (R-L)/2: %v \n", l+(r-l)/2)   // -3
	fmt.Printf("m = L + (R-L)>>1: %v \n", l+(r-l)>>1) // -3
}
