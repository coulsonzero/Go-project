package main

import (
	"encoding/json"
	"fmt"
)

type BarData struct {
	Type     string                 `json:"type"`
	Data     []int                  `json:"data"`
	AxisLine map[string]interface{} `json:"axisLine"`
}

func main() {
	obj := []byte(`{
		"type": "bar",
		"data": [120, 200, 150, 80, 70, 110, 130],
		"axisLine": {"lineStyle": {"type": "solid", "color": "blue"}}
	}`)

	res := BarData{}
	err := json.Unmarshal(obj, &res)
	if err != nil {
		return
	}

	fmt.Println("res: ", res)
	fmt.Println("type", res.Type)
	fmt.Println("data: ", res.Data)
	fmt.Println("axisLine:", res.AxisLine)
}
