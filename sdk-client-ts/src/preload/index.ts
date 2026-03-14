import { contextBridge, ipcRenderer } from 'electron'
import { electronAPI } from '@electron-toolkit/preload'
import type { IAppBindParams, IOpenParams, ICloseParams } from '@type/sdk'

const api = {}
const sdkAPI = {
  /** 初始化sdk */
  appBind: (data: IAppBindParams) => ipcRenderer.invoke('app-bind', data),
  /** 更新usersig */
  appTokenUpdate: (data: string) => ipcRenderer.invoke('app-token-update', data),
  /** 启动环境 */
  browserOpen: (data: IOpenParams) => ipcRenderer.invoke('app-browser-open', data),
  /** 关闭环境 */
  browserClose: (data: ICloseParams) => ipcRenderer.invoke('app-browser-close', data),
  /** 关闭sdk */
  appShutdown: () => ipcRenderer.invoke('app-shutdown')
}

if (process.contextIsolated) {
  try {
    contextBridge.exposeInMainWorld('electron', electronAPI)
    contextBridge.exposeInMainWorld('api', api)
    contextBridge.exposeInMainWorld('sdkAPI', sdkAPI)
  } catch (error) {
    console.error(error)
  }
} else {
  // @ts-ignore (define in dts)
  window.electron = electronAPI
  // @ts-ignore (define in dts)
  window.api = api
  // @ts-ignore (define in dts)
  window.sdkAPI = sdkAPI
}
