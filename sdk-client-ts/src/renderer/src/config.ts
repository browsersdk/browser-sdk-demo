// app配置文件
// 优先读环境变量 VITE_API_URL，未配置时按开发/生产模式取默认值

/**
 * 获取 API 基础 URL
 * 优先级: 环境变量 > 默认值
 */
export const BASE_API_URL: string =
  (import.meta.env.VITE_API_URL as string | undefined) ||
  (import.meta.env.DEV ? 'http://localhost:7888' : 'http://192.168.0.127:7888')

/**
 * 应用配置对象
 */
export const APP_CONFIG = {
  name: import.meta.env.VITE_APP_NAME || 'Browser SDK Client',
  version: import.meta.env.VITE_APP_VERSION || '1.0.0',
  apiUrl: BASE_API_URL,
  enableDebug: import.meta.env.VITE_ENABLE_DEBUG === 'true',
  mode: import.meta.env.MODE, // 当前构建模式: development | production | staging
} as const

/**
 * 获取当前环境信息
 */
export const ENV_INFO = {
  isDev: import.meta.env.DEV,
  isProd: import.meta.env.PROD,
  mode: import.meta.env.MODE,
  baseUrl: BASE_API_URL,
} as const




