package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"os"
)

func main() {
	f, _ := os.ReadFile("bar.json")
	res := gjson.Get(string(f), "xAxis")
	fmt.Println(res)
}
