package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	db := RedisDB()
	defer db.Close()
	db.SetString("score", 97.2)
	println(db.GetString("score"))
}

type RedisConn struct {
	redis.Conn
}

// RedisConn connect redis
func RedisDB() (c *RedisConn) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(fmt.Errorf("connect redis error:", err))
	}
	log.Println("connect redis success!")
	return &RedisConn{conn}
}

// SetString SET filed value
func (conn *RedisConn) SetString(field string, value interface{}) error {
	_, err := conn.Do("SET", field, value)
	return err
}

// GetString GET field
func (conn *RedisConn) GetString(field string) string {
	res, err := redis.String(conn.Do("GET", field))
	if err != nil {
		log.Fatal(err)
	}
	return res
}
