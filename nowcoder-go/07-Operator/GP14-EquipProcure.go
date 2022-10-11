package main

import "sort"

/**
 * question: GP14 器材采购
 * description: 现要采购一批器材，同时有三个厂商在卖，单位价格分别为 a,b,c, 比较这三家厂商的价格，返回 价格的差值=最高价格-最低价格。
 * input : 2,3,1
 * output: 2
 */

func compare(x int, y int, z int) int {
	// write code here
	nums := []int{x, y, z}
	sort.Ints(nums)
	return nums[len(nums)-1] - nums[0]

}
