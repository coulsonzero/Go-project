package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn := redisConnect()
	defer conn.Close()

	setString(conn, "country", "China")
	getString(conn, "country")

}

// connect redis
func redisConnect() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return nil
	}
	fmt.Println("connect redis success!")

	return conn
}

// setString SET filed value
func setString(conn redis.Conn, field string, value interface{}) {
	_, _ = conn.Do("SET", field, value)
}

// getString GET field
func getString(conn redis.Conn, field string) {
	res, _ := redis.String(conn.Do("GET", field))
	fmt.Printf("Get %s: %s \n", field, res)
}
