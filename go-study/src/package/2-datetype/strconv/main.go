package main

import (
	"strconv"
)

func main() {
	println(strconv.Itoa(123))

	num, _ := strconv.Atoi("123")
	println(num)
}
