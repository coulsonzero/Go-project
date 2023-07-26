package hash

import "sort"

/**
 * BM52 数组中只出现一次的两个数字
 * 描述：一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字
 * 要求：时间复杂度O(n), 空间复杂度O(1)
 * 提示：输出时按非降序排列
 * 输入：[1,4,1,6]
 * 输出：[4,6]
 * 输入：[1,2,3,3,2,9]
 * 输出：[1,9]
 */

func FindNumsAppearOnce(nums []int) []int {
	// write code here
	var res []int
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	// 输出时按非降序(升序)排列 ！
	sort.Ints(res)
	return res
}
