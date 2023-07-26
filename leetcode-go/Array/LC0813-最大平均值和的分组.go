package main

import (
	"fmt"
	"strconv"
)

func main() {
	c, _ := strconv.ParseFloat(strconv.FormatFloat(2, 'f', 5, 64), 64)
	fmt.Printf("%v, %T \n", c, c)

	f, _ := strconv.ParseFloat("2.0", 64)
	fmt.Printf("%v, %T \n", f, f)

	h := strconv.FormatFloat(2, 'e', 5, 64)
	fmt.Printf("%v, %T \n", h, h)

	l, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("%v, %T \n", l, l)
}

func avg(arr []int) string {
	var res int
	for _, v := range arr {
		res += v
	}

	v := float64(res) / float64(len(arr))
	return strconv.FormatFloat(v, 'f', 5, 64)
}
