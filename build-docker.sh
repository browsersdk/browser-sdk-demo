#!/bin/bash

# Browser SDK Demo Docker构建脚本

set -e

echo "🚀 Browser SDK Demo Docker Build Script"
echo "======================================"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 检查必要工具
check_prerequisites() {
    echo "🔍 Checking prerequisites..."
    
    if ! command -v docker &> /dev/null; then
        echo -e "${RED}Error: Docker is not installed${NC}"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        echo -e "${YELLOW}Warning: docker-compose is not installed, using docker compose instead${NC}"
    fi
    
    echo -e "${GREEN}✓ Prerequisites check passed${NC}"
}

# 构建镜像
build_image() {
    local tag=${1:-"browser-sdk-demo:latest"}
    
    echo "🏗️ Building Docker image: $tag"
    
    # 清理之前的构建缓存（可选）
    # docker builder prune -f
    
    docker build -t "$tag" .
    
    echo -e "${GREEN}✓ Image built successfully${NC}"
}

# 启动服务
start_services() {
    echo "🚢 Starting services..."
    
    if command -v docker-compose &> /dev/null; then
        docker-compose up -d
    else
        docker compose up -d
    fi
    
    echo -e "${GREEN}✓ Services started${NC}"
}

# 停止服务
stop_services() {
    echo "🛑 Stopping services..."
    
    if command -v docker-compose &> /dev/null; then
        docker-compose down
    else
        docker compose down
    fi
    
    echo -e "${GREEN}✓ Services stopped${NC}"
}

# 查看日志
show_logs() {
    echo "📋 Showing service logs..."
    
    if command -v docker-compose &> /dev/null; then
        docker-compose logs -f
    else
        docker compose logs -f
    fi
}

# 显示帮助
show_help() {
    echo "Usage: $0 [COMMAND]"
    echo ""
    echo "Commands:"
    echo "  build     Build the Docker image"
    echo "  start     Start services"
    echo "  stop      Stop services"
    echo "  logs      Show service logs"
    echo "  restart   Restart services"
    echo "  help      Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 build                    # Build with default tag"
    echo "  $0 build my-tag             # Build with custom tag"
    echo "  $0 start                    # Start all services"
    echo "  $0 stop                     # Stop all services"
}

# 主逻辑
main() {
    check_prerequisites
    
    case "${1:-help}" in
        build)
            build_image "${2:-browser-sdk-demo:latest}"
            ;;
        start)
            start_services
            ;;
        stop)
            stop_services
            ;;
        restart)
            stop_services
            start_services
            ;;
        logs)
            show_logs
            ;;
        help|--help|-h)
            show_help
            ;;
        *)
            echo -e "${RED}Unknown command: $1${NC}"
            show_help
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"