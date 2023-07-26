package main

/**
 * 给你一个字符串 s 和一个字符 c ，且 c 是 s 中出现过的字符。
 * 返回一个整数数组 answer ，其中 answer.length == s.length 且 answer[i] 是 s 中从下标 i 到离它 最近 的字符 c 的 距离 。
 * 两个下标 i 和 j 之间的 距离 为 abs(i - j) ，其中 abs 是绝对值函数。
 *
 * 输入：s = "loveleetcode", c = "e"
 * 输出：[3,2,1,0,1,0,0,1,2,2,1,0]
 * 输入：s = "aaab", c = "b"
 * 输出：[3,2,1,0]
 */

func shortestToChar(s string, c byte) []int {
	var temp []int
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			temp = append(temp, i)
		}
	}
	// fmt.Printf("%v \n", temp)

	var res []int
	var index int
	var distance int
	for i := 0; i < len(s) && index < len(temp); i++ {
		// fmt.Printf("index: %d, i: %d, v: %c, dis:%d\n", temp[index], i, s[i], temp[index]-i)
		if i != temp[index] {
			// 寻找最近的c比较距离
			if index == 0 {
				distance = abs(temp[index] - i)
			} else {
				distance = min(abs(temp[index]-i), abs(temp[index-1]-i))
			}
			res = append(res, distance)

		} else {
			res = append(res, 0)
			index++
			// 防索引溢出
			if index == len(temp) {
				index--
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
