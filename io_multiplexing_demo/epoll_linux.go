//go:build linux
// +build linux

package main

import (
	"fmt"
	"log"
	"net"
	"syscall"

	"golang.org/x/sys/unix"
)

// Linux 特定的 epoll 实现
func epollServerNative(cm *ConnectionManager) {
	fmt.Println("=== Native Epoll Server (Linux Only) ===")
	fmt.Println("特点：")
	fmt.Println("1. 使用 Linux 原生的 epoll 系统调用")
	fmt.Println("2. 最高效的 I/O 多路复用机制")
	fmt.Println("3. 适合处理大量并发连接")
	fmt.Println("4. O(1) 的时间复杂度，效率最高")

	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Printf("Listen error: %v\n", err)
		return
	}
	defer listener.Close()

	// 获取文件描述符
	file, err := listener.(*net.TCPListener).File()
	if err != nil {
		log.Printf("File error: %v\n", err)
		return
	}
	fd := int(file.Fd())

	// 创建 epoll 实例
	epollFd, err := unix.EpollCreate1(0)
	if err != nil {
		log.Printf("EpollCreate1 error: %v\n", err)
		return
	}
	defer unix.Close(epollFd)

	// 添加监听套接字到 epoll
	var event unix.EpollEvent
	event.Events = unix.EPOLLIN
	event.Fd = int32(fd)

	if err := unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, fd, &event); err != nil {
		log.Printf("EpollCtl error: %v\n", err)
		return
	}

	// 创建事件数组
	events := make([]unix.EpollEvent, 10)

	for {
		// 调用 epoll_wait
		n, err := unix.EpollWait(epollFd, events, 5000) // 5秒超时
		if err != nil {
			if err == syscall.EINTR {
				continue // 被信号中断，继续
			}
			log.Printf("EpollWait error: %v\n", err)
			continue
		}

		if n == 0 {
			fmt.Println("[Native Epoll] 等待新连接...")
			continue
		}

		// 处理事件
		for i := 0; i < n; i++ {
			if int(events[i].Fd) == fd {
				conn, err := listener.Accept()
				if err != nil {
					log.Printf("Accept error: %v\n", err)
					continue
				}
				fmt.Printf("[Native Epoll] 新连接: %s\n", conn.RemoteAddr())
				cm.Add(conn)
				go handleConnection(conn, cm, "Native Epoll")
			}
		}
	}
}
