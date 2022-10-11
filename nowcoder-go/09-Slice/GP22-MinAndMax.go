package main

import "sort"

/**
 * question: GP22 评委打分
 * description: 小明参加某个歌唱比赛，评委们进行打分,要求去掉最高分，和最低分，将最高分和最低分依次存入切片并返回。
 * input : [1,2,3,4,5,6,7,8,9]
 * output: [1,9]
 */
func minAndMax(s []int) []int {
	// write code here
	sort.Ints(s)
	return []int{s[0], s[len(s)-1]}
}
