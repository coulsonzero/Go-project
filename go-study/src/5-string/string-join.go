package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	a := []string{"a", "b", "c"}

	// 方式1：
	ret1 := a[0] + a[1] + a[2]

	// 方式2：
	ret2 := fmt.Sprintf("%s%s%s", a[0], a[1], a[2])

	// 方式3：
	ret3 := strings.Join(a, "")

	// 方式4：
	var b strings.Builder
	b.WriteString(a[0])
	b.WriteString(a[1])
	b.WriteString(a[2])

	ret4 := b.String()
	ret4_1 := build(a)

	// 方式5：
	buf := new(bytes.Buffer)
	buf.Write([]byte(a[0]))
	buf.Write([]byte(a[1]))
	buf.Write([]byte(a[2]))

	ret5 := buf.String()
	ret5_1 := buffer(a)

	fmt.Println(ret1)
	fmt.Println(ret2)
	fmt.Println(ret3)
	fmt.Println(ret4)
	fmt.Println(ret5)

	fmt.Println(ret4_1)
	fmt.Println(ret5_1)
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
