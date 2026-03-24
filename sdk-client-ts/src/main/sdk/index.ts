import { app, ipcMain } from 'electron'
import path from 'path'
import type { IResponse, IAppBindParams, IOpenParams, ICloseParams } from '@type/sdk'
import BroSDK from './brosdk'

export default class SDK {
  private bindStatus = false
  private broSDK: BroSDK
  private workDir: string = ''
  constructor() {
    this.broSDK = new BroSDK()

    ipcMain.handle('app-bind', this.init)
    ipcMain.handle('app-token-update', this.tokenUpdate)
    ipcMain.handle('app-browser-open', this.browserOpen)
    ipcMain.handle('app-browser-close', this.browserClose)
    ipcMain.handle('app-shutdown', this.shutdown)
  }
  init = async (_event, data: IAppBindParams): Promise<IResponse> => {
    const isWindows = process.platform === 'win32' // Windows 系统
    const isMac = process.platform === 'darwin' // macOS 系统
    // console.log(path.join(app.getAppPath(), '..','workDir'), path.join(path.join(app.getPath('exe')), '..', '..', '..','workDir'))

    // 设置工作目录
    if (app.isPackaged) {
      if (isWindows) {
        this.workDir = path.join(process.resourcesPath, 'workDir')
      }
      if (isMac) {
        const exePath = app.getPath('exe')
        const appDir = path.join(path.dirname(exePath), '..', '..')
        const macArm64Dir = path.join(appDir, '..')
        this.workDir = path.join(macArm64Dir, 'workDir')
      }
    } else {
      this.workDir = path.join(app.getAppPath(), '..', 'workDir-win')
    }
    console.log('内核路径：', this.workDir)

    const initParam = {
      port: 65535,
      userSig: data.usersin,
      workDir: this.workDir
    }
    this.broSDK.registerCookiesStorageCb((cookies) => {
      console.log('cookies...', cookies)
      return null
    })
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

    const code = await this.broSDK.shutdown()
    if (code === 0) {
      this.bindStatus = false
    }
    return {
      code,
      msg: ''
    }
  }
}
