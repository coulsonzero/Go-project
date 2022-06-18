package main

import "fmt"

/**
 * goto跳转需要配合return来使用，否则每条goto语句都会被顺序输出
 */

func main() {
	fmt.Println(gotoFun(13))
}

func gotoFun(age int) string {
	if age > 18 {
		goto OK
	} else if age < 6 {
		goto NO
	} else {
		return "right"
	}

NO:
	return "too young!"
OK:
	return "more than 18 years old"
}
