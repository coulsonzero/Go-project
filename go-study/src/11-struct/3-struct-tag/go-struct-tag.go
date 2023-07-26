package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name" level:"12" csv:"name"`
	Age  int    `json:"age" csv:"age"`
}

func getStructTag(s any, field string, tag string) string {
	t, _ := reflect.TypeOf(s).FieldByName(field)
	return t.Tag.Get(tag)
}

func main() {
	// getStructTags()

	// var s Student
	// t, _ := reflect.TypeOf(s).FieldByName("Name")
	// fmt.Println(t.Tag.Get("level")) // 12

	// println(getStructTag(Student{}, "Name", "level"))

	fmt.Printf("%v", getStructTags(Student{}, "json")) // ["name", "age"]
}

// 获取struct所有指定的tag名称
func getStructTags(s any, tag string) (res []string) {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		res = append(res, t.Field(i).Tag.Get(tag))
	}
	return
}

func getTagName() {
	s := Student{}
	t := reflect.TypeOf(s)
	if fieldName, ok := t.FieldByName("Name"); ok {
		fmt.Println(fieldName.Tag.Get("json"))  // "name"
		fmt.Println(fieldName.Tag.Get("level")) // "12"
	}
}
