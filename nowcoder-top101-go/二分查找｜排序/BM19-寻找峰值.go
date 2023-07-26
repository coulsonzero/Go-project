package 二分查找_排序

/**
 * 给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可。
 */

// 普通解法 O(N)
func findPeakElement(nums []int) int {
	// write code here
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}

	if nums[0] < nums[len(nums)-1] {
		return len(nums) - 1
	}
	return 0
}

// 二分法 O(logN)
func findPeakElement2(nums []int) int {
	// write code here
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	// 会形成两个山峰，一个是索引为1，峰值为4的山峰，另一个是索引为5，峰值为8的山峰
	println(findPeakElement([]int{2, 4, 1, 2, 7, 8, 4}))
	println(findPeakElement([]int{1, 2, 3, 1}))
	// notice
	println(findPeakElement([]int{2}))
	println(findPeakElement([]int{3, 2}))
	println(findPeakElement([]int{1, 2, 3, 4}))
	println(findPeakElement([]int{4, 3, 2, 1}))
}
