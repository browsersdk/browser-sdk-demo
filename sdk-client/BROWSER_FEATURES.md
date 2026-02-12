# 浏览器环境管理功能说明

## 功能概述

本项目实现了完整的浏览器环境管理功能，包括创建、读取、更新、删除(CRUD)操作。

## 技术实现

### 后端API (browser.go)
- **QueryPage**: 获取浏览器环境列表（分页）
- **Get**: 获取单个浏览器环境详情
- **Create**: 创建新的浏览器环境
- **Update**: 更新浏览器环境信息
- **Del**: 删除浏览器环境

### 前端实现

#### 1. API服务层 (`src/services/api.ts`)
扩展了`ApiService`类，添加了以下方法：
- `getBrowserList()`: 获取环境列表
- `getBrowser()`: 获取单个环境
- `createBrowser()`: 创建环境
- `updateBrowser()`: 更新环境
- `deleteBrowser()`: 删除环境

#### 2. 状态管理层 (`src/stores/browser.ts`)
创建了`useBrowserStore` Pinia store，包含：
- 浏览器环境列表状态管理
- 加载状态跟踪
- 分页信息管理
- 完整的CRUD操作方法

#### 3. UI组件层 (`src/layouts/MainLayout.vue`)
在主控制台界面中集成了：
- 环境列表展示
- 创建/编辑模态框
- 删除确认对话框
- 加载状态指示器
- 响应式设计

## 数据结构

### Browser模型
```typescript
interface Browser {
  id: number;        // 主键ID（服务器生成）
  name: string;      // 环境名称（必填）
  envId: number;     // 环境ID（服务器自动生成）
  userId: number;    // 用户ID（必填）
  data: string;      // 配置数据(JSON格式，必填）
  createdAt: string; // 创建时间（服务器生成）
  updatedAt: string; // 更新时间（服务器生成）
}
```

## 使用方法

### 1. 环境列表查看
- 登录后自动加载环境列表
- 支持按名称搜索过滤
- 显示环境的基本信息和配置数据

### 2. 创建新环境
- 点击"CREATE NEW ENVIRONMENT"按钮
- 填写环境名称和配置数据
- **envId由服务器自动生成，无需填写**
- 提交后自动刷新列表

### 3. 编辑环境
- 点击环境卡片的"EDIT"按钮
- 修改环境名称和配置数据
- **envId保持不变，由服务器管理**
- 保存后实时更新列表

### 4. 删除环境
- 点击环境卡片的"DELETE"按钮
- 确认删除操作
- 成功后从列表中移除

## 数据库配置

### 表结构 (`sys_browser`)
```sql
CREATE TABLE `sys_browser` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `env_id` bigint unsigned NOT NULL DEFAULT '0',
  `user_id` int NOT NULL DEFAULT '0',
  `data` text,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_browser_deleted_at` (`deleted_at`),
  KEY `idx_sys_browser_user_id` (`user_id`),
  KEY `idx_sys_browser_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

## 注意事项

1. **认证要求**: 所有API调用都需要有效的JWT令牌
2. **数据格式**: `data`字段应为JSON字符串格式
3. **环境ID**: `envId`由服务器自动生成和管理，前端无需提供
4. **用户隔离**: 每个用户只能访问自己的环境数据
5. **必填字段**: 只有`name`和`data`是必填项

## 错误处理

- 网络错误会显示友好提示
- API错误会记录到控制台
- 删除操作需要二次确认
- 表单验证确保数据完整性

## 测试

可以使用提供的测试脚本验证API功能：
```bash
node test-browser-api.js
```

该脚本会依次测试创建、查询、更新、删除等完整流程，特别验证了envId由服务器自动生成的功能。