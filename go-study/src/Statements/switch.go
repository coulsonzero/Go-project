package main

import "fmt"

func main() {
	x := 8
	switch y := x % 2; y {
	case 1:
		fmt.Println("1")
		// statement(s)        //不需要break
	default:
		fmt.Println("0")
		// statement(s)
	}
}

// GP34. 推箱子
func pushBox(forwards string) bool {
	// write code here
	x, y := 0, 0
	charArr := []byte(forwards)
	for _, v := range charArr {
		switch v {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		default:
			break
		}
	}
	return x == 0 && y == 0
}
