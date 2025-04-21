# I/O 多路复用实现示例

这个项目演示了三种常见的 I/O 多路复用技术的实现：select、poll 和 epoll。

## 项目结构

- `main.go`: 跨平台实现，使用 Go 语言标准库模拟 I/O 多路复用
- `select_linux.go`: Linux 平台特定的 select 系统调用实现
- `poll_linux.go`: Linux 平台特定的 poll 系统调用实现
- `epoll_linux.go`: Linux 平台特定的 epoll 系统调用实现
- `main_linux.go`: Linux 平台特定的初始化代码
- `go.mod`: Go 模块定义文件
- `Dockerfile`: Docker 镜像构建文件
- `docker-compose.yml`: Docker Compose 配置文件

## 功能特点

1. **Select 服务器 (端口 8080)**
   - 使用 select 模式
   - 可以同时监控多个文件描述符
   - 适合处理少量连接
   - O(n) 的时间复杂度

2. **Poll 服务器 (端口 8081)**
   - 使用 poll 模式
   - 可以监控大量文件描述符
   - 比 select 更高效
   - O(n) 的时间复杂度，但比 select 效率高

3. **Epoll 服务器 (端口 8082)**
   - 使用 epoll 模式
   - 最高效的 I/O 多路复用机制
   - 适合处理大量并发连接
   - O(1) 的时间复杂度，效率最高

## Linux 特定实现

在 Linux 平台上，还会启动以下服务器，使用原生系统调用：

- **原生 Select 服务器 (端口 8083)**
- **原生 Poll 服务器 (端口 8084)**
- **原生 Epoll 服务器 (端口 8085)**

## 编译和运行

### 在 Docker 中运行（推荐）

```bash
# 进入项目目录
cd io_multiplexing_demo

# 构建 Docker 镜像
docker-compose build

# 运行容器
docker-compose up -d

# 查看日志
docker-compose logs -f
```

### 直接编译运行

```bash
# 进入项目目录
cd io_multiplexing_demo

# 安装依赖
go mod download

# 编译
go build -o server

# 运行
./server
```

## 测试方法

可以使用 telnet 或 nc 工具连接到不同的服务器端口进行测试：

```bash
# 连接到 select 服务器
telnet localhost 8080

# 连接到 poll 服务器
telnet localhost 8081

# 连接到 epoll 服务器
telnet localhost 8082
```

在 Linux 系统上，还可以连接到原生实现的服务器：

```bash
# 连接到原生 select 服务器
telnet localhost 8083

# 连接到原生 poll 服务器
telnet localhost 8084

# 连接到原生 epoll 服务器
telnet localhost 8085
```

## 性能比较

三种 I/O 多路复用技术的主要区别：

1. **Select**
   - 同时监控文件描述符的数量有限制（一般为1024）
   - 每次调用都需要将文件描述符集合从用户空间拷贝到内核空间
   - O(n) 的时间复杂度，n为文件描述符数量

2. **Poll**
   - 没有文件描述符数量的限制
   - 依然需要将文件描述符集合从用户空间拷贝到内核空间
   - O(n) 的时间复杂度，但比 select 效率高

3. **Epoll**
   - 没有文件描述符数量的限制
   - 使用事件通知机制，避免了大量的文件描述符拷贝
   - 只关注活跃的文件描述符，不需要遍历所有文件描述符
   - O(1) 的时间复杂度，效率最高

## 环境要求

- Go 1.21 或更高版本
- 如需运行原生实现，需要在 Linux 系统上运行
- Docker 和 Docker Compose (如果使用 Docker 运行)

## 项目源码结构详解

1. **main.go**：
   - 定义了连接管理器（ConnectionManager）
   - 实现了跨平台版本的 select、poll 和 epoll 服务器
   - 处理客户端连接和数据收发

2. **select_linux.go**：
   - 使用 Linux 系统调用实现 select 多路复用
   - 定义了 FD_SET 和 FD_ISSET 等底层操作

3. **poll_linux.go**：
   - 使用 Linux 系统调用实现 poll 多路复用
   - 直接操作底层文件描述符

4. **epoll_linux.go**：
   - 使用 Linux 系统调用实现 epoll 多路复用
   - 演示了事件驱动的 I/O 处理模型

5. **main_linux.go**：
   - 初始化 Linux 特定的服务器
   - 只在 Linux 平台上编译和运行 