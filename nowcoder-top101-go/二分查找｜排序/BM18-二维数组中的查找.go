package 二分查找_排序

/**
 * 题目：二维数组中的查找
 * 描述：在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
		请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 * 输入：7,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
 * 输出：true
 * 说明：存在7，返回true
 * 输入：3,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
 * 输出：false
 * 说明：不存在3，返回false
*/

// Find 时间复杂度O(N)
func Find(target int, array [][]int) bool {
	// write code here
	j := len(array[0]) - 1 // 每个一维数组的长度相同
	for i := 0; i < len(array) && j >= 0; {
		if array[i][j] == target {
			return true
		} else if array[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

// Find2 时间复杂度 O(N^2)
func Find2(target int, array [][]int) bool {
	// write code here
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if array[i][j] == target {
				return true
			} else if array[i][j] > target {
				break
			}
		}
	}
	return false
}
