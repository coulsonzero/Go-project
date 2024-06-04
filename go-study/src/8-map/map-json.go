package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	json_to_map()
	// map_to_json()
}

func map_to_json() {
	m := map[string]any {
		"name": "coulsonzero",
		"age": 21,
		"country": "China",
	}


	jsonStr, _ := json.Marshal(m)
	fmt.Println(string(jsonStr))
	// {"age":21,"country":"China","name":"coulsonzero"}

}

func json_to_map() {
	jsonStr := `
		{
			"name": "coulsonzero",
			"age": 21,
			"country": "China"
		}
	`
	var m map[string]any
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
	// map[age:21 country:China name:coulsonzero]
}