package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn := RedisConnect()
	defer conn.Close()

	db := Conn{conn}
	db.SetString("score", 97.2)
	db.GetString("score")
}

// RedisConnect connect redis
func RedisConnect() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return nil
	}
	fmt.Println("connect redis success!")

	return conn
}

type Conn struct {
	redis.Conn
}

// SetString SET filed value
func (conn *Conn) SetString(field string, value interface{}) {
	_, _ = conn.Do("SET", field, value)
}

// GetString GET field
func (conn *Conn) GetString(field string) {
	res, _ := redis.String(conn.Do("GET", field))
	fmt.Printf("Get %s: %s \n", field, res)
}
