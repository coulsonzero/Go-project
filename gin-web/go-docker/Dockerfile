FROM golang:1.18-alpine

# 作者署名
MAINTAINER CoulsonZero

# golang 环境变量
ENV GO111MODULE="on" \
    GOPROXY="https://goproxy.cn,direct" \
    CGO_ENABLED="0" \
    GOOS="linux" \
    GOARCH="amd64"

# 切换到当前工作目录
WORKDIR /project/go-docker

# Copy go.mod go.sum and 下载golang环境依赖
COPY go.* ./
RUN go mod download

# Copy 当前工作目录下的所有文件到docker目录中
COPY . .
RUN go build -o /project/go-docker/v1 .

EXPOSE 8080

# 声明docker目录
ENTRYPOINT ["/project/go-docker/v1"]

