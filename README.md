# Browser SDK Demo

[![License](https://img.shields.io/github/license/your-org/browser-sdk-demo)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/your-org/browser-sdk-demo)](go.mod)
[![Node.js](https://img.shields.io/badge/node-%3E%3D20.19.0-brightgreen)](https://nodejs.org/)
[![Vue](https://img.shields.io/badge/vue-3.5.27-brightgreen)](https://vuejs.org/)

**中文** | [English](./README.en.md)

## 🌟 项目简介

Browser SDK Demo 是一个完整的浏览器端SDK演示项目，包含后台管理系统、服务端API和桌面客户端应用。该项目展示了现代Web技术栈的最佳实践，适用于SDK集成方、前端开发者和平台运维人员。

## 🏗️ 项目架构

```
browser-sdk-demo/
├── sdk-admin/     # 后台管理系统 (Vue 3 + TypeScript)
├── sdk-server/    # 服务端API (Go + Gin)
├── sdk-client/    # 桌面客户端 (Electron + Vue 3)
```

### 技术栈概览

**前端技术栈:**
- Vue 3 + TypeScript
- Vite 构建工具
- Pinia 状态管理
- Element Plus UI 组件库
- Electron 桌面应用框架

**后端技术栈:**
- Go 语言 (1.25+)
- Gin Web 框架
- GORM ORM
- JWT 认证
- Redis 缓存
- MySQL 数据库

## 🚀 快速开始

### 环境要求

- **Node.js**: >= 20.19.0
- **Go**: >= 1.25
- **pnpm**: 包管理器
- **MySQL**: 数据库
- **Redis**: 缓存服务

### 安装依赖

```bash
# 克隆项目
git clone https://github.com/your-org/browser-sdk-demo.git
cd browser-sdk-demo

# 安装前端依赖
cd sdk-admin && pnpm install
cd ../sdk-client && npm install

# 安装后端依赖
cd ../sdk-server && go mod tidy
```

### 启动服务

**1. 启动后端服务:**
```bash
cd sdk-server
go run main.go start -c resources/config.dev.yaml
```

**2. 启动后台管理:**
```bash
cd sdk-admin
pnpm dev
```

**3. 启动桌面客户端:**
```bash
cd sdk-client
npm run dev
```

## 📁 目录结构

```
browser-sdk-demo/
├── sdk-admin/              # 后台管理系统
│   ├── src/               # 源代码
│   │   ├── components/    # 组件
│   │   ├── layout/        # 布局
│   │   ├── router/        # 路由
│   │   ├── store/         # 状态管理
│   │   └── views/         # 页面视图
│   ├── package.json       # 依赖配置
│   └── vite.config.ts     # 构建配置
│
├── sdk-server/            # 服务端
│   ├── cmd/               # 命令行工具
│   ├── common/            # 公共工具
│   ├── modules/           # 业务模块
│   ├── resources/         # 配置资源
│   ├── go.mod             # Go模块
│   └── main.go            # 入口文件
│
├── sdk-client/            # 客户端
│   ├── src/               # Vue源码
│   ├── electron/          # Electron主进程
│   ├── package.json       # 依赖配置
│   └── vite.config.ts     # 构建配置
│
└── README.md              # 项目说明
```

## 🔧 核心功能

### 后台管理系统 (sdk-admin)
- 用户认证与权限管理
- 浏览器环境监控
- SDK配置管理
- 数据统计分析
- 国际化支持

### 服务端API (sdk-server)
- RESTful API 接口
- JWT Token 认证
- 数据库操作
- 日志记录
- 中间件支持

### 桌面客户端 (sdk-client)
- Electron 桌面应用
- 浏览器环境控制
- 实时状态监控
- 本地数据存储
- 跨平台支持

## ⚙️ 配置说明

### bro-sdk 配置

项目中的 `bro-sdk` 配置用于连接浏览器 SDK 服务：

```yaml
bro-sdk:
  endpoint: http://192.168.0.188:9988 # 请替换为实际值 
  api-key: Ls/e7BzPbtHfcmfR3V+kYJpWHsPffdXKWhEYJdABTmYp+CV/FA1G7EHzk6coamX5 # 请替换为实际值
```

**配置项说明：**
- `endpoint`: 浏览器 SDK 服务的访问地址
- `api-key`: 用于身份验证的 API 密钥

**注意事项：**
- 请根据实际部署环境修改 `endpoint` 地址
- `api-key` 为敏感信息，请妥善保管
- 在生产环境中建议使用 HTTPS 协议

### 其他重要配置

- **数据库配置**: 支持 SQLite 和 MySQL
- **缓存配置**: 支持内存和 Redis 缓存
- **日志配置**: 可配置日志级别和输出格式
- **JWT配置**: 用于用户认证的 Token 配置

## 🛠️ 开发指南

### 代码规范

**前端开发:**
```bash
# 代码检查
pnpm lint

# 代码格式化
pnpm format

# 类型检查
pnpm type-check
```

**后端开发:**
```bash
# 代码格式化
go fmt ./...

# 代码检查
go vet ./...
```

### 构建部署

**构建前端:**
```bash
# 后台管理构建
cd sdk-admin && pnpm build

# 客户端构建
cd sdk-client && npm run build
```

**构建后端:**
```bash
cd sdk-server
go build -ldflags "-w -s" -o bin/server main.go
```

### 测试

```bash
# 前端测试
cd sdk-admin && pnpm test

# 后端测试
cd sdk-server && go test ./...
```

## 📊 API文档

API文档通过 Swagger 自动生成，启动服务后访问:
```
http://localhost:7888/swagger/index.html
```

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

### 提交规范

遵循 [Conventional Commits](https://www.conventionalcommits.org/) 规范:
- `feat`: 新功能
- `fix`: Bug修复
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 📞 联系方式

- 项目主页: [https://github.com/your-org/browser-sdk-demo](https://github.com/your-org/browser-sdk-demo)
- Issues: [https://github.com/your-org/browser-sdk-demo/issues](https://github.com/your-org/browser-sdk-demo/issues)
- 邮箱: your-email@example.com

---

<p align="center">Made with ❤️ by Your Team</p>