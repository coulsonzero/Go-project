package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var DB *sql.DB

type User struct {
	id    int
	name  string
	email string
}

func main() {
	db := ConnectDB()
	defer db.Close()
	// Query(2)
	// QueryAll()

	// Insert()
	// Insert2()
	// Insert3()

	// parms := []string{"name", "email", "password"}
	// data := []map[string]string{
	// 	{"username": "test1", "email": "test1@gmail.com", "password": "123456"},
	// 	{"username": "test2", "email": "test2@gmail.com", "password": "admin123"},
	// }
	// InsertAny(parms, data)

	// Update()
	// Delete()
}

func ConnectDB() *sql.DB {
	dsn := "root:root@tcp(localhost:3306)/golang_api?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败！")
	}

	DB = db
	return db
}

// id查询一条数据
func Query(id int) {
	sqlStr := "select id, name, email from users where id = ?"
	var u User
	err := DB.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id: %d, name: %s, email: %s\n", u.id, u.name, u.email)
}

// 查询全部数据
func QueryAll() {
	sqlStr := "select id, name, email from users where id > ?"
	rows, _ := DB.Query(sqlStr, 0)
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s, email: %s\n", u.id, u.name, u.email)
	}
}

// 插入数据
func Insert() {
	sqlStr := "insert into users(name, email) values(?, ?)"
	res, _ := DB.Exec(sqlStr, "aqwq", "aqwq@gmail.com")
	id, _ := res.LastInsertId()
	fmt.Printf("id: %d insert success!", id)
}

// 插入多条数据
func Insert2() {
	sqlStr := "insert into users(name, email) values (?, ?), (?, ?)"

	res, _ := DB.Exec(sqlStr, "tom", "tom@gmail.com", "john", "john@gmail.com")
	id, _ := res.LastInsertId()
	fmt.Printf("id: %d insert success!", id)
}

// 插入多条数据
func Insert3() {
	data := []map[string]string{
		{"name": "test1", "email": "test1@gmail.com"},
		{"name": "test2", "email": "test2@gmail.com"},
		{"name": "test3", "email": "test3@gmail.com"},
	}
	sqlStr := "insert into users(name, email) values"
	vals := []interface{}{}
	for index, row := range data {
		if index == len(data)-1 {
			sqlStr += "(?, ?)"
		} else {
			sqlStr += "(?, ?), "
		}
		vals = append(vals, row["name"], row["email"])
	}

	res, _ := DB.Exec(sqlStr, vals...) // vals...: 解构
	id, _ := res.LastInsertId()
	fmt.Printf("lastId: %d insert success!", id)
}

func InsertAny(parms []string, data []map[string]string) {
	sqlStr := "insert into users("
	for i, v := range parms {
		if i == len(parms)-1 {
			sqlStr += v
		} else {
			sqlStr += v + ", "
		}
	}
	sqlStr += ") values "

	vals := []interface{}{}
	sqlQuery := fmt.Sprintf("(%s)", strings.Join(strings.Split(strings.Repeat("?", len(parms)), ""), ","))
	for index, row := range data {
		if index == len(data)-1 {
			sqlStr += sqlQuery
		} else {
			sqlStr += sqlQuery + ", "
		}

		for _, v := range parms {
			vals = append(vals, row[v])
		}
	}

	res, _ := DB.Exec(sqlStr, vals...) // vals...: 解构
	id, _ := res.LastInsertId()
	fmt.Printf("lastId: %d insert success!", id)
}

func Update() {
	sqlStr := "update users set name = ?, email = ? where id = ?"
	res, _ := DB.Exec(sqlStr, "yop", "yop@gmail.com", 10)
	affected, _ := res.RowsAffected()
	fmt.Println(affected) // return 1: success
}

func Delete() {
	sqlStr := "delete from users where id = ?"
	res, _ := DB.Exec(sqlStr, 12)
	rows, _ := res.RowsAffected()
	fmt.Println(rows) // return 1: success
}
