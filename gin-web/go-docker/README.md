# Go Docker 的使用

```sh
$ cd <project-name>
# 打包go项目
$ go build .

# 打包docker程序
$ docker build -t go-docker:v1 .
#$ docker build -f Dockerfile -t go-docker:latest .

# 运行docker程序
$ docker run -d -p 8080:8080 go-docker:v1
#$ docker run -d -p 9000:8080 go-docker:latest
```

```shell
$ docker --version
# 查看docker打包程序
$ docker images
# 查看docker进程
$ docker ps -a
```
