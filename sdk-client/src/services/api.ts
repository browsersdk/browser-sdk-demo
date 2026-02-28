import { TokenManager, type AuthTokens } from '@/utils/tokenManager';
import axios from '@/utils/axios'
import type { LoginRequest, LoginResponse, UserInfo, UserInfoResponse, Browser, BrowserDto, BrowserGetPageReq, PageResponse, BaseResponse } from './type'

const BASE_URL = 'http://192.168.0.127:7888';

export class ApiService {

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

    const res: LoginResponse = await axios.post('/api/app/login', requestBody);
    const { code, data, msg } = res;

    // 根据实际API响应调整判断条件
    if (code === 200 || code === 0) {
      return data;
    } else {
      throw new Error(msg || 'Login failed');
    }
  }

  static async getUserInfo(): Promise<UserInfo> {
    const res: UserInfoResponse = await axios.get('/api/app/user/info');
    const { code, data, msg } = res;

    if (code === 200) {
      return data;
    } else {
      throw new Error(msg || 'Failed to get user info');
    }
  }

  static logout(): void {
    TokenManager.clearTokens();
  }

  // Browser相关API方法
  static async getBrowserList(params: BrowserGetPageReq): Promise<PageResponse<Browser>> {
    const res: BaseResponse<PageResponse<Browser>> = await axios.post('/api/app/browser/page', params);
    const { code, data, msg } = res;

    if (code === 200) {
      return data;
    } else {
      throw new Error(msg || 'Failed to get browser list');
    }
  }

  static async getBrowser(id: number): Promise<Browser> {
    const res: BaseResponse<Browser> = await axios.post('/api/app/browser/get', { id });
    const { code, data, msg } = res;

    if (code === 200) {
      return data;
    } else {
      throw new Error(msg || 'Failed to get browser');
    }
  }

  static async createBrowser(browser: BrowserDto): Promise<Browser> {
    const res: BaseResponse<Browser> = await axios.post('/api/app/browser/create', browser);
    const { code, data, msg } = res;

    if (code === 200) {
      return data;
    } else {
      throw new Error(msg || 'Failed to create browser');
    }
  }

  static async updateBrowser(browser: BrowserDto): Promise<Browser> {
    const res: BaseResponse<Browser> = await axios.post('/api/app/browser/update', browser);
    const { code, data, msg } = res;

    if (code === 200) {
      return data;
    } else {
      throw new Error(msg || 'Failed to update browser');
    }
  }

  static async deleteBrowser(ids: number[]): Promise<void> {
    const res: BaseResponse<void> = await axios.post('/api/app/browser/del', { ids });
    const { code, data, msg } = res;

    if (code !== 200) {
      throw new Error(msg || 'Failed to delete browser');
    }
  }
}