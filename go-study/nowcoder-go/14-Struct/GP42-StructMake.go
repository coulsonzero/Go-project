package main

import "fmt"

/**
 * question: GP42 学生信息II
 * description: 每个学生都有如下信息：姓名name，性别sex，年龄age，分数score，定义一个结构体Student，来表示小明的信息，
 *              小明的信息如下：姓名：小明，性别：true，年龄：23，分数：88, 输出并打印小明信息
 * input :
 * output:
 *         小明
 *         true
 *         23
 *         88
 */

type Student struct {
	Name  string
	Sex   bool
	Age   int
	Score int
}

func main() {
	s := Student{
		Name:  "小明",
		Sex:   true,
		Age:   23,
		Score: 88,
	}
	fmt.Println(s.Name)
	fmt.Println(s.Sex)
	fmt.Println(s.Age)
	fmt.Println(s.Score)
}
