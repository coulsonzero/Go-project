package main

import (
	"github.com/coulsonzero/gopkg/pro/sqlz"
)

func main() {
	zeroSql, _ := sqlz.Load("./gopkg/zsql/query.sql")
	// first
	println(zeroSql.LookupQuery("create-user")) // INSERT INTO users (name, email) VALUES(?, ?);

	// secend
	println(zeroSql.LookupQueryAny("drop-users-table")) //  DROP TABLE users;

	// third
	println(zeroSql.Get("type").QueryAny("find-one-user-by-email")) // SELECT id,name,email FROM users WHERE email = ? LIMIT 1;
	println(zeroSql.Get("tag").QueryAny("drop-users-table"))        // DROP TABLE users;

}
