package main

/**
 * question: GP24 判断两个切片是否有相同的元素
 * description: 给定两个切片，判断这两个切片中的元素是否完全一样。
 * input : [1,2,3,4],[1,2,3,4]
 * output: true
 */
func equal(s1 []int, s2 []int) bool {
	// write code here
	if len(s1) != len(s2) {
		return false
	}
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
