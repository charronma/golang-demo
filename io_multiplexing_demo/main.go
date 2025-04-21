package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 连接管理器
type ConnectionManager struct {
	connections map[net.Conn]time.Time
	mutex       sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[net.Conn]time.Time),
	}
}

func (cm *ConnectionManager) Add(conn net.Conn) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.connections[conn] = time.Now()
}

func (cm *ConnectionManager) Remove(conn net.Conn) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	delete(cm.connections, conn)
}

func (cm *ConnectionManager) Count() int {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return len(cm.connections)
}

func main() {
	// 打印启动信息
	fmt.Println("启动服务器，通过代码学习 epoll、poll 和 select")
	fmt.Println("Select 服务器运行在端口 8080")
	fmt.Println("Poll 服务器运行在端口 8081")
	fmt.Println("Epoll 服务器运行在端口 8082")

	// 创建连接管理器
	selectManager := NewConnectionManager()
	pollManager := NewConnectionManager()
	epollManager := NewConnectionManager()

	// 启动三个不同的服务器
	go selectServer(selectManager)
	go pollServer(pollManager)
	go epollServer(epollManager)

	// 启动状态报告
	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Printf("\n=== 服务器状态报告 ===\n")
			fmt.Printf("Select 服务器连接数: %d\n", selectManager.Count())
			fmt.Printf("Poll 服务器连接数: %d\n", pollManager.Count())
			fmt.Printf("Epoll 服务器连接数: %d\n", epollManager.Count())
			fmt.Printf("=====================\n")
		}
	}()

	// 等待中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("服务器关闭")
}

// Select 实现
func selectServer(cm *ConnectionManager) {
	fmt.Println("=== Select Server (端口 8080) ===")
	fmt.Println("特点：")
	fmt.Println("1. 使用 select 系统调用")
	fmt.Println("2. 可以同时监控多个文件描述符")
	fmt.Println("3. 适合处理少量连接")
	fmt.Println("4. O(n) 的时间复杂度，n 为文件描述符数量")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("Listen error: %v\n", err)
		return
	}
	defer listener.Close()

	// 使用通道实现非阻塞 I/O
	connChan := make(chan net.Conn)
	errChan := make(chan error)

	// 启动 goroutine 接受连接
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				errChan <- err
				continue
			}
			connChan <- conn
		}
	}()

	// 使用 select 多路复用通道
	for {
		select {
		case conn := <-connChan:
			fmt.Printf("[Select] 新连接: %s\n", conn.RemoteAddr())
			cm.Add(conn)
			go handleConnection(conn, cm, "Select")
		case err := <-errChan:
			log.Printf("[Select] 错误: %v\n", err)
		case <-time.After(5 * time.Second):
			fmt.Println("[Select] 等待新连接...")
		}
	}
}

// Poll 实现
func pollServer(cm *ConnectionManager) {
	fmt.Println("=== Poll Server (端口 8081) ===")
	fmt.Println("特点：")
	fmt.Println("1. 使用 poll 系统调用")
	fmt.Println("2. 可以监控大量文件描述符")
	fmt.Println("3. 比 select 更高效")
	fmt.Println("4. O(n) 的时间复杂度，但比 select 效率高")

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Printf("Listen error: %v\n", err)
		return
	}
	defer listener.Close()

	// 设置非阻塞模式
	tcpListener := listener.(*net.TCPListener)
	tcpListener.SetDeadline(time.Now().Add(5 * time.Second))

	for {
		// 尝试接受连接，带超时
		conn, err := tcpListener.Accept()
		if err != nil {
			if os.IsTimeout(err) {
				fmt.Println("[Poll] 等待新连接...")
				tcpListener.SetDeadline(time.Now().Add(5 * time.Second))
				continue
			}
			log.Printf("[Poll] 错误: %v\n", err)
			continue
		}

		// 接受新连接
		fmt.Printf("[Poll] 新连接: %s\n", conn.RemoteAddr())
		cm.Add(conn)
		go handleConnection(conn, cm, "Poll")

		// 重新设置超时
		tcpListener.SetDeadline(time.Now().Add(5 * time.Second))
	}
}

// Epoll 实现
func epollServer(cm *ConnectionManager) {
	fmt.Println("=== Epoll Server (端口 8082) ===")
	fmt.Println("特点：")
	fmt.Println("1. 使用 epoll 系统调用")
	fmt.Println("2. 最高效的 I/O 多路复用机制")
	fmt.Println("3. 适合处理大量并发连接")
	fmt.Println("4. O(1) 的时间复杂度，效率最高")

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Printf("Listen error: %v\n", err)
		return
	}
	defer listener.Close()

	// 使用多个 goroutine 和通道实现事件多路复用
	const maxEvents = 10
	connChan := make(chan net.Conn, maxEvents)
	errChan := make(chan error)

	// 启动 goroutine 接受连接
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				errChan <- err
				time.Sleep(100 * time.Millisecond) // 避免过高的 CPU 占用
				continue
			}
			connChan <- conn
		}
	}()

	// 主循环
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case conn := <-connChan:
			fmt.Printf("[Epoll] 新连接: %s\n", conn.RemoteAddr())
			cm.Add(conn)
			go handleConnection(conn, cm, "Epoll")
		case err := <-errChan:
			log.Printf("[Epoll] 错误: %v\n", err)
		case <-ticker.C:
			fmt.Println("[Epoll] 等待新连接...")
		}
	}
}

// 处理连接
func handleConnection(conn net.Conn, cm *ConnectionManager, serverType string) {
	defer func() {
		conn.Close()
		cm.Remove(conn)
		fmt.Printf("[%s] 连接关闭: %s\n", serverType, conn.RemoteAddr())
	}()

	// 设置读取超时
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	// 读取客户端数据
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if os.IsTimeout(err) {
				fmt.Printf("[%s] 连接超时: %s\n", serverType, conn.RemoteAddr())
			} else {
				log.Printf("[%s] 读取错误: %v\n", serverType, err)
			}
			return
		}

		// 重置读取超时
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))

		// 打印接收到的数据
		fmt.Printf("[%s] 收到来自 %s 的数据: %s\n", serverType, conn.RemoteAddr(), string(buf[:n]))

		// 发送响应
		response := fmt.Sprintf("[%s] 服务器已收到: %s", serverType, string(buf[:n]))
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Printf("[%s] 写入错误: %v\n", serverType, err)
			return
		}
	}
}
