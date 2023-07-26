package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// MD5 加密方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func main() {
	str := "123456"
	fmt.Printf("MD5(%s): %s\n", str, MD5(str))

}

/*
MD5(123456): e10adc3949ba59abbe56e057f20f883e
*/
