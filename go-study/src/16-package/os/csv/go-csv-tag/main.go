package main

import (
	"fmt"
	"os"
	"path"

	csvtag "github.com/artonge/go-csv-tag/v2"
)

type Demo struct {
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

func main() {
	dir, _ := os.Getwd()
	filepath := path.Join(dir, "/conf/file.csv")

	tab := []Demo{}
	file, _ := os.Open(filepath)
	csvtag.LoadFromReader(file, &tab)
	fmt.Println(tab) // [{name1 1 1.2} {name2 2 2.3} {name3 3 3.4}]

	// fmt.Printf("%v\n", reflect.TypeOf(&tab).Elem().Kind())	// slice
}
