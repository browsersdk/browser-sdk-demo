import type { IResponse, IOpenParams, ICloseParams } from '@type/sdk'
import axios from '@/utils/axios'

const BASE_URL = 'http://localhost'

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
