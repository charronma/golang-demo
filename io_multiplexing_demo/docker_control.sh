#!/bin/bash

# Docker容器控制脚本 - IO多路复用演示项目

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}=============================================${NC}"
echo -e "${CYAN}      I/O 多路复用演示 - Docker 控制      ${NC}"
echo -e "${CYAN}=============================================${NC}"
echo

# 检查 Docker 和 Docker Compose 是否可用
if ! command -v docker &> /dev/null; then
    echo -e "${RED}错误: Docker 未安装或不在PATH中。${NC}"
    exit 1
fi

if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
    echo -e "${RED}错误: Docker Compose 未安装或不在PATH中。${NC}"
    exit 1
fi

# 判断使用docker-compose还是docker compose命令
if command -v docker-compose &> /dev/null; then
    DOCKER_COMPOSE="docker-compose"
else
    DOCKER_COMPOSE="docker compose"
fi

# 显示状态
function show_status {
    echo -e "\n${GREEN}==== 容器状态 ====${NC}"
    $DOCKER_COMPOSE ps
    echo
}

# 启动容器
function start_container {
    echo -e "\n${GREEN}启动容器...${NC}"
    $DOCKER_COMPOSE up -d
    echo -e "${GREEN}容器已启动!${NC}"
    show_status
}

# 停止容器
function stop_container {
    echo -e "\n${YELLOW}停止容器...${NC}"
    $DOCKER_COMPOSE down
    echo -e "${YELLOW}容器已停止!${NC}"
}

# 重建并启动容器
function rebuild_container {
    echo -e "\n${BLUE}重新构建镜像...${NC}"
    $DOCKER_COMPOSE down
    $DOCKER_COMPOSE build --no-cache
    $DOCKER_COMPOSE up -d
    echo -e "${GREEN}容器已重建并启动!${NC}"
    show_status
}

# 显示日志
function show_logs {
    echo -e "\n${PURPLE}显示容器日志 (按Ctrl+C退出)...${NC}"
    $DOCKER_COMPOSE logs -f
}

# 访问容器Shell
function access_shell {
    echo -e "\n${BLUE}进入容器Shell...${NC}"
    $DOCKER_COMPOSE exec linux-env /bin/bash
}

# 运行测试脚本
function run_test {
    echo -e "\n${PURPLE}运行测试脚本...${NC}"
    if [ -f ./test_servers.sh ]; then
        chmod +x ./test_servers.sh
        ./test_servers.sh
    else
        echo -e "${RED}错误: 未找到测试脚本 test_servers.sh${NC}"
    fi
}

# 主菜单
function show_menu {
    echo -e "${GREEN}请选择操作:${NC}"
    echo "1) 启动容器"
    echo "2) 停止容器"
    echo "3) 重新构建并启动容器"
    echo "4) 显示容器状态"
    echo "5) 查看容器日志"
    echo "6) 进入容器Shell"
    echo "7) 运行测试脚本"
    echo "0) 退出"
    echo
}

# 主循环
while true; do
    show_menu
    read -p "请输入选项 [0-7]: " CHOICE
    
    case $CHOICE in
        0)
            echo -e "\n${YELLOW}退出脚本。${NC}"
            exit 0
            ;;
        1)
            start_container
            ;;
        2)
            stop_container
            ;;
        3)
            rebuild_container
            ;;
        4)
            show_status
            ;;
        5)
            show_logs
            ;;
        6)
            access_shell
            ;;
        7)
            run_test
            ;;
        *)
            echo -e "\n${RED}无效选项，请重新选择。${NC}"
            ;;
    esac
    
    # 如果不是日志或Shell选项，等待用户按任意键继续
    if [[ $CHOICE != 5 && $CHOICE != 6 ]]; then
        echo
        read -p "按回车键继续..."
    fi
done 