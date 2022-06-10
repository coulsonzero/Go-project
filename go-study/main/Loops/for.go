package main

import "fmt"

func main() {

}

// 9乘9格式化
func MultiplyTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			// %-3d ——以10进制显示，3表示输出的数字占3个字符的位置，-表示左对齐；
			fmt.Printf("%d*%d=%-3d ", j, i, i*j)
		}
		fmt.Println()
	}
}

// 回文字符串
func isEve() bool {
	s := "14141"
	j := len(s) - 1
	for i := 0; i < j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

/**
 * GP36. 二维数组转一维数组
 * input:  [[1,2,3],[4,5,6],[7,8,9]]
 * output: [1,2,3,4,5,6,7,8,9]
 */
func transform(array [][]int) []int {
	// write code here
	res := []int{}
	for _, v := range array {
		res = append(res, v...)
	}
	return res
}
