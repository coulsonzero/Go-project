package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GetHtml 获取网页文本
func GetHtml(params string) string {
	url := fmt.Sprintf("http://voice.baidu.com/act/newpneumonia/newpneumonia/%s", params)
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	html := string(body)

	return html
}

// WriteHtml 写入html
func WriteHtml() {
	file, _ := os.Create("newpneumonia.html")
	defer file.Close()
	file.WriteString(GetHtml("?city=北京-北京"))
}

func main() {
	GetHtml("")
	GetHtml("?city=北京-北京")
}