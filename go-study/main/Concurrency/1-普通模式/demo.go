package main

import (
	"fmt"
)

func out(from, to int) {
	for i := from; i <= to; i++ {
		fmt.Println(i)
	}
}

func main() {
	out(0, 5)
	out(6, 10)
}

/*
0
1
2
3
4
5
6
7
8
9
10
*/
