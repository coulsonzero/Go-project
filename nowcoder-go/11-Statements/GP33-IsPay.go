package main

/**
 * question: GP33 游乐园门票
 * description: 有个游乐园，如果身高小于160.0，则可以免费入场，身高大于等于160.0岁则要收费，根据输入的身高判断是否需要收费，收费返回true,不收费返回false。
 * input : 156.000000
 * output: false
 */

func ispay(hight float64) bool {
	// write code here
	return hight >= 160
}
