package _6_TypeConver

import "fmt"

/**
 * Question: GP9 格式化字符串
 * Descript: 给定一个正整数，将其转换为字符串类型。
 * input : 100
 * output: "100"
 *
 * Tip: fmt.Sprintf 格式化字符串
 *      %T   数据类型和值
 *      %b   整型以二进制方式显示
 *      %d   整型以十进制方式显示
 *      %-3d 整数以十进制方式显示, 3表示输出的数字占3个字符的位置，-表示左对齐
 *      %f   浮点数
 *      %t   布尔值
 *      %s   字符串
 *      %p   指针，十六进制方式显示
 *      %v   值
 *      %+v
 *      %#v
 *      %U   Unicode 字符
 *      %%   输出 % 本体
 */

func formatstr(num int) string {
	// write code here
	return fmt.Sprintf("%d", num)
}
