package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := People{Name: "coulson", Age: 24}
	fmt.Printf("json data: %s", toJson(&p)) // json data : {"name":"coulson","age":24}
}

func toJson(res *People) string {
	jsons, _ := json.Marshal(res)
	return string(jsons)
}
