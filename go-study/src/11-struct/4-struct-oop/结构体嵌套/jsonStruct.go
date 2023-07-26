package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Title string
	URL   string
}

type Response struct {
	// 结构体嵌套：常用于返回或解析网络API接口的Response请求
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func main() {
	jsonStr := `{
		"data": {
			"children":[{
				"data": {
					"title": "Coulson's blog",
					"url": "http://docs.coulsonzero.top"
				}
			}]
		}
	}`
	res := Response{}
	json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(res)
	// {{[{{Coulson's blog http://docs.coulsonzero.top}}]}}
}
