package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

func main() {
	db := RedisConnect()
	defer db.Close()

	// 认证信息
	// if _, err := conn.Do("username", "password"); err != nil {
	// 	conn.Close()
	// }

	// 删：key键
	// db.DelKey("score")
	// 删：单条数据
	// db.DelSortSetByField("score", "Tom")

	// 增
	db.SetSortSet("score", 96, "Coulson")
	db.SetSortSet("score", 92, "Tom")
	db.SetSortSet("score", 97, "Jack")
	// 查：长度
	fmt.Println(db.GetSortSetLength("score"))
	// 查：排名(asc: 升序, desc: 降序, int: 降序前n条数据)
	nameRank, scoreRank := db.RankSortSet("score", "asc")
	fmt.Println(nameRank, scoreRank)
	// 查：单条数据
	fmt.Println(db.GetSortSetByField("score", "Tom"))

}

// RedisConnect connect redis
func RedisConnect() *Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return nil
	}
	fmt.Println("connect redis success!")

	return &Conn{conn}
}

type Conn struct {
	redis.Conn
}

func (conn *Conn) DelKey(key string) {
	_, err := conn.Do("DEL", key)
	if err != nil {
		log.Printf("Failed to del the key: %s \n", key)
	}
	fmt.Printf("success to del the key: %s! \n", key)
}

// SetSortSet ZADD key ...
func (conn *Conn) SetSortSet(key string, value int, field string) {
	_, err := conn.Do("ZADD", key, value, field)
	if err != nil {
		// error: WRONGTYPE Operation against a key holding the wrong kind of value
		log.Printf("Already exist duplicate keys: `%s` \n", key)
		log.Fatal(err)
	}
	log.Printf("success to set key: %s %d %s", key, value, field)
}

// GetSortSetLength ZCARD key
func (conn *Conn) GetSortSetLength(key string) interface{} {
	res, err := conn.Do("ZCARD", key)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// RankSortSet sort
func (conn *Conn) RankSortSet(key string, limit interface{}) (nameArr []string, valueArr []string) {
	// ZRANGE score 0 -1 WITHSCORES        # 递增排列
	// ZRANGE salary 0 <n> WITHSCORES      # [0, n)
	// fmt.Println(reflect.TypeOf(limit).Name())
	cmd, index := "ZREVRANGE", -1
	switch {
	case limit == "asc":
		cmd = "ZRANGE"
		index = -1
	case limit == "desc":
		cmd = "ZREVRANGE"
		index = -1
	case reflect.TypeOf(limit).Name() == "int" && int(reflect.ValueOf(limit).Int()) > 0:
		index = int(reflect.ValueOf(limit).Int()) - 1
	default:
		cmd = "ZREVRANGE"
		index = -1
	}
	// scoreMap, err := redis.StringMap(conn.Do("ZRANGE", "score", 0, -1, "withscores"))
	scoreMap, err := redis.StringMap(conn.Do(cmd, "score", 0, index, "withscores"))
	if err != nil {
		log.Printf("Failed to rank by key %s \n", key)
	}
	// sort
	for name := range scoreMap {
		nameArr = append(nameArr, name)
		valueArr = append(valueArr, scoreMap[name])
	}
	return nameArr, valueArr
}

// GetSortSetByField 获取单个数据
func (conn *Conn) GetSortSetByField(field string, value interface{}) interface{} {
	res, err := redis.Int(conn.Do("ZSCORE", field, value))
	if err != nil {
		log.Printf("can't find %v of %v \n", field, value)
	}
	return res
}

// DelSortSetByField 删除单条数据
func (conn *Conn) DelSortSetByField(field string, value interface{}) interface{} {
	res, err := conn.Do("ZREM", field, value)
	if err != nil {
		log.Printf("failed to del %v by %v \n", field, value)
	}
	return res
}
