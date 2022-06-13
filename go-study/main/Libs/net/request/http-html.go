package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

/**
发送get请求html资源，然后写入html文件中
*/

func main() {
	// Get请求html
	resp, _ := http.Get("http://www.coulsonzero.cn")
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	// 创建文件
	os.Remove("index.html")
	os.Create("index.html")

	res := ""
	// 爬取html
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10000; i++ {
		// fmt.Println(scanner.Text())
		res += scanner.Text()
	}

	// 写入文件
	os.WriteFile("index.html", []byte(res), os.ModePerm)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
