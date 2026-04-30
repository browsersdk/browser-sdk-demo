import { TokenManager } from '@/utils/tokenManager';
import type { IResponse } from '../type'
import type { IOpenParams, ICloseParams } from './type'
import axios from '@/utils/axios'

const BASE_URL = 'http://localhost';

/**
 * SDK http服务
 */
export class SdkHttpService {
  static open(data: IOpenParams): Promise<IResponse> {
    return axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/browser/open`, data)
  }
  static close(data: ICloseParams): Promise<IResponse> {
    return axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/browser/close`, data)
  }
  static tokenUpdate(userSig: string): Promise<IResponse> {
    return axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/token/update`, {
      userSig
    })
  }
}