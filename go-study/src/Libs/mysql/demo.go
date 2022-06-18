package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

var db *sql.DB

type User struct {
	id    int
	name  string
	email string
}

func main() {
	if err := initDB(); err != nil {
		fmt.Println(err, "连接失败")
	} else {
		fmt.Println("连接成功")
	}
	queryAll()
	// insertDatas2()

	// parms := []string{"name", "email", "password"}
	// data := []map[string]string{
	// 	{"username": "test1", "email": "test1@gmail.com", "password": "123456"},
	// 	{"username": "test2", "email": "test2@gmail.com", "password": "admin123"},
	// }
	// insertDatas3(parms, data)

}

// id查询一条数据
func query() {
	sqlStr := "select id, name, email from users where id = ?"
	var u User
	if err := db.QueryRow(sqlStr, 2).Scan(&u.id, &u.name, &u.email); err != nil {
		fmt.Println("scan fail: ", err)
		return
	}
	fmt.Printf("id: %d, name: %s, email: %s\n", u.id, u.name, u.email)
}

// 查询全部数据
func queryAll() {
	sqlStr := "select id, name, email from users where id > ?"
	rows, _ := db.Query(sqlStr, 0)
	defer rows.Close()
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.id, &u.name, &u.email); err != nil {
			fmt.Println("scan fail: ", err)
			return
		}
		fmt.Printf("id: %d, name: %s, email: %s\n", u.id, u.name, u.email)
	}
}

// 插入数据
func insertData() {
	sqlStr := "insert into users(name, email) values(?, ?)"
	res, err := db.Exec(sqlStr, "aqwq", "aqwq@gmail.com")
	if err != nil {
		fmt.Println("res fail: ", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("lastInsertId fail: %s", err)
	}
	fmt.Printf("id: %d insert success!", id)
}

// 插入多条数据
func insertDatas() {
	sqlStr := "insert into users(name, email) values (?, ?), (?, ?)"

	res, err := db.Exec(sqlStr, "tom", "tom@gmail.com", "john", "john@gmail.com")
	if err != nil {
		fmt.Println("res fail: ", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("lastInsertId fail: %s", err)
	}
	fmt.Printf("id: %d insert success!", id)
}

// 插入多条数据
func insertDatas2() {
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
	fmt.Println("sqlStr:", sqlStr)
	fmt.Println("vals: ", vals)

	res, _ := db.Exec(sqlStr, vals...) // vals...: 解构
	id, _ := res.LastInsertId()
	fmt.Printf("lastId: %d insert success!", id)
}

func insertDatas3(parms []string, data []map[string]string) {
	fmt.Println(parms)

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

	// fmt.Println("sqlStr:", sqlStr)
	// fmt.Println("vals: ", vals)

	res, _ := db.Exec(sqlStr, vals...) // vals...: 解构
	id, _ := res.LastInsertId()
	fmt.Printf("lastId: %d insert success!", id)

}

func test() {
	db, err := sql.Open("mysql", "root:root@/golang_api")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func initDB() (err error) {
	dsn := "root:root@tcp(localhost:3306)/golang_api?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
