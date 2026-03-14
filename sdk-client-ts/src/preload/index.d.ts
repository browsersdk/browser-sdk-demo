import { ElectronAPI } from '@electron-toolkit/preload'
import type { IResponse, IAppBindParams, IOpenParams, ICloseParams } from '@type/sdk'

interface ISdkAPI {
  /** 初始化sdk */
  appBind: (port: IAppBindParams) => Promise<IResponse>
  /** 更新usersig */
  appTokenUpdate: (userSig: string) => Promise<IResponse>
  /** 启动环境 */
  browserOpen: (data: IOpenParams) => Promise<IResponse>
  /** 关闭环境 */
  browserClose: (data: ICloseParams) => Promise<IResponse>
  /** 关闭sdk */
  appShutdown: () => Promise<IResponse>
}

declare global {
  interface Window {
    electron: ElectronAPI
    api: unknown
    sdkAPI: ISdkAPI
  }
}
