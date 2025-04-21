//go:build linux
// +build linux

package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/sys/unix"
)

// Linux 特定的 poll 实现
func pollServerNative(cm *ConnectionManager) {
	fmt.Println("=== Native Poll Server (Linux Only) ===")
	fmt.Println("特点：")
	fmt.Println("1. 使用 Linux 原生的 poll 系统调用")
	fmt.Println("2. 可以监控大量文件描述符")
	fmt.Println("3. 比 select 更高效")
	fmt.Println("4. O(n) 的时间复杂度，但比 select 效率高")

	listener, err := net.Listen("tcp", ":8084")
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

	// 创建 poll 文件描述符
	pollFd := []unix.PollFd{
		{
			Fd:      int32(fd),
			Events:  unix.POLLIN,
			Revents: 0,
		},
	}

	for {
		// 调用 poll
		n, err := unix.Poll(pollFd, 5000) // 5秒超时
		if err != nil {
			log.Printf("Poll error: %v\n", err)
			continue
		}

		if n == 0 {
			fmt.Println("[Native Poll] 等待新连接...")
			continue
		}

		// 检查是否有新的连接
		if pollFd[0].Revents&unix.POLLIN != 0 {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Accept error: %v\n", err)
				continue
			}
			fmt.Printf("[Native Poll] 新连接: %s\n", conn.RemoteAddr())
			cm.Add(conn)
			go handleConnection(conn, cm, "Native Poll")
		}
	}
}
