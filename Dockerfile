# 多阶段构建：第一阶段构建Go后端
FROM golang:1.22-alpine as go-builder

# 设置工作目录
WORKDIR /app

# 拷贝Go模块文件
COPY sdk-server/go.mod sdk-server/go.sum ./

# 下载依赖
RUN go mod download

# 拷贝Go源码
COPY sdk-server/. .

# 构建Go应用
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go build -o bin/server main.go

# 多阶段构建：第二阶段构建Vue前端
FROM node:20-alpine as vue-builder

# 设置工作目录
WORKDIR /app

# 安装pnpm
RUN npm install -g pnpm

# 拷贝前端依赖文件
COPY sdk-admin/package.json sdk-admin/pnpm-lock.yaml ./

# 安装前端依赖
RUN pnpm install --frozen-lockfile

# 拷贝前端源码
COPY sdk-admin/. .

# 构建前端应用
RUN pnpm build

# 多阶段构建：第三阶段创建最终镜像
FROM alpine:latest

# 安装nginx
RUN apk add --no-cache nginx

# 修改时区
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# 创建应用目录
RUN mkdir -p /app/server/bin \
    && mkdir -p /app/admin/dist \
    && mkdir -p /var/log/nginx

# 从Go构建阶段复制二进制文件
COPY --from=go-builder /app/bin/server /app/server/bin/server
COPY --from=go-builder /app/resources /app/server/resources/
COPY --from=go-builder /app/docs /app/server/docs/

# 从Vue构建阶段复制静态文件
COPY --from=vue-builder /app/dist /app/admin/dist

# 复制nginx配置
COPY nginx.conf /etc/nginx/nginx.conf

# 复制启动脚本
COPY start.sh /start.sh
RUN chmod +x /start.sh

# 暴露端口
EXPOSE 80 7888

# 设置工作目录
WORKDIR /app

# 启动脚本
CMD ["/start.sh"]