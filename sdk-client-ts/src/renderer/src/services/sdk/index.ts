import { TokenManager } from '@/utils/tokenManager'
import { ApiService } from '@/services'
import { SdkCService } from './sdk'
import { SdkHttpService } from './sdkHttp'
import type { IResponse } from '../type'
import type { IOpenParams, ICloseParams } from '@type/sdk'

/**
 * SDK http服务
 */
export class SdkService {
  /**
   * sdk请求模式 0:C语言 1:http
   */
  static mode = 1
  /**
   * 初始化
   * @param port http服务端口
   */
  static async appInit(port: number): Promise<number> {
    await this.updateUsersig(false)
    // 初始化SDK
    const res = await SdkCService.appInit(port)
    const { code, msg } = res
    console.log('绑定SDK', code, msg)
    if (code === 0) localStorage.setItem('appPost', port + '')
    if (code !== 0) localStorage.setItem('appPost', '')
    return code
  }
  /**
   * 启动环境
   * @param data
   * @returns
   */
  static async open(data: IOpenParams): Promise<number> {
    await this.updateUsersig()
    const res: IResponse = await (this.mode === 0 ? window.sdkAPI.browserOpen(data) : SdkHttpService.open(data))
    const { code, msg } = res
    console.log('open', code, msg)
    return code
  }
  /**
   * 关闭环境
   * @param data
   * @returns
   */
  static async close(data: ICloseParams): Promise<number> {
    await this.updateUsersig()
    const res: IResponse = await (this.mode === 0 ? window.sdkAPI.browserClose(data) : SdkHttpService.close(data))
    const { code, msg } = res
    console.log('close', code, msg)
    return code
  }
  /**
   * 刷新UserSig
   * @param init 是否同步sdk
   */
  static async updateUsersig(init = true): Promise<void> {
    if (TokenManager.isUsersigExpired()) {
      console.log('usersig已过期')
      const userSig = await ApiService.getSdkUserSig()
      TokenManager.setUsersig(userSig)
      if (init) await SdkHttpService.tokenUpdate(userSig.userSig)
    }
  }
}
