// app配置文件
// 优先读环境变量 VITE_API_URL，未配置时按开发/生产模式取默认值
export const BASE_API_URL: string =
  (import.meta.env.VITE_API_URL as string | undefined) ||
  (import.meta.env.DEV ? 'http://localhost:7888' : 'http://192.168.0.127:7888')



