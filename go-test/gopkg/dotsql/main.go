package main

import (
	"fmt"

	"github.com/coulsonzero/gopkg/pro/dotsql"
)

func main() {
	dotsql, _ := dotsql.LoadFromFile("./gopkg/zsql/query.sql")
	fmt.Println(dotsql.LookupQuery("create-user")) // INSERT INTO users (name, email) VALUES(?, ?);
}
