package main

import "fmt"

/**
 * question: GP43 学生信息III
 * description: 每个学生都有如下信息：姓名name，性别sex，年龄age，分数score，地址信息address,其中address地址信息又包含城市city,省份province等信息， 定义一个结构体Student，来表示小明的信息，小明的信息如下：姓名：小明，性别：true，年龄：23，分数：88，province：湖南省，city：长沙市
 * input :
 * output:
 *        小明
 *        true
 *        23
 *        88
 *        湖南省
 *        长沙市
 */

type Address struct {
	city     string
	province string
}

type Student struct {
	name    string
	sex     bool
	age     int
	score   int
	address Address
}

func main() {
	s := Student{"小明", true, 23, 88, Address{"湖南省", "长沙市"}}
	fmt.Println(s.name)
	fmt.Println(s.sex)
	fmt.Println(s.age)
	fmt.Println(s.score)
	fmt.Println(s.address.city)
	fmt.Println(s.address.province)
}
