package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	// connect redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	fmt.Println("connect redis success!")
	defer conn.Close()

	// SET Age 20
	_, _ = conn.Do("SET", "name", "Coulson")

	// GET name
	name, _ := redis.String(conn.Do("GET", "name"))
	fmt.Printf("Get name: %s \n", name)
}
