package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db := ConnectDB()
	ConnectTest(db)
	// CreateTable(db)
	// InsertData(db)
	// QueryData(db)
	QueryAllData(db)
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_study?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}

func ConnectTest(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic("ping 失败")
	}
	fmt.Println("pong")
}

// Create Table
func CreateTable(db *sql.DB) {
	query := `
-- 	DROP TABLE IF EXISTS users;
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );
    `

	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Create table failed!")
		return
	}
	fmt.Println("Create table users success.")
}

// INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)
func InsertData(db *sql.DB) {
	username := "Rel Poul"
	password := "admin123"
	createdAt := time.Now()

	result, _ := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	userID, _ := result.LastInsertId()
	fmt.Printf("Insert success: %v. \n", userID)
}

// SELECT id, username, password, created_at FROM users WHERE id = ?
func QueryData(db *sql.DB) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	// Query the database and scan the values into out variables. Don't forget to check for errors.
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	if err != nil {
		panic("Query data failed!")
		return
	}
	fmt.Printf("id = %d, username = %s, password = %s, createdAt = %s\n", id, username, password, createdAt)
}

func QueryAllData(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}
	rows, _ := db.Query(`SELECT id, username, password, created_at FROM users`)
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", users)
}

func DeleteData(db *sql.DB, id int) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
}
