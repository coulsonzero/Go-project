package main

import "fmt"

/**
 * Title: GP4 国家名称
 * Question: 定义三个常量来分别表示中国，英国，美国这三个国家的名称，并按照中国，英国，美国的顺序逐行打印
 * input : 无
 * output:
 *		  "中国"
 *		  "英国"
 *        "美国"
 */

func main() {
	const cn, uk, us = "中国", "英国", "美国"
	fmt.Printf("%s\n%s\n%s", cn, uk, us)

}
