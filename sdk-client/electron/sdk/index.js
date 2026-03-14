const { app, ipcMain } = require('electron')
const path = require('path')
const BroSDK = require('./brosdk')

class SDK {
  bindStatus = false
  broSDK
  constructor() {
    this.broSDK = new BroSDK()

    ipcMain.handle('app-bind', this.init)
    ipcMain.handle('app-token-update', this.tokenUpdate)
    ipcMain.handle('app-browser-open', this.browserOpen)
    ipcMain.handle('app-browser-close', this.browserClose)
    ipcMain.handle('app-shutdown', this.shutdown)
  }
  init = async (_event, data) => {
    const initParam = {
      port: 65535,
      userSig: data.usersin,
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
      msg,
    }
  }

  tokenUpdate = async (_event, data) => {
    const res = await this.broSDK.tokenUpdate(JSON.stringify(data))
    return {
      code: res,
      msg: '',
    }
  }
  browserOpen = async (_event, data) => {
    console.log('启动环境', data)
    const res = await this.broSDK.browserOpen(JSON.stringify(data))
    console.log(res)
    return {
      code: res,
      msg: '',
    }
  }
  browserClose = async (_event, data) => {
    console.log('关闭环境', data)
    const res = await this.broSDK.browserClose(data)
    console.log(res)
    return {
      code: res,
      msg: '',
    }
  }
  shutdown = async (_event) => {
    if (!this.bindStatus) return

    const res = await this.broSDK.shutdown()
    let msg = 'Shutdown failed.'
    if (res.code === 0) {
      msg = 'Shutdown successful.'
      this.bindStatus = false
    }
    return {
      code: res.code,
      msg: '',
    }
  }
}

module.exports = SDK
