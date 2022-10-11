package main

import (
	"encoding/json"
	"fmt"
)

// map -> json
func mapToJson() {
	m := map[string]interface{}{
		"name":  "John",
		"email": "john@gmail.com",
		"age":   20,
	}

	jsonStr, _ := json.Marshal(&m)
	fmt.Println(string(jsonStr))
	// Output: {"age":20,"email":"john@gmail.com","name":"John"}
}

// json -> map
func jsonToMap() {
	jsonStr := `{
		"name":  "John",
		"email": "john@gmail.com",
		"age":   20,
		"data": [120, 200, 150, 80, 70, 110, 130]
	}`
	var m map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &m)
	fmt.Println(m)
	// Output: map[age:20 data:[120 200 150 80 70 110 130] email:john@gmail.com name:John]
}
