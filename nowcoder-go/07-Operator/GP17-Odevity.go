package main

/**
 * question: GP17 联谊活动
 * description: 某公司举办了一个联谊活动，现在要统计参加活动人数的单双，如果是单数，返回false，偶数返回true
 * input : 1
 * output: false
 */
func odevity(x int) bool {
	// write code here
	// return x % 2 == 0
	return x&1 == 0
}
