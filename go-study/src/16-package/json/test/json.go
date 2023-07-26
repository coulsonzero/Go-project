package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name  string
	Email string
	Age   int
}

func main() {
	// mapToJson()
	// jsonToMap()

}

// object -> json
func objectToJson(obj interface{}) string {
	res, _ := json.Marshal(obj)
	return string(res)
}

// json -> object(map/struct)
func jsonToObject(data string, ptr interface{}) interface{} {
	err := json.Unmarshal([]byte(data), &ptr)
	if err != nil {
		log.Fatal(err)
	}
	return ptr
}

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

// struct -> json
func structToJson() {
	s := Student{
		Name:  "John",
		Email: "john@gmail.com",
		Age:   20,
	}

	res, _ := json.Marshal(s)
	fmt.Println(string(res))
	// Output: {"Name":"John","Email":"john@gmail.com","Age":20}
}

// json -> struct
func jsonToStruct() {
	obj := []byte(`{"Name":"John","Email":"john@gmail.com","Age":20}`)
	s := Student{}
	json.Unmarshal(obj, &s)
	fmt.Println(s)
	// Output: {John john@gmail.com 20}
}

func jsonFromObjDemo() {
	jsonStr := `{	
		"name":  "John",
		"email": "john@gmail.com",
		"age":   20,
		"data": [120, 200, 150, 80, 70, 110, 130]
	}`
	var m map[string]interface{}
	fmt.Println(jsonToObject(jsonStr, m))
}
