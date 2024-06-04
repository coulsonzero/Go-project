package main

import "fmt"

func main() {
	ans := 0
	for i := 1; i <= 1000; i++ {
		ans += i
	}

	/* for版的while */
	res := 0
	i := 1
	for i <= 1000 {
		res += i
		i++
	}
	fmt.Println(res) // 500500
}
