package main

import (
	"fmt"
)

/**
* LC.121 买卖股票的最佳时机
* Input: prices = [7,1,5,3,6,4]
* Output: 5
* 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
* 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票
 */

// 贪心算法
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 累加
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
