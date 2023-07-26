package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	nums := []int{0, 2, 4, 6, 8}

	res := sum(nums...)
	fmt.Println(res) // 20
}
