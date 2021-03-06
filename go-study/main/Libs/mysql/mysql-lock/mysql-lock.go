package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * 分布式锁：FOR UPDATE
 */

var db *sql.DB

func MySQLLock() (*sql.Tx, error) {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_study")
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		// FOR UPDATE：分布式锁
		fmt.Printf("SELECT * FROM methodLock WHERE `method_name`='method_name' FOR UPDATE", err)
		return nil, nil
	}
	return tx, err
}

func main() {
	res, _ := MySQLLock()
	// 执行逻辑
	fmt.Println("执行逻辑")
	res.Commit() // 提交事务 (解锁)
	fmt.Println("exec transaction success!")
}
