package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "3.1415926535"
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 3.1415926535
	}

	v64 := "-3546343826"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int64, -3546343826
	}

	num := 14141
	s := strconv.Itoa(num) // int -> string: "14141"
	fmt.Printf(s)
}
