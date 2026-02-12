# Browser SDK Demo Docker部署指南

## 🐳 概述

本指南介绍了如何使用Docker一键部署Browser SDK Demo，包含后端API服务(sdk-server)和前端管理界面(sdk-admin)。

## 🏗️ 架构说明

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   用户浏览器     │───▶│     Nginx       │───▶│   Go后端服务     │
│  (端口 80)      │    │  (反向代理)     │    │   (端口 7888)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │
                              ▼
                       ┌─────────────────┐
                       │  Vue前端静态文件  │
                       │    (Admin界面)   │
                       └─────────────────┘
```

## 🚀 快速开始

### 1. 环境准备

确保已安装：
- Docker 20.10+
- Docker Compose 1.29+

```bash
# 检查Docker版本
docker --version
docker-compose --version
```

### 2. 构建和启动

```bash
# 克隆项目
git clone https://github.com/your-org/browser-sdk-demo.git
cd browser-sdk-demo

# 给脚本执行权限
chmod +x build-docker.sh

# 一键构建和启动
./build-docker.sh build
./build-docker.sh start
```

或者使用传统方式：
```bash
# 构建镜像
docker build -t browser-sdk-demo .

# 启动服务
docker-compose up -d
```

### 3. 访问应用

服务启动后，可以通过以下地址访问：

- **前端管理界面**: http://localhost
- **API文档**: http://localhost/swagger/index.html
- **健康检查**: http://localhost/health
- **后端API直连**: http://localhost:7888

## 📁 项目结构

```
browser-sdk-demo/
├── Dockerfile              # 多阶段构建Dockerfile
├── docker-compose.yml      # Docker Compose配置
├── nginx.conf             # Nginx配置文件
├── start.sh               # 容器启动脚本
├── build-docker.sh        # 构建部署脚本
├── DOCKER_DEPLOY.md       # 本部署文档
├── sdk-server/            # Go后端服务
│   └── resources/
│       └── config.docker.yaml  # Docker环境配置
└── sdk-admin/             # Vue前端管理界面
```

## ⚙️ 配置说明

### 环境变量

在 `docker-compose.yml` 中可以配置：

```yaml
environment:
  - TZ=Asia/Shanghai  # 时区设置
```

### 端口映射

默认端口配置：
- `80:80` - 前端和API统一入口
- `7888:7888` - 后端API直连端口（可选）

### 数据库配置

默认使用本地配置，如需外部数据库，请修改 `config.docker.yaml`：

```yaml
database:
  host: "mysql"  # docker-compose服务名
  port: 3306
  database: "browser_sdk"
  username: "browser"
  password: "browser123"

redis:
  addr: "redis:6379"  # docker-compose服务名
```

## 🛠️ 管理命令

### 使用构建脚本（推荐）

```bash
# 构建镜像
./build-docker.sh build [tag]

# 启动服务
./build-docker.sh start

# 停止服务
./build-docker.sh stop

# 重启服务
./build-docker.sh restart

# 查看日志
./build-docker.sh logs

# 显示帮助
./build-docker.sh help
```

### 使用Docker Compose

```bash
# 构建镜像
docker-compose build

# 启动服务
docker-compose up -d

# 停止服务
docker-compose down

# 查看状态
docker-compose ps

# 查看日志
docker-compose logs -f

# 重启服务
docker-compose restart
```

### 直接使用Docker

```bash
# 构建镜像
docker build -t browser-sdk-demo .

# 运行容器
docker run -d -p 80:80 -p 7888:7888 --name browser-sdk-demo browser-sdk-demo

# 查看日志
docker logs -f browser-sdk-demo

# 停止容器
docker stop browser-sdk-demo

# 删除容器
docker rm browser-sdk-demo
```

## 🔧 高级配置

### 挂载卷配置

```yaml
volumes:
  # 挂载配置文件
  - ./sdk-server/resources/config.docker.yaml:/app/server/resources/config.docker.yaml
  # 挂载日志目录
  - ./logs:/app/server/logs
  # 挂载静态文件（开发时）
  - ./sdk-admin/dist:/app/admin/dist
```

### 健康检查

容器内置健康检查：
```bash
curl http://localhost/health
```

### 性能调优

Nginx配置已优化：
- Gzip压缩静态资源
- 静态文件缓存策略
- 连接超时设置
- WebSocket支持

## 🔒 安全建议

1. **生产环境配置**：
   - 修改JWT密钥
   - 使用HTTPS
   - 配置防火墙规则
   - 定期更新基础镜像

2. **敏感信息处理**：
   ```bash
   # 使用Docker secrets
   echo "your-secret-key" | docker secret create jwt_secret -
   ```

3. **网络隔离**：
   ```yaml
   networks:
     frontend:
     backend:
   
   services:
     browser-sdk-demo:
       networks:
         - frontend
         - backend
   ```

## 📊 监控和日志

### 查看容器状态

```bash
# 查看运行状态
docker-compose ps

# 查看资源使用
docker stats browser-sdk-demo

# 查看详细信息
docker inspect browser-sdk-demo
```

### 日志管理

```bash
# 实时日志
docker-compose logs -f

# 指定服务日志
docker-compose logs -f browser-sdk-demo

# 日志轮转配置在nginx.conf中
```

## 🐛 故障排除

### 常见问题

1. **端口占用**：
   ```bash
   # 检查端口占用
   netstat -tlnp | grep :80
   lsof -i :80
   
   # 修改docker-compose.yml端口映射
   ports:
     - "8080:80"  # 改为8080端口
   ```

2. **构建失败**：
   ```bash
   # 清理构建缓存
   docker builder prune -a
   
   # 重新构建
   docker-compose build --no-cache
   ```

3. **服务无法启动**：
   ```bash
   # 查看详细日志
   docker-compose logs browser-sdk-demo
   
   # 检查配置文件
   docker exec browser-sdk-demo cat /app/server/resources/config.docker.yaml
   ```

4. **数据库连接失败**：
   ```bash
   # 检查网络连通性
   docker-compose exec browser-sdk-demo ping mysql
   
   # 检查数据库服务状态
   docker-compose logs mysql
   ```

### 调试模式

```bash
# 交互式进入容器
docker-compose exec browser-sdk-demo sh

# 查看进程
ps aux

# 测试网络连接
curl -v http://localhost:7888/health
```

## 📈 性能基准

### 资源需求

**最小配置**：
- CPU: 1核
- 内存: 512MB
- 磁盘: 1GB

**推荐配置**：
- CPU: 2核
- 内存: 1GB
- 磁盘: 2GB

### 性能指标

- 并发用户数: 100-500
- 响应时间: < 200ms
- 吞吐量: 1000+ QPS

## 🔄 升级和维护

### 版本升级

```bash
# 拉取最新代码
git pull origin main

# 重新构建
./build-docker.sh build

# 平滑升级
docker-compose up -d --no-deps --build browser-sdk-demo
```

### 备份策略

```bash
# 备份配置
docker-compose exec browser-sdk-demo tar czf /backup/config.tar.gz /app/server/resources/

# 备份日志
docker cp browser-sdk-demo:/app/server/logs ./backup/logs-$(date +%Y%m%d)
```

## 🤝 贡献和支持

如有问题或建议，请：
1. 查看[Issues](https://github.com/your-org/browser-sdk-demo/issues)
2. 提交[Pull Request](https://github.com/your-org/browser-sdk-demo/pulls)
3. 联系维护团队

---
**注意**: 本文档适用于Docker部署场景，其他部署方式请参考主README文档。