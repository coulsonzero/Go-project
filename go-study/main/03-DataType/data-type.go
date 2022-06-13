package main

import "fmt"

func main() {
	dataType()
}

func dataType() {
	var (
		name     string
		age      int
		gender   bool
		score    float64
		skills   []string
		scoreMap map[string]int
	)
	fmt.Printf("name =  %s, age = %d, gender = %t, score = %f \n", name, age, gender, score)
	// name = "", age = 0, gender = false, score = 0.000000
	fmt.Printf("%T, %T, %T, %T \n", name, age, gender, score)
	// string, int, bool, float64
	fmt.Printf("%v %T \n", skills, skills)
	// [] []string
	fmt.Printf("%v, %T \n", scoreMap, scoreMap)
	// map[], map[string]int
}
