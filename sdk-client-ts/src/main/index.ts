import { app, shell, BrowserWindow, session, globalShortcut } from 'electron'
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

  const shortcutKey = 'CommandOrControl+F12'
  const isRegistered = globalShortcut.register(shortcutKey, () => {
    mainWindow.webContents.toggleDevTools()
  })

  // 可选：额外注册 Mac 独有的 Control+F12（纯 Ctrl 键）
  if (process.platform === 'darwin') {
    globalShortcut.register('Control+F12', () => {
      mainWindow.webContents.toggleDevTools()
    })
  }

  // 验证注册结果
  if (!isRegistered) {
    console.log(`${shortcutKey} 快捷键注册失败（可能被占用）`)
  }
}

app.whenReady().then(async () => {
  electronApp.setAppUserModelId('com.yunlogin.fbmain')

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

app.on('will-quit', () => {
  globalShortcut.unregisterAll()
})
