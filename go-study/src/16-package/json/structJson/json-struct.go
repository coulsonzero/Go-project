package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	GetJson()
}

// GetJson struct -> json
func GetJson() {
	s := Student{
		Name:  "John",
		Email: "john@gmail.com",
		Age:   20,
	}

	res, _ := json.Marshal(&s)
	fmt.Println(string(res))
	// Output: {"Name":"John","Email":"john@gmail.com","Age":20}
}

// SetJson json -> struct
func SetJson() {
	obj := []byte(`{"Name":"John","Email":"john@gmail.com","Age":20}`)

	s := Student{}
	json.Unmarshal(obj, &s)
	fmt.Println(s)
	// Output: {John john@gmail.com 20}
}
