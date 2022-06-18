# Go-project

## 开发环境搭建

### 安装

[Go 环境安装](https://golang.google.cn/dl/)

检查版本: `go version`

### 配置环境变量

1. Windows

`我的电脑` -> `属性` -> `高级系统设置` -> `环境变量` -> `系统变量` -> `path` -> 复制 bin 路径 `...go\bin`

2. Mac

```sh
vim ~/.bash_profile
export PATH=$PATH:/usr/local/go/bin
source ~/.bash_profile
```

## Code

### hello.go

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world!")
}
```

### linux

[Go Docs](https://golang.google.cn/doc/)

```
go mod init example.com/hello
```
