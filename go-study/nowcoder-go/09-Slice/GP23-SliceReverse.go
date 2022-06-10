package main

/**
 * question: GP23 调整顺序
 * description: [1,3,2,3,4,6]，重新排列后为[6,4,3,2,3,1]
 * input : [1,2,3,4,5,6,7,8,9]
 * output: [9,8,7,6,5,4,3,2,1]
 */
func convert(s []int) []int {
	// write code here
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	return s
}
