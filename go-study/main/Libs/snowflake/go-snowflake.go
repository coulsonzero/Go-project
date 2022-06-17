package main

import (
	"fmt"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/golang/glog"
	"sync"
	"time"
)

/**
 * Snowflake
 * 由twitter设计，使用snowflake算法为分布式系统生成id。
 * 单机系统中数据库自增id有利于排序和索引，但是在分布式系统中使用它就会引起冲突，自增id非常容易被爬虫爬去数据
 * uuid是随机字符串，无法被排序
 */

func main() {
	var wg sync.WaitGroup
	// 创建Snowflake对象
	s, err := snowflake.NewSnowflake(int64(1), int64(1))
	if err != nil {
		glog.Error(err)
		return
	}
	var check sync.Map
	t1 := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 生成唯一ID
			val := s.NextVal()
			// 打印生成的id
			fmt.Println(val)
			if _, ok := check.Load(val); ok {
				// id冲突检查
				glog.Error(fmt.Errorf("error#unique: val:%v", val))
				return
			}
			check.Store(val, 0)
			if val == 0 {
				glog.Error(fmt.Errorf("error"))
				return
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(t1)
	glog.Infof("generate 20k ids elapsed: %v", elapsed)
}

/*
325615759504052224
325615759504052225
325615759504052226
325615759504052227
325615759504052228
325615759504052229
325615759504052230
*/
