package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr []string
	for j := 0; j <= 1000; j++ {
		var x string
		fmt.Scan(&x)
		arr = append(arr, x)
	}

	// arr := []string{"5", "2", "C", "D", "+"}
	// arr := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	// arr := []string{"1"}
	// arr := []string{"+"}

	fmt.Println(solution(arr))
}

func solution(arr []string) int {
	var res []int
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case "C":
			if len(res) < 1 {
				return -1
			}
			res = res[:len(res)-1]
		case "D":
			if len(res) < 1 {
				return -1
			}
			res = append(res, 2*res[len(res)-1])
		case "+":
			if len(res) < 2 {
				return -1
			}
			res = append(res, res[len(res)-1]+res[len(res)-2])
		default:
			num, err := strconv.Atoi(arr[i])
			if err != nil {
				fmt.Println("please enter a valid char !")
			}
			res = append(res, num)
		}
	}

	fmt.Println(res)
	return sum(res)
}

func sum(nums []int) (res int) {
	for _, v := range nums {
		res += v
	}
	return res
}
