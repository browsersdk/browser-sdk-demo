import type { IResponse, IBrowserInfoResponse, IOpenParams, ICloseParams } from '@type/sdk'
import axios from '@/utils/axios'

const BASE_URL = 'http://localhost'

/**
 * SDK http服务
 */
export class SdkHttpService {
  static info(): Promise<IBrowserInfoResponse> {
    console.log('sdk info');
    return axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/browser/info`)
  }
  static open(data: IOpenParams): Promise<IResponse> {
    console.log('sdk open', data);
    return axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/browser/open`, data)
  }
  static close(data: ICloseParams): Promise<IResponse> {
    console.log('sdk close', data);
    return axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/browser/close`, data)
  }
  static tokenUpdate(userSig: string): Promise<IResponse> {
    console.log('sdk tokenUpdate', userSig);
    return axios.post(`${BASE_URL}:${localStorage.getItem('appPost')}/sdk/v1/token/update`, {
      userSig
    })
  }
}
