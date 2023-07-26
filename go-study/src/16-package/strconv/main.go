package main

import (
	"strconv"
)

func main() {
	println(strconv.Itoa(123))

	num, _ := strconv.Atoi("123")
	println(num)

	// int64(10) -> int(2)
	println(strconv.FormatInt(int64(11), 2)) // 1011

}
