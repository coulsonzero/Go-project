package main

import "fmt"

func main() {
	countries := map[string]string{
		"us": "USA",
		"fr": "France",
		"cn": "China", // 末尾加逗号，或者将大括号放在此行！
	}
	fmt.Println(countries) // map[cn:China fr:France us:USA]
}
