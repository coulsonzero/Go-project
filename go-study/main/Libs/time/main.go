package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 获取当前时间
	fmt.Println("now: ", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Printf("时间格式化: %d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	week := now.Weekday() // 周
	fmt.Printf("%s, %v, %T \n", week, week, week)
	fmt.Printf("%s, %v, %T \n", week.String(), week.String(), week.String())

	// 今年第几周
	yearTh, weekTh := now.ISOWeek()
	fmt.Printf("%d第%d周 \n", yearTh, weekTh)

	timestamp := now.Unix() // 时间戳
	fmt.Printf("现在的时间戳：%v\n", timestamp)
	timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
	fmt.Printf("时间戳->时间: %v\n", timeObj)

	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
	/*
		时间加减：Add, Sub
		时间比较：Equal, Before, After
	*/

	timeToString()
}

func timeToString() {
	t := time.Now()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)
}

func timerTick() {
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i)
	}
}

func timeFormat() string {
	now := time.Now()
	t := now.Format("2006/01/02 15:04:02 PM Mon Jan")
	return t
}

func timer() {
	start := time.Now()
	Sum()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func Sum() int {
	sum := 0
	for i := 0; i < 1000000000; i++ {
		sum += i
	}
	return sum
}
