import { TokenManager } from '@/utils/tokenManager';
import type { IResponse } from './type/type'
import type { IOpenParams, ICloseParams } from './type/app'
import axios from '@/utils/axios'

const BASE_URL = 'http://localhost';

/**
 * SDK http服务
 */
export class SdkHttpService {
  static async open(data: IOpenParams): Promise<void> {
    const res = await window.sdkAPI.browserOpen(data)

    // const res: IResponse = await axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/browser/open`, data)
    const { code, msg } = res
    console.log('open', code, msg)
  }
  static async close(data: ICloseParams): Promise<void> {
    const res: IResponse = await axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/browser/close`, data)
    const { code, msg } = res
    console.log('close', code, msg)
  }
}