//go:build linux
// +build linux

package main

import (
	"fmt"
)

func init() {
	// 注册Linux特定的服务器启动
	fmt.Println("检测到Linux系统，将启动原生的 select/poll/epoll 实现")
	fmt.Println("原生 Select 服务器运行在端口 8083")
	fmt.Println("原生 Poll 服务器运行在端口 8084")
	fmt.Println("原生 Epoll 服务器运行在端口 8085")

	// 创建连接管理器
	selectNativeManager := NewConnectionManager()
	pollNativeManager := NewConnectionManager()
	epollNativeManager := NewConnectionManager()

	// 启动原生的Linux实现
	go selectServerNative(selectNativeManager)
	go pollServerNative(pollNativeManager)
	go epollServerNative(epollNativeManager)
}
