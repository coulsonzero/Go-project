package main

import (
	"bytes"
	"strings"
)

func main() {
	// a := []string{"a", "b", "c"}

	// //方式1：
	// ret := a[0] + a[1] + a[2]
	// //方式2：
	// ret := fmt.Sprintf("%s%s%s", a[0], a[1], a[2])

	// //方式3：
	// var sb strings.Builder
	// sb.WriteString(a[0])
	// sb.WriteString(a[1])
	// sb.WriteString(a[2])
	// ret := sb.String()
	// //方式4：
	// buf := new(bytes.Buffer)
	// buf.Write(a[0])
	// buf.Write(a[1])
	// buf.Write(a[2])
	// ret := buf.String()
	// //方式5：
	// ret := strings.Join(a,"")
	// fmt.Println(ret)
}

func build(arr []string) string {
	var sb strings.Builder
	for _, v := range arr {
		sb.WriteString(v)
	}
	return sb.String()
}

func buffer(arr []string) string {
	buf := new(bytes.Buffer)
	for _, v := range arr {
		buf.Write([]byte(v))
	}
	return buf.String()
}
