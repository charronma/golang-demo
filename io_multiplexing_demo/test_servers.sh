#!/bin/bash

# I/O 多路复用服务器测试脚本

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${GREEN}=============================================${NC}"
echo -e "${GREEN}   I/O 多路复用技术演示 - 测试脚本   ${NC}"
echo -e "${GREEN}=============================================${NC}"
echo

# 检查 nc 命令是否可用
if ! command -v nc &>/dev/null; then
    echo -e "${YELLOW}警告: 未找到 nc (netcat) 命令，将使用 telnet 作为备选。${NC}"
    USE_TELNET=true
else
    USE_TELNET=false
fi

# 检查 telnet 命令是否可用
if $USE_TELNET && ! command -v telnet &>/dev/null; then
    echo -e "${RED}错误: 未找到 nc 或 telnet 命令，无法进行测试。${NC}"
    echo -e "${RED}请安装 nc (netcat) 或 telnet 后再试。${NC}"
    exit 1
fi

# 检测系统类型
if [[ "$(uname)" == "Linux" ]]; then
    echo -e "${BLUE}检测到 Linux 系统，可以使用原生函数实现的服务器。${NC}"
    NATIVE_SUPPORT=true
else
    echo -e "${YELLOW}检测到非 Linux 系统，将通过 Docker 容器访问所有服务器。${NC}"
    NATIVE_SUPPORT=false
fi

# 显示菜单
function show_menu {
    echo
    echo -e "${GREEN}请选择要测试的服务器:${NC}"
    echo "1) Select 服务器 (端口 8080)"
    echo "2) Poll 服务器 (端口 8081)"
    echo "3) Epoll 服务器 (端口 8082)"
    echo "4) 原生 Select 服务器 (端口 8083)"
    echo "5) 原生 Poll 服务器 (端口 8084)"
    echo "6) 原生 Epoll 服务器 (端口 8085)"
    echo "0) 退出"
    echo
}

# 测试服务器函数
function test_server {
    local port=$1
    local name=$2

    echo -e "\n${GREEN}测试 $name 服务器 (端口 $port)...${NC}"
    echo -e "${YELLOW}连接到服务器，输入消息并查看响应。按 Ctrl+C 退出。${NC}"
    echo -e "${BLUE}---------- 连接开始 ----------${NC}"

    if $USE_TELNET; then
        telnet localhost $port
    else
        nc localhost $port
    fi

    echo -e "${BLUE}---------- 连接结束 ----------${NC}"
}

# 主循环
while true; do
    show_menu
    read -p "请输入选项: " CHOICE

    case $CHOICE in
    0)
        echo -e "\n${YELLOW}退出测试。${NC}"
        exit 0
        ;;
    1)
        test_server 8080 "Select"
        ;;
    2)
        test_server 8081 "Poll"
        ;;
    3)
        test_server 8082 "Epoll"
        ;;
    4)
        test_server 8083 "原生 Select"
        ;;
    5)
        test_server 8084 "原生 Poll"
        ;;
    6)
        test_server 8085 "原生 Epoll"
        ;;
    *)
        echo -e "\n${RED}无效选项，请重新选择。${NC}"
        ;;
    esac

    # 等待用户按任意键继续
    echo
    read -p "按回车键继续..."
done
