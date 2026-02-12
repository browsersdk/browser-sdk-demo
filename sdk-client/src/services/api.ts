import { TokenManager, type AuthTokens } from '@/utils/tokenManager';

const BASE_URL = 'http://localhost:7888';

interface LoginRequest {
  code: string;
  devId: string;
  inviteCode: string;
  name: string;
  os: string;
  osVer: string;
  password: string;
  src: string;
  tid: string;
  username: string;
  uuid: string;
  ver: string;
}

interface LoginResponse {
  code: number;
  data: AuthTokens;
  msg: string;
  reqId: string;
}

interface UserInfo {
  id: number;
  username: string;
  phone: string;
  email: string;
  nickname: string;
  name: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

interface UserInfoResponse {
  reqId: string;
  code: number;
  msg: string;
  data: UserInfo;
}

// Browser相关接口定义
interface Browser {
  id: number;
  name: string;
  envId: number;
  userId: number;
  data: string;
  createdAt: string;
  updatedAt: string;
  status: number;  // 1: 停止, 3: 启动
}

interface BrowserDto {
  id?: number;
  name: string;
  envId?: number;  // 变为可选
  userId: number;
  data: string;
  status?: number;  // 可选，创建时可不传
}

interface BrowserGetPageReq {
  page: number;
  size: number;
  name?: string;
  envId?: number;
  userId?: number;
}

interface PageResponse<T> {
  list: T[];
  total: number;
  page: number;
  size: number;
}

interface BaseResponse<T> {
  code: number;
  data: T;
  msg: string;
  reqId: string;
}

interface ReqId {
  id: number;
}

interface ReqIds {
  ids: number[];
}

export class ApiService {
  private static async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const url = `${BASE_URL}${endpoint}`;
    
    const defaultHeaders = {
      'Content-Type': 'application/json',
    };

    // 如果已认证，添加Authorization头
    const accessToken = TokenManager.getAccessToken();
    if (accessToken && !TokenManager.isTokenExpired()) {
      (defaultHeaders as Record<string, string>)['Authorization'] = `Bearer ${accessToken}`;
    }

    const config: RequestInit = {
      ...options,
      headers: {
        ...defaultHeaders,
        ...options.headers,
      },
    };

    try {
      const response = await fetch(url, config);
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('API request failed:', error);
      throw error;
    }
  }

  static async login(username: string, password: string): Promise<AuthTokens> {
    const requestBody: LoginRequest = {
      code: '666666',
      devId: '',
      inviteCode: '',
      name: '',
      os: '',
      osVer: '',
      password: password,
      src: '',
      tid: '',
      username: username,
      uuid: '',
      ver: ''
    };

    const response = await this.request<LoginResponse>('/api/app/login', {
      method: 'POST',
      body: JSON.stringify(requestBody),
    });

    // 根据实际API响应调整判断条件
    if (response.code === 200 || response.code === 0) {
      return response.data;
    } else {
      throw new Error(response.msg || 'Login failed');
    }
  }

  static async getUserInfo(): Promise<UserInfo> {
    const response = await this.request<UserInfoResponse>('/api/app/user/info');
    
    if (response.code === 200) {
      return response.data;
    } else {
      throw new Error(response.msg || 'Failed to get user info');
    }
  }

  static logout(): void {
    TokenManager.clearTokens();
  }

  // Browser相关API方法
  static async getBrowserList(params: BrowserGetPageReq): Promise<PageResponse<Browser>> {
    const response = await this.request<BaseResponse<PageResponse<Browser>>>('/api/app/browser/page', {
      method: 'POST',
      body: JSON.stringify(params),
    });
    
    if (response.code === 200) {
      return response.data;
    } else {
      throw new Error(response.msg || 'Failed to get browser list');
    }
  }

  static async getBrowser(id: number): Promise<Browser> {
    const response = await this.request<BaseResponse<Browser>>('/api/app/browser/get', {
      method: 'POST',
      body: JSON.stringify({ id } as ReqId),
    });
    
    if (response.code === 200) {
      return response.data;
    } else {
      throw new Error(response.msg || 'Failed to get browser');
    }
  }

  static async createBrowser(browser: BrowserDto): Promise<Browser> {
    const response = await this.request<BaseResponse<Browser>>('/api/app/browser/create', {
      method: 'POST',
      body: JSON.stringify(browser),
    });
    
    if (response.code === 200) {
      return response.data;
    } else {
      throw new Error(response.msg || 'Failed to create browser');
    }
  }

  static async updateBrowser(browser: BrowserDto): Promise<Browser> {
    const response = await this.request<BaseResponse<Browser>>('/api/app/browser/update', {
      method: 'POST',
      body: JSON.stringify(browser),
    });
    
    if (response.code === 200) {
      return response.data;
    } else {
      throw new Error(response.msg || 'Failed to update browser');
    }
  }

  static async deleteBrowser(ids: number[]): Promise<void> {
    const response = await this.request<BaseResponse<any>>('/api/app/browser/del', {
      method: 'POST',
      body: JSON.stringify({ ids } as ReqIds),
    });
    
    if (response.code !== 200) {
      throw new Error(response.msg || 'Failed to delete browser');
    }
  }
}