# Swagger

```sh
# 安装swag (go-v1.17版本使用install)
$ go install github.com/swaggo/swag/cmd/swag
# 查看swag是否安装成功
$ swag -v
# 生成docs文档
$ swag init
# 启动项目, 并查看接口文档: http://localhost:8080/swagger/index.html
$ go run main.go
```
