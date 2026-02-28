import axios from 'axios'
import type { InternalAxiosRequestConfig, AxiosResponse } from 'axios'
import { TokenManager, type AuthTokens } from '@/utils/tokenManager';
import { BASE_API_URL } from '@/config'

// 创建 axios 实例
const requester = axios.create({
  baseURL: BASE_API_URL, // 你的 API 基础地址
  timeout: 10000, // 请求超时时间
  headers: {
    'Content-Type': 'application/json',
    // 可以添加默认请求头，如 Token 等
  },
})

// 设置基础路径
if (import.meta.env.DEV) {
  requester.defaults.baseURL = 'http://192.168.0.127:7888'
} else {
  requester.defaults.baseURL = 'http://192.168.0.127:7888'
}

/**
 * http request 拦截器
 */
requester.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const accessToken = TokenManager.getAccessToken();
    if (accessToken && !TokenManager.isTokenExpired()) {
      config.headers.Authorization = `Bearer ${accessToken}`
    }
    return config
  },
  (err) => {
    // 请求错误的处理
    return Promise.reject(err)
  }
)

// http response 拦截器
requester.interceptors.response.use(
  // 处理响应数据
  (response: AxiosResponse) => {
    return response.data
  },
  // 抛出错误
  (error) => {
    return Promise.reject(error.response)
  }
)

export default requester
