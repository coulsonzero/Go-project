package main

/**
 * 解题思路：
 * 遍历字符串s，如果是'D'就将最后一个数字插入到当前索引位置并删除该末尾数字，如果是'I'就继续
 * 返回的数组长度比字符串s长度多一个，排序后范围为[0, len(s)]
 */

// 执行时间：0ms 内存消耗：4.9MB
func diStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	l, r := 0, n
	for i, c := range s {
		if c == 'I' {
			res[i] = l
			l++
		} else {
			res[i] = r
			r--
		}
	}
	// 数组长度比字符串s长度多一个，此时 l=r
	res[n] = l
	return res
}

// 执行时间：0ms 内存消耗：6.1MB
func diStringMatch2(s string) []int {
	var res []int
	n := len(s)
	l, r := 0, n
	for _, c := range s {
		if c == 'I' {
			res = append(res, l)
			l++
		} else {
			res = append(res, r)
			r--
		}
	}
	// 数组长度比字符串s长度多一个，此时 l=r
	res = append(res, l)
	return res
}

// 数组原地操作(插入、删除)：104ms
// func diStringMatch3(s string) []int {
// 	nums := make([]int, 0)
// 	for i := 0; i < len(s)+1; i++ {
// 		nums = append(nums, i)
// 	}
// 	fmt.Printf("%v \n", nums)
//
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == 'D' {
// 			nums = insert(nums, i, nums[len(nums)-1])
// 			nums = delete(nums, len(nums)-1)
// 			fmt.Printf("%v", nums)
// 		}
// 	}
//
// 	// fmt.Printf("%v", nums)
// 	return nums
// }
//
// func insert(slice []int, index int, value int) []int {
// 	return append(slice[:index], append([]int{value}, slice[index:]...)...)
// }
//
// func delete(slice []int, i int) []int {
// 	return append(slice[:i], slice[i+1:]...)
// }
