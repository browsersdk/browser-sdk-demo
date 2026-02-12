# 浏览器SDK客户端

这是一个基于Vue3 + TypeScript + Electron的浏览器环境管理客户端应用程序。

## 功能特性

- ✅ 用户登录认证（集成服务器API）
- ✅ 用户信息管理（获取、修改用户信息）
- ✅ 浏览器环境管理（列表展示、新建、修改、删除环境配置）
- ✅ 响应式界面设计
- ✅ 多标签页导航
- ✅ 模态框交互
- ✅ 服务器API集成
- ✅ 科技风格UI设计

## 技术栈

- **前端框架**: Vue 3 + TypeScript
- **状态管理**: Pinia
- **路由**: Vue Router
- **构建工具**: Vite
- **桌面应用**: Electron
- **UI设计**: 自定义科技风格

## 快速开始

### 环境要求

- Node.js >= 16.0.0
- npm 或 yarn

### 安装依赖

```bash
npm install
```

### 开发模式

#### 同时启动Web和Electron（推荐）
```bash
npm run dev
```

#### 仅启动Web开发服务器
```bash
npm run dev:web
```

#### 仅启动Electron（需要先运行Web服务器）
```bash
npm run dev:electron
```

### 构建生产版本

#### 构建Web版本
```bash
npm run build
```

#### 构建Electron应用
```bash
npm run build:electron
```

## 核心功能

### 用户认证
- JWT令牌认证机制
- 自动令牌刷新
- 本地存储管理

### 浏览器环境管理
- **创建环境**: 添加新的浏览器环境配置
- **查看环境**: 展示所有环境列表，支持分页
- **编辑环境**: 修改现有环境的配置信息
- **删除环境**: 移除不需要的环境配置
- **数据持久化**: 所有操作同步到后端数据库

### 系统监控
- CPU、内存、存储使用率显示
- 实时状态更新

## 项目结构

```
src/
├── components/          # 组件目录
│   ├── LoginForm.vue    # 登录表单组件
├── layouts/            # 布局组件
│   ├── MainLayout.vue   # 主布局组件
├── services/           # API服务
│   ├── api.ts          # API服务类
├── stores/             # 状态管理
│   ├── user.ts         # 用户状态store
│   ├── browser.ts      # 浏览器环境store
├── utils/              # 工具函数
│   ├── tokenManager.ts # Token管理器
├── router/             # 路由配置
├── App.vue             # 根组件
└── main.ts             # 应用入口

electron/
├── main.js             # Electron主进程
└── preload.js          # 预加载脚本
```

## 默认账户信息

- **用户名**: `13800138000`
- **密码**: `Tt123456`
- **验证码**: `666666`

## API集成

后端服务器地址: `http://localhost:7888`

### 支持的API端点

#### 用户认证
1. **用户登录**
   - POST `/api/app/login`
   - 无需认证

2. **获取用户信息**
   - GET `/api/app/user/info`
   - 需要Bearer Token认证

#### 浏览器环境管理
3. **获取环境列表**
   - POST `/api/app/browser/page`
   - 需要Bearer Token认证

4. **获取单个环境**
   - POST `/api/app/browser/get`
   - 需要Bearer Token认证

5. **创建环境**
   - POST `/api/app/browser/create`
   - 需要Bearer Token认证

6. **更新环境**
   - POST `/api/app/browser/update`
   - 需要Bearer Token认证

7. **删除环境**
   - POST `/api/app/browser/del`
   - 需要Bearer Token认证

## 本地存储结构

```javascript
// localStorage中存储的数据结构
{
  "access_token": "JWT访问令牌",
  "refresh_token": "JWT刷新令牌", 
  "token_expires": "访问令牌过期时间",
  "refresh_expires": "刷新令牌过期时间",
  "username": "用户名",
  "user_info": "{用户详细信息JSON}"
}
```

## 开发指南

### 添加新功能

1. 在相应的目录下创建新组件或服务
2. 如果需要状态管理，在`stores/`目录下创建新的store
3. 在`router/index.ts`中添加新的路由
4. 更新相关类型定义

### 样式规范

- 使用scoped CSS
- 遵循BEM命名规范
- 科技风格配色：淡蓝紫色渐变背景
- 毛玻璃效果和现代化动画

## 测试

### API测试
```bash
# 测试用户认证流程
node test-login-flow.js

# 测试浏览器环境API
node test-browser-api.js
```

## 打包发布

Electron应用会自动打包为对应平台的安装包：

- **Windows**: NSIS安装程序和ZIP包
- **macOS**: DMG镜像和ZIP包  
- **Linux**: AppImage和DEB包

## 故障排除

### 常见问题

1. **开发服务器无法启动**
   - 确保端口5173未被占用
   - 检查Node.js版本是否符合要求

2. **Electron窗口空白**
   - 确保先启动Vite开发服务器
   - 检查控制台是否有JavaScript错误

3. **API调用失败**
   - 确认后端服务器是否正常运行
   - 检查网络连接和CORS设置

4. **登录后无法跳转**
   - 检查浏览器控制台错误信息
   - 确认路由配置是否正确
   - 验证认证状态管理逻辑

5. **环境管理功能异常**
   - 确认数据库表结构是否正确
   - 检查后端服务是否正常运行
   - 验证API端点和请求格式

## 详细文档

- [浏览器环境管理功能说明](BROWSER_FEATURES.md)
- [科技风格设计指南](TECH_STYLE_GUIDE.md)

## 许可证

MIT License