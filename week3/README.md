## 第三周作业
    题目：  
        - 构建本地镜像。  
        - 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。  
        - 将镜像推送至 Docker 官方镜像仓库。  
        - 通过 Docker 命令本地启动 httpserver。  
        - 通过 nsenter 进入容器查看 IP 配置。  
        作业需编写并提交 Dockerfile 及源代码。  
        提交链接：https://jinshuju.net/f/rxeJhn  

解答：  
## 将week2的代码目录拷贝至week3/，并执行docker build
     docker build -t prosperousli/httpserver:1.0 .
## 多段构建详情见Dokerfile内容：
    #编译阶段
        FROM golang:1.16-alpine as build  

    #将源码拷贝到编译容器里面
        COPY ./week2 /build/

    #强制开启gomod， 使用国内七牛云包代理
        ENV GO111MODULE=on \
            GOPROXY=https://goproxy.cn,direct

    #指定容器内默认工作目录
        WORKDIR /build

    #预下载依赖modules到容器本地cache
        RUN go mod download

    #编译go程序，关闭cgo，否则在alpine不可运行。
        RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpserver

    #运行阶段
        FROM alpine:latest as run

    #从编译容器中将编译好的程序copy到运行容器中
        COPY --from=build /build/httpserver /
    #修改可执行程序的权限，以便可以运行
        RUN chmod a+x /httpserver

    #容器启动执行命令，启动httpserver服务
        ENTRYPOINT ["/httpserver"]
## docker push到dockerhub
    push之前需要先登录dockerhub, 即docker login 输入账号密码
    docker push
    
    镜像地址 ： https://registry.hub.docker.com/repository/docker/prosperousli/httpserver  
