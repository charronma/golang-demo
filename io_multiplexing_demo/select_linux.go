//go:build linux
// +build linux

package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/sys/unix"
)

// Linux 特定的 select 实现
func selectServerNative(cm *ConnectionManager) {
	fmt.Println("=== Native Select Server (Linux Only) ===")
	fmt.Println("特点：")
	fmt.Println("1. 使用 Linux 原生的 select 系统调用")
	fmt.Println("2. 可以同时监控多个文件描述符")
	fmt.Println("3. 适合处理少量连接")
	fmt.Println("4. O(n) 的时间复杂度，n 为文件描述符数量")

	listener, err := net.Listen("tcp", ":8083")
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

	// 创建文件描述符集合
	readFds := &unix.FdSet{}
	FD_SET(fd, readFds)

	// 设置超时时间
	timeout := &unix.Timeval{
		Sec:  5,
		Usec: 0,
	}

	for {
		// 复制文件描述符集合
		readFdsCopy := *readFds

		// 调用 select
		n, err := unix.Select(fd+1, &readFdsCopy, nil, nil, timeout)
		if err != nil {
			log.Printf("Select error: %v\n", err)
			continue
		}

		if n == 0 {
			fmt.Println("[Native Select] 等待新连接...")
			continue
		}

		// 检查是否有新的连接
		if FD_ISSET(fd, &readFdsCopy) {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Accept error: %v\n", err)
				continue
			}
			fmt.Printf("[Native Select] 新连接: %s\n", conn.RemoteAddr())
			cm.Add(conn)
			go handleConnection(conn, cm, "Native Select")
		}
	}
}

// FD_SET 宏的实现
func FD_SET(fd int, set *unix.FdSet) {
	set.Bits[fd/64] |= 1 << (uint(fd) % 64)
}

// FD_ISSET 宏的实现
func FD_ISSET(fd int, set *unix.FdSet) bool {
	return set.Bits[fd/64]&(1<<(uint(fd)%64)) != 0
}
