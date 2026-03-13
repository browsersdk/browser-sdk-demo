import { ipcMain } from 'electron'
import type { IResponse, IAppBindParams, IOpenParams, ICloseParams } from '@type/sdk'
import BroSDK from './brosdk'

export default class SDK {
  private bindStatus = false
  private broSDK: BroSDK
  constructor() {
    this.broSDK = new BroSDK()

    ipcMain.handle('app-bind', this.init)
    ipcMain.handle('app-token-update', this.tokenUpdate)
    ipcMain.handle('app-browser-open', this.browserOpen)
    ipcMain.handle('app-browser-close', this.browserClose)
    ipcMain.handle('app-shutdown', this.shutdown)
  }
  init = async (_event, data: IAppBindParams): Promise<IResponse> => {
    const initParam = {
      port: 65535,
      userSig: data.usersin
    }
    const res = await this.broSDK.init(JSON.stringify(initParam))
    console.log('res...', initParam, res)
    let msg = 'Initialization failed.'
    if (res.code === 0) {
      msg = 'Initialization successful.'
      this.bindStatus = true
    }
    return {
      code: res.code,
      msg
    }
  }

  tokenUpdate = async (_event, data): Promise<IResponse> => {
    const res = await this.broSDK.tokenUpdate(JSON.stringify(data))
    return {
      code: res,
      msg: ''
    }
  }
  browserOpen = async (_event, data: IOpenParams): Promise<IResponse> => {
    console.log('启动环境', data)
    const res = await this.broSDK.browserOpen(JSON.stringify(data))
    console.log(res)
    return {
      code: res,
      msg: ''
    }
  }
  browserClose = async (_event, data: ICloseParams): Promise<IResponse> => {
    console.log('关闭环境', data)
    const res = await this.broSDK.browserClose(data)
    console.log(res)
    return {
      code: res,
      msg: ''
    }
  }
  shutdown = async (_event): Promise<IResponse | void> => {
    if (!this.bindStatus) return

    const res = await this.broSDK.shutdown()
    if (res.code === 0) {
      this.bindStatus = false
    }
    return {
      code: res.code,
      msg: ''
    }
  }
}
