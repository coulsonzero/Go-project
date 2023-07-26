package main

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := make(map[string]int)
	for _, email := range emails {
		arr := strings.Split(email, "@")
		name, domain := arr[0], arr[1]
		// 1.替换所有的句点('.')
		name = strings.ReplaceAll(name, ".", "")
		// 2.忽略第一个加号('+')后面的所有内容
		if strings.Index(name, "+") != -1 {
			name = name[:strings.Index(name, "+")]
		}
		// 除重统计
		m[name+"@"+domain]++
	}

	// fmt.Printf("%v", m)

	return len(m)
}
