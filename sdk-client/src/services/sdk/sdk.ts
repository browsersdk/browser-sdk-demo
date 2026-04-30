import { TokenManager } from '@/utils/tokenManager';
import type { IOpenParams } from '@/services'
import type { IResponse } from '../type'

/**
 * sdk C服务
 * @description 通过C服务操作SDK，electron项目推荐http服务方式
 */
export class SdkCService {
  /** 
   * 初始化
   * @param port http服务端口
   */
  static appInit(port: number): Promise<IResponse> {
    return window.sdkAPI.appBind({
      port,
      usersin: TokenManager.getUsersig(),
    })
  }

  static open(data: IOpenParams): Promise<IResponse> {
    return window.sdkAPI.browserOpen(data)
  }
  static async close(id: number): Promise<IResponse> {
    return window.sdkAPI.browserClose({
      envId: id
    })
  }
}