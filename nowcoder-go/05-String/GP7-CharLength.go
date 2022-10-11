package main

/**
 * Question: GP7 字符数量
 * Descript: 给定一个字符串，统计其中的字符个数（一个中文算一个）
 * input : "小明的英文名叫jack"
 * output: 11
 *
 * Tip: 1. 汉字是采用unicode编码，占三个字节
 *      2. rune是int32的别名（-231~231-1），对比byte（-128～127），可表示的字符更多
 */

func count(s string) int {
	// write code here
	return len([]rune(s))
}
