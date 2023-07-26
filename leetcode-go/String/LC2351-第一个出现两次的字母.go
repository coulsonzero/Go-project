package main

/**
 * 题目：2351. 第一个出现两次的字母
 * 难度：简单
 * 描述：给你一个由小写英文字母组成的字符串 s ，请你找出并返回第一个出现 两次 的字母。
 * 输入：s = "abccbaacz"
 * 输出："c"
 * 输入：s = "abcdd"
 * 输出："d"
 */

func repeatedCharacter(s string) byte {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]] == 2 {
			return s[i]
		}
	}

	return 0
}
