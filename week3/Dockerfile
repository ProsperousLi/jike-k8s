# 编译阶段
FROM golang:1.16-alpine as build

# 将源码拷贝到编译容器里面
COPY ./week2 /build/

# 强制开启gomod， 使用国内七牛云包代理
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 指定容器内默认工作目录
WORKDIR /build

# 预下载依赖modules到容器本地cache
RUN go mod download

# 编译go程序，关闭cgo，否则在alpine不可运行。
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpserver

# 运行阶段
FROM alpine:latest as run

# 从编译容器中将编译好的程序copy到运行容器中
COPY --from=build /build/httpserver /
# 修改可执行程序的权限，以便可以运行
RUN chmod a+x /httpserver

# 容器启动执行命令，启动httpserver服务
ENTRYPOINT ["/httpserver"]
