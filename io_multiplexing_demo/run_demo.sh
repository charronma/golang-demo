#!/bin/bash

# I/O 多路复用演示运行脚本

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${GREEN}=============================================${NC}"
echo -e "${GREEN}   I/O 多路复用技术演示 - 启动脚本   ${NC}"
echo -e "${GREEN}=============================================${NC}"
echo

# 检测系统类型
if [[ "$(uname)" == "Linux" ]]; then
    echo -e "${BLUE}检测到 Linux 系统，将启动全部 6 个服务器...${NC}"
    NATIVE_SUPPORT=true
else
    echo -e "${YELLOW}检测到非 Linux 系统，将只启动跨平台服务器...${NC}"
    NATIVE_SUPPORT=false
fi

# 选择运行模式
echo
echo -e "${GREEN}请选择运行模式:${NC}"
echo "1) 使用 Docker 运行 (推荐，可在任何平台运行全部服务器)"
echo "2) 直接在本机运行 (依赖本地 Go 环境)"
echo "3) 退出"
echo

read -p "请输入选项 [1-3]: " CHOICE

case $CHOICE in
    1)
        echo -e "\n${GREEN}使用 Docker 运行演示...${NC}"
        
        # 检查 Docker 是否可用
        if ! command -v docker &> /dev/null || ! command -v docker-compose &> /dev/null; then
            echo -e "${RED}错误: 未找到 Docker 或 Docker Compose，请先安装。${NC}"
            exit 1
        fi
        
        echo -e "${YELLOW}构建 Docker 镜像...${NC}"
        docker-compose build
        
        echo -e "${YELLOW}启动容器...${NC}"
        docker-compose up -d
        
        echo -e "${GREEN}服务器已启动!${NC}"
        echo -e "${YELLOW}查看日志...${NC}"
        docker-compose logs -f
        ;;
        
    2)
        echo -e "\n${GREEN}直接在本机运行演示...${NC}"
        
        # 检查 Go 是否可用
        if ! command -v go &> /dev/null; then
            echo -e "${RED}错误: 未找到 Go 环境，请先安装。${NC}"
            exit 1
        fi
        
        # 下载依赖
        echo -e "${YELLOW}下载依赖...${NC}"
        go mod download
        
        # 编译
        echo -e "${YELLOW}编译程序...${NC}"
        go build -o server
        
        echo -e "${GREEN}服务器已编译，开始运行...${NC}"
        ./server
        ;;
        
    3)
        echo -e "\n${YELLOW}退出演示。${NC}"
        exit 0
        ;;
        
    *)
        echo -e "\n${RED}无效选项，请重新运行脚本并选择 1-3。${NC}"
        exit 1
        ;;
esac 