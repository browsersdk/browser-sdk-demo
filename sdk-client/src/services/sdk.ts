import { TokenManager } from '@/utils/tokenManager';
import type { IOpenParams } from '@/services'

/**
 * sdk C服务
 * @description 通过C服务操作SDK，electron项目推荐http服务方式
 */
export class SdkService {
  /** 
   * 初始化
   * @param port http服务端口
   */
  static async appInit(port: number): Promise<void> {
    // 初始化SDK
    const res = await window.sdkAPI.appBind({
      port,
      usersin: TokenManager.getUsersig(),
    })
    const { code, msg } = res
    console.log('绑定SDK', code, msg)
    if (code === 0) localStorage.setItem('appPost', port + '')
    if (code !== 0) localStorage.setItem('appPost', '')
  }

  static async open(data: IOpenParams): Promise<void> {
    const res = await window.sdkAPI.browserOpen(data)
    const { code, msg } = res
    console.log('open', code, msg)
  }
  static async close(id: number): Promise<void> {
    const res = await window.sdkAPI.browserClose({
      envId: id
    })
    const { code, msg } = res
    console.log('close', code, msg)
  }
}