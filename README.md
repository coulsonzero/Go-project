# Go-project


<p align='center'>
<a href="https://go.dev"><img alt="Support Go version" src="https://img.shields.io/badge/Go-v1.18-26C2F0"></a>
<a href="https://github.com/coulsonzero/gopkg/blob/master/LICENSE" ><img src="https://img.shields.io/github/license/coulsonzero/gopkg?label=License"></a>
<a href="https://github.com/coulsonzero/gopkg"><img alt="visitor" src="https://visitor-badge.laobi.icu/badge?page_id=coulsonzero.go-project"></a>
<img src="https://img.shields.io/badge/Code%20rows-9416-success">

</p>

```md
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                             580           3224           4002          14259
XML                             11              0              0           2046
JavaScript                      40             75            150            320
SQL                             16             33             16            274
Markdown                         5             46              0            148
JSON                             3              1              0             71
YAML                             4              3              0             60
INI                              2              2              0             31
make                             2             12             11             30
Bourne Shell                     1              1              1             14
Dockerfile                       1              8              6             14
C++                              1              0              0             13
HTML                             1              0              0             10
CSV                              2              0              0              9
-------------------------------------------------------------------------------
SUM:                           669           3405           4186          17299
-------------------------------------------------------------------------------
```



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
