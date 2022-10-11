package main

/**
 * question: GP46 体温异常
 * description: 实现一个函数，该函数的功能是 给定一个float类型变量表示某个人的体温，如果有人体温大于37.5，抛出"体温异常"，并输出。
 * input : 38.000000
 * output: "体温异常"
 */

func temperature(t float64) (res string) {
	// write code here
	defer func() {
		if err := recover(); err != nil {
			res = "体温异常"
		}
	}()

	if t > 37.5 {
		panic("体温异常")
	}
	return res
}
