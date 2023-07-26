package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	t := strconv.Itoa(num)
	for i := 0; i < len(t); i++ {
		m, _ := strconv.Atoi(string(t[i]))
		fmt.Printf("%c", m+97)
	}

	return 0
}

func main() {
	translateNum(12258)
	fmt.Printf("%c", 1+97)
}
