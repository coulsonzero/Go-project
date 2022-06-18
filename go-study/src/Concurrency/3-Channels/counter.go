package main

import "fmt"

func count(target int, nums []int, ch chan int) {
	cnt := 0
	for _, v := range nums {
		if target == v {
			cnt++
		}
	}
	ch <- cnt
}

func main() {
	nums := []int{12, 36, 12, 2, 5, 12, 36}
	//var input int
	//fmt.Scanln(&input)
	input := 12

	ch1 := make(chan int)
	ch2 := make(chan int)

	go count(input, nums[:len(nums)/2], ch1)
	go count(input, nums[len(nums)/2:], ch2)

	fmt.Println(<-ch1 + <-ch2)
}
