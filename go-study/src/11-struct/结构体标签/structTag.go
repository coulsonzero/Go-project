package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name" level:"12"`
	Age  int    `json:"age"`
}

func main() {
	s := Student{}
	t := reflect.TypeOf(s)
	if fieldName, ok := t.FieldByName("Name"); ok {
		fmt.Println(fieldName.Tag.Get("json"))  // "name"
		fmt.Println(fieldName.Tag.Get("level")) // "12"
	}

}
