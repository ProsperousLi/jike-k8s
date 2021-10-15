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
    
1. dockerhub 拉取linux小镜像，alpine  

   ​	`docker pull alpine:latest`

2. dockerhub 拉取 golang编译容器golang:alpine

   ​	`docker pull`

3. 进入容器，拉取代码编译

   

4. 使用Dockerfile将程序和环境打入alpine

5. 运行Dokerfile

6. 将镜像推送到dockerhub
