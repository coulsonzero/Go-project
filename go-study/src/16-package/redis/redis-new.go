package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 使用默认数据库
	})

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("无法连接到 Redis:", err)
		return
	}
	fmt.Println("连接成功:", pong)

	// string
	rdb.Set(ctx, "name", "john smith", 0)
	val, _ := rdb.Get(ctx, "name").Result()
	fmt.Println("name:", val)

	// list
	rdb.Del(ctx, "db")
	rdb.RPush(ctx, "db", "mysql", "redis", "mongodb")
	db, _ := rdb.LRange(ctx, "db", 0, -1).Result()
	fmt.Println("db: ", db)

	// hash
	rdb.HSet(ctx, "myhash", "name", "John", "age", 30)
	name, _ := rdb.HGet(ctx, "myhash", "name").Result()
	fmt.Println("myhash name:", name)

	// set
	rdb.SAdd(ctx, "myset", "java", "go", "python")
	tags, _ := rdb.SMembers(ctx, "myset").Result()
	fmt.Println("set myset:", tags)

	// zset
	scores, _ := rdb.ZRangeWithScores(ctx, "myzset", 0, -1).Result()
	fmt.Println("zset score: ", scores)
}
