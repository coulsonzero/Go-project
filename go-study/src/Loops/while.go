package main

import "fmt"

func main() {
	/* for版的while */
	sum := 1
	res := 0
	for sum <= 1000 {
		res += sum
		sum++
	}
	fmt.Println(res) // 500500
}
