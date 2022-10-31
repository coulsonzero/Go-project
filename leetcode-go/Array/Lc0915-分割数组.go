package main

func main() {
	arr := []int{5, 0, 3, 8, 6}
	println(partitionDisjoint(arr)) // 3

	arr2 := []int{1, 1, 1, 0, 6, 12}
	println(partitionDisjoint(arr2)) // 4

}

// 如果遇到比左侧最大值小的元素，说明左侧应该包含该值，更新下pos的位置为当前遍历位置，并将左侧最大值更新为最大值
func partitionDisjoint(nums []int) int {
	lMax, curMax, pos := nums[0], nums[0], 0
	for i := 1; i < len(nums); i++ {
		curMax = max(curMax, nums[i])
		if nums[i] < lMax {
			pos = i
			lMax = curMax
		}
	}
	return pos + 1
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

/*
 * 执行结果：通过
 * 执行用时：120 ms, 在所有 Go 提交中击败了50.75% 的用户
 * 内存消耗：8.1 MB, 在所有 Go 提交中击败了96.66% 的用户
 */

// 以下超出时间限制

// func partitionDisjoint2(nums []int) int {
// 	for i := 1; i < len(nums); i++ {
// 		if maxNum(nums[:i]) <= minNum(nums[i:]) {
// 			return i
// 		}
// 	}
// 	return -1
// }
//
// func maxNum(nums []int) int {
// 	res := nums[0]
// 	for _, v := range nums {
// 		if v >= res {
// 			res = v
// 		}
// 	}
// 	return res
// }
//
// func minNum(nums []int) int {
// 	res := nums[0]
// 	for _, v := range nums {
// 		if v <= res {
// 			res = v
// 		}
// 	}
// 	return res
// }
