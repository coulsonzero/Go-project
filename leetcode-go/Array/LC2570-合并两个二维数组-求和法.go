package main

import "sort"

// 额外空间：O(N)
// 使用map哈希表存储数据，排序keys后输出而为数组
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	// map to plus
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v[0]] = v[1]
	}
	for _, v := range nums2 {
		m[v[0]] += v[1]
	}

	// sort keys
	keys := []int{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// ans
	ans := [][]int{}
	for _, v := range keys {
		ans = append(ans, []int{v, m[v]})
	}
	return ans
}

// 额外空间：O(1)
// 由于两个数组均为排序数组，所以可以通过比较大小来添加到新的而为数组中，无需排序
func mergeArrays_2(nums1 [][]int, nums2 [][]int) [][]int {
	var ans [][]int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] < nums2[j][0] {
			ans = append(ans, nums1[i])
			i++
		} else if nums1[i][0] > nums2[j][0] {
			ans = append(ans, nums2[j])
			j++
		} else {
			ans = append(ans, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		}
	}

	ans = append(ans, nums1[i:]...)
	ans = append(ans, nums2[j:]...)
	// fmt.Printf("ans: %v \n", ans)
	// fmt.Printf("nums1: %v; nums2: %v", nums1[i:], nums2[j:])

	return ans
}
