package main

import (
	"errors"
	"fmt"
)

// 错误被认为是一种可以预期的结果，而异常则是一种非预期的结果
// 发生异常可能表示程序中存在bug或发生了其它不可控的问题

// 错误处理：errors.New(), log.Fatal()/log.Fatalf(), panic
// 异常捕获: defer()函数中直接调用recover(), 使用panic/log将异常抛出为明确的错误信息进行处理

// errors.New(): 抛出错误
// errors.ToJson(err)：将错误编码为json字符串

// err := errors.NewWithCode(404, "not found")
// err.(errors.Error).Code(): HTTP错误状态码

// f, err := os.Open()
// if err != nil {}

// if v, ok := m["key"]; ok {}

func test(x int) (int, error) {
	if x < 0 {
		return -1, errors.New("can't work with 42")
	}
	return x + 2, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
