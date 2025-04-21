FROM golang:1.21

WORKDIR /usr/src/app

# 设置时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 安装必要的工具
RUN apt-get update && apt-get install -y \
    vim \
    net-tools \
    telnet \
    && rm -rf /var/lib/apt/lists/*

# 设置环境变量
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

# 预复制 go.mod 和 go.sum 以缓存依赖
COPY go.mod ./
RUN go mod download

# 复制源代码
COPY . .

# 暴露端口
EXPOSE 8080 8081 8082 8083 8084 8085

# 设置启动命令 - 先构建再运行
CMD go build -o server && ./server
