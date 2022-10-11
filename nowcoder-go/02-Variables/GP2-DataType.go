package main

import (
	"fmt"
)

/**
 * Question: GP2 小明信息
 * Description: 已知小明的信息（姓名：小明，年龄：23，性别：男），
 *              定义三个变量，分别表示姓名（string类型），年龄（int类型），性别（男true,女false,bool类型），
 *              然后按照姓名，年龄，性别的顺序，逐行输出对应的信息。
 * input : 无
 * output:
 *         "小明"
 *          23
 *          true
 */

func main() {
	var name, age, gender = "小明", 23, true
	fmt.Printf("%s\n%d\n%t\n", name, age, gender)
}
