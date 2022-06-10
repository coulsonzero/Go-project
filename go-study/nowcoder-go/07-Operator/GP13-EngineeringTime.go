package main

import (
	"fmt"
)

/**
 * question: GP13 工程时间
 * description: 某项工程完成用了97天，用了多少个星期，零几天, 按照星期，天数逐行打印输出。
 * input: 无
 * output: 12
 *         6
 */

func main() {
	fmt.Println(97 / 7)
	fmt.Println(97 % 7)
}
