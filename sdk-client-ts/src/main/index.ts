import { app, shell, BrowserWindow, session } from 'electron'
import path, { join } from 'path'
import { electronApp, optimizer, is } from '@electron-toolkit/utils'
import icon from '../../resources/icon.png?asset'
import Logger from './logs/logs'
import SDK from './sdk'

let mainWindow: BrowserWindow

const log = Logger.getInstance()
new SDK()

async function installVueDevtools(): Promise<void> {
  if (app.isPackaged) return

  try {
    // 项目内插件存放路径（示例：项目根目录下的 extensions/vue-devtools）
    const vueDevtoolsPath = path.join(
      process.cwd(), // 项目根目录
      'extensions', // 插件存放目录
      'vue-devtools' // Vue Devtools 插件目录
    )

    // 加载项目内的 Vue Devtools 插件
    await session.defaultSession.loadExtension(vueDevtoolsPath)
  } catch {
    log.info('vueDevtoolsPath加载错误')
  }
}

function createWindow(): void {
  log.info('Creating main window...')
  log.info('process.resourcesPath', path.join(process.resourcesPath, 'sdk', 'brosdk.dll'))
  log.info('__dirname', path.join(__dirname, 'brosdk', 'brosdk.dll'))
  mainWindow = new BrowserWindow({
    width: 1200,
    height: 800,
    minWidth: 800,
    minHeight: 600,
    show: false,
    autoHideMenuBar: true,
    ...(process.platform === 'linux' ? { icon } : {}),
    webPreferences: {
      preload: join(__dirname, '../preload/index.js'),
      sandbox: false,
      devTools: true,
      webSecurity: false
    }
  })

  mainWindow.on('ready-to-show', () => {
    mainWindow.show()
    if (!app.isPackaged) {
      mainWindow.webContents.openDevTools()
    }
  })

  mainWindow.webContents.setWindowOpenHandler((details) => {
    shell.openExternal(details.url)
    return { action: 'deny' }
  })

  if (is.dev && process.env['ELECTRON_RENDERER_URL']) {
    mainWindow.loadURL(process.env['ELECTRON_RENDERER_URL'])
  } else {
    mainWindow.loadFile(join(__dirname, '../renderer/index.html'))
  }
}

app.whenReady().then(async () => {
  electronApp.setAppUserModelId('com.electron')

  app.on('browser-window-created', (_, window) => {
    optimizer.watchWindowShortcuts(window)
  })

  await installVueDevtools()
  createWindow()

  app.on('activate', function () {
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
  })
})

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})
