package main

/**
 * question: GP29 字符串构成
 * description: 给定两个字符串des 和src  ，判断 des能不能由 src 里面的字符构成,//如果可以，返回 true ；否则返回 false,
 *              src中的每个字符只能在 des 中使用一次
 * input : "ab","aab"
 * output: true
 */
func canConstruct(des string, src string) bool {
	// write code here
	m := make(map[rune]int)
	for _, k := range src {
		m[k]++
	}

	for _, v := range des {
		if m[v] != 0 {
			m[v]--
		} else {
			return false
		}
	}

	return true
}
