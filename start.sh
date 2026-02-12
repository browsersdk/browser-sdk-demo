#!/bin/sh

# 设置环境变量
export SERVER_CONFIG="/app/server/resources/config.docker.yaml"

# 检查配置文件是否存在
if [ ! -f "$SERVER_CONFIG" ]; then
    echo "Error: Config file $SERVER_CONFIG not found!"
    exit 1
fi

echo "Starting Browser SDK Demo..."

# 创建日志目录
mkdir -p /app/server/logs

# 启动Go后端服务（后台运行）
echo "Starting backend server..."
/app/server/bin/server start -c $SERVER_CONFIG &
BACKEND_PID=$!

# 等待后端服务启动
echo "Waiting for backend to start..."
sleep 5

# 检查后端服务是否正常运行
if ! kill -0 $BACKEND_PID 2>/dev/null; then
    echo "Error: Backend server failed to start!"
    exit 1
fi

echo "Backend server started with PID: $BACKEND_PID"

# 启动Nginx
echo "Starting Nginx..."
nginx -g "daemon off;" &

# 等待信号处理
trap "echo 'Received SIGTERM, shutting down...'; kill $BACKEND_PID; exit 0" SIGTERM
trap "echo 'Received SIGINT, shutting down...'; kill $BACKEND_PID; exit 0" SIGINT

# 等待所有后台进程
wait