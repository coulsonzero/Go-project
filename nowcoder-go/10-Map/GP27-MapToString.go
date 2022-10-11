package main

import "fmt"

/**
 * question: GP27 成绩表
 * description: 某大学宿舍6人的数学成绩分别为 小明：60，小王：70，张三：95，李四：98，王五：100，张伟：88 ，
 *              现要将六人的成绩录入成绩表中，这个成绩表用一个map来表示，成绩表的键为宿舍成员的姓名，值为对应的分数 。打印该成绩表
 * input :
 * output: map[小明:60 小王:70 张三:95 张伟:88 李四:98 王五:100]
 */

func main() {
	scoreMap := map[string]int{
		"小明": 60,
		"小王": 70,
		"张三": 95,
		"张伟": 88,
		"李四": 98,
		"王五": 100,
	}
	fmt.Println(scoreMap)
}
