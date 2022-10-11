package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {
	const obj = `{
	"type": "bar",
	"data": [120, 200, 150, 80, 70, 110, 130],
	"axisLine": {"lineStyle": {"type": "solid", "color": "blue"}}
}`
	value, _ := sjson.Set(obj, "axisLine.lineStyle.color", "skyblue")
	fmt.Println(value)
	fmt.Println(obj)

	res := gjson.Get(obj, "axisLine.lineStyle")
	fmt.Println(res)
}

func getJson() {
	const obj = `{
	"type": "bar",
	"data": [120, 200, 150, 80, 70, 110, 130],
	"axisLine": {"lineStyle": {"type": "solid", "color": "blue"}}
}`
	res := gjson.Get(obj, "axisLine.lineStyle")
	fmt.Println(res)
}

func setJson() {
	const obj = `{
	"type": "bar",
	"data": [120, 200, 150, 80, 70, 110, 130],
	"axisLine": {"lineStyle": {"type": "solid", "color": "blue"}}
}`
	sjson.Set(obj, "axisLine.lineStyle.color", "skyblue")
}
