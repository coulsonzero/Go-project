package main

import "strconv"

/**
 * 输入：s = "dfa12321afd"
 * 输出：2
 * 解释：出现在 s 中的数字包括 [1, 2, 3] 。第二大的数字是 2 。
 * 输入：s = "abc1111"
 * 输出：-1
 * 解释：出现在 s 中的数字
 * 包含 [1] 。没有第二大的数字。
 */

func secondHighest(s string) int {
	max1, max2 := -1, -1
	for _, c := range s {
		if '0' <= c && c <= '9' {
			num := int(c - '0')
			if num > max1 {
				max1, max2 = num, max1
			} else if max2 < num && num < max1 {
				max2 = num
			}
		}
	}
	return max2
}

func secondHighest2(s string) int {
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			m[num]++
		}
	}

	// fmt.Println(m)

	var res []int
	for i := 9; i >= 0; i-- {
		if m[i] != 0 {
			res = append(res, i)
		}
	}

	if len(res) <= 1 {
		return -1
	}
	return res[1]
}
