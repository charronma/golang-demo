# I/O 多路复用技术演示 - 快速入门指南

这个演示项目展示了三种主要的 I/O 多路复用技术（select、poll 和 epoll）的实现和使用方法。

## 快速开始

### 1. 使用脚本启动

最简单的方法是使用提供的脚本：

```bash
# 运行演示服务器
./run_demo.sh

# **测试服务器连接**
./test_servers.sh
```

### 2. 手动启动

如果你想更详细了解启动过程，可以手动执行以下步骤：

#### Docker 方式（推荐）

```bash
# 构建 Docker 镜像
docker-compose build

# 启动服务
docker-compose up -d

# 查看日志
docker-compose logs -f
```

#### 本地方式

```bash
# 下载依赖
go mod download

# 编译
go build -o server

# 运行
./server
```

## 测试连接

启动服务器后，可以使用 telnet 或 nc 连接到不同端口进行测试：

```bash
# 测试 select 服务器
telnet localhost 8080

# 测试 poll 服务器
telnet localhost 8081

# 测试 epoll 服务器
telnet localhost 8082
```

连接后，输入任何文本消息，服务器会返回相应的响应。

## 注意事项

1. 原生的 select、poll 和 epoll 实现（端口 8083-8085）仅在 Linux 系统上可用。

2. 在非 Linux 系统上（如 macOS 或 Windows），只能使用跨平台实现（端口 8080-8082）。

3. 使用 Docker 运行时，所有实现（包括 Linux 原生实现）都可以在任何平台上运行。

## 各服务器端口列表

| 服务器 | 端口 | 实现方式 | 平台支持 |
|-------|-----|--------|---------|
| Select | 8080 | Go 标准库 | 全平台 |
| Poll | 8081 | Go 标准库 | 全平台 |
| Epoll | 8082 | Go 标准库 | 全平台 |
| 原生 Select | 8083 | Linux 系统调用 | 仅 Linux |
| 原生 Poll | 8084 | Linux 系统调用 | 仅 Linux |
| 原生 Epoll | 8085 | Linux 系统调用 | 仅 Linux | 