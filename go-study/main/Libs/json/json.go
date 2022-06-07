package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Email string
	Age   int
}

func main() {
	SetJson2()
}

func GetJson() {
	s := Student{
		Name:  "John",
		Email: "john@gmail.com",
		Age:   20,
	}

	// struct => json
	res, _ := json.Marshal(s)
	fmt.Println(string(res))
	// Output: {"Name":"John","Email":"john@gmail.com","Age":20}
}

func SetJson() {
	obj := []byte(`{"Name":"John","Email":"john@gmail.com","Age":20}`)

	s := Student{}
	json.Unmarshal(obj, &s)
	fmt.Println(s)
	// Output: {John john@gmail.com 20}
}

func SetJson2() {
	obj := []byte(`{
	"type": "bar",
	"data": [120, 200, 150, 80, 70, 110, 130],
	"axisLine": {"lineStyle": {"type": "solid", "color": "blue"}}
}`)
	var s interface{}
	json.Unmarshal(obj, &s)
	fmt.Println(s)

}
