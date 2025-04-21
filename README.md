# Go 语言编程示例集合

这个仓库包含了 Go 语言各种特性和应用场景的示例代码集合，旨在帮助开发者快速掌握 Go 语言的核心概念和实践技巧。

## 示例集合

### 已实现示例

[io_multiplexing_demo](./io_multiplexing_demo) - I/O 多路复用技术（select、poll 和 epoll）的实现和比较。

### 计划添加的示例

#### 网络编程
- **TCP/UDP 服务器** - 基础网络编程示例
- **HTTP 服务器** - RESTful API 与中间件实现
- **WebSocket** - 实时通信示例
- **gRPC** - 高性能 RPC 框架使用

#### 系统编程
- **内存管理** - Go 内存模型与垃圾回收机制探索
- **并发模式** - goroutine 与 channel 最佳实践
- **系统信号处理** - 优雅启动与关闭
- **文件操作** - 高效文件读写与处理

#### Go 语言特性
- **反射机制** - 运行时类型检查与动态调用
- **接口实现** - 面向接口编程的最佳实践
- **泛型编程** - Go 1.18+ 泛型特性示例
- **错误处理** - 错误传播与包装模式

#### 工具与模式
- **上下文管理** - context 包的使用场景
- **依赖注入** - 松耦合设计与测试
- **单元测试** - 测试驱动开发示例
- **性能优化** - 性能分析与基准测试

## 如何使用

每个示例都在独立的目录中，包含自己的 README 文件、源代码和运行脚本。要运行特定示例，请进入相应目录并按照其 README 指导操作。

例如，运行 I/O 多路复用演示：

```bash
# 进入多路复用演示目录
cd io_multiplexing_demo

# 使用脚本启动
./run_demo.sh

# 测试服务器连接
./test_servers.sh
```

## 环境要求

- Go 1.21 或更高版本
- Docker 和 Docker Compose (用于容器化运行)
- 特定示例可能有额外的依赖，详见各示例的 README

## 学习资源

这些示例旨在作为学习工具，建议配合以下资源一起使用：

1. [Go 官方文档](https://golang.org/doc/)
2. [Go 标准库参考](https://pkg.go.dev/std)
3. [Go by Example](https://gobyexample.com/)
4. [Go 语言高级编程](https://github.com/chai2010/advanced-go-programming-book)
5. [Go 系统编程](https://github.com/astaxie/go-system-programming)
6. [Linux 系统编程手册](http://man7.org/linux/man-pages/index.html)

## 贡献

欢迎通过以下方式贡献：

- 添加新的示例程序
- 改进现有示例的代码质量
- 完善文档和注释
- 修复 bug 和问题

请提交 Pull Request 或 Issue 来参与项目改进。

## 许可

MIT License 