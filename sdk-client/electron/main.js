const { app, BrowserWindow, ipcMain, session } = require('electron')
const path = require('path')

const logger = require('./logs/logs')
const SDK = require('./sdk')

let mainWindow

const log = logger.getInstance()
const sdk = new SDK()

async function installVueDevtools() {
  if (app.isPackaged) return

  try {
    // 项目内插件存放路径（示例：项目根目录下的 extensions/vue-devtools）
    const vueDevtoolsPath = path.join(
      process.cwd(), // 项目根目录
      'extensions', // 插件存放目录
      'vue-devtools', // Vue Devtools 插件目录
    )

    // 加载项目内的 Vue Devtools 插件
    const extension = await session.defaultSession.loadExtension(vueDevtoolsPath)
  } catch (err) {
    log.info('vueDevtoolsPath加载错误')
  }
}

function createWindow() {
  log.info('Creating main window...')

  mainWindow = new BrowserWindow({
    width: 1200,
    height: 800,
    minWidth: 800,
    minHeight: 600,
    webPreferences: {
      nodeIntegration: false,
      contextIsolation: true,
      preload: path.join(__dirname, 'preload.js'),
      webSecurity: false, // 禁用 CSP 限制
    },
    icon: path.join(__dirname, '../public/favicon.ico'),
    backgroundColor: '#0a0a0a', // 改为深色背景匹配应用主题
    autoHideMenuBar: true, // 自动隐藏菜单栏
    resizable: true, // 允许调整窗口大小
    scrollBounce: false, // 禁用滚动弹性效果
    frame: true, // 保持窗口框架
    show: false, // 先隐藏窗口，等加载完成后再显示
  })
  // 等待页面加载完成后再显示窗口
  mainWindow.once('ready-to-show', () => {
    mainWindow.show()
    if (!app.isPackaged) {
      mainWindow.webContents.openDevTools()
    }
  })

  // 开发环境下加载Vite开发服务器
  if (!app.isPackaged) {
    console.log('Loading development server from http://localhost:5173')
    mainWindow.loadURL('http://localhost:5173')
  } else {
    // 生产环境下加载打包后的文件
    log.info('PROD MODE')

    const indexPath = path.join(__dirname, '../dist/index.html')

    log.info('indexPath:', indexPath)

    mainWindow.loadFile(indexPath)
  }

  // 监听窗口大小变化
  mainWindow.on('resize', () => {
    const [width, height] = mainWindow.getSize()
    // console.log(`Window resized to ${width}x${height}`);

    // 通知渲染进程窗口大小变化
    mainWindow.webContents.send('window-resize', { width, height })
  })

  // 监听页面加载事件
  mainWindow.webContents.on('did-finish-load', () => {
    console.log('Page loaded successfully')
    log.debug('初始化成功')
    // 获取当前窗口尺寸并发送给渲染进程
    const [width, height] = mainWindow.getSize()
    mainWindow.webContents.send('window-ready', { width, height })

    // 注入CSS来隐藏滚动条和优化布局
    mainWindow.webContents.insertCSS(`
      ::-webkit-scrollbar {
        width: 0px !important;
        height: 0px !important;
        background: transparent !important;
      }
      ::-webkit-scrollbar-track {
        background: transparent !important;
      }
      ::-webkit-scrollbar-thumb {
        background: transparent !important;
      }
      html {
        scrollbar-width: none !important;
      }
      body {
        -ms-overflow-style: none !important;
        height: 100vh !important;
        overflow: hidden !important;
      }
      #app {
        height: 100vh !important;
        overflow: hidden !important;
      }
    `)
  })

  mainWindow.webContents.on('did-fail-load', (event, errorCode, errorDescription, validatedURL) => {
    console.error('Page failed to load:', errorCode, errorDescription, validatedURL)
  })

  mainWindow.on('closed', () => {
    mainWindow = null
  })
}

app.whenReady().then(async () => {
  console.log('App is ready, creating window...')
  await installVueDevtools() // 安装 Vue Devtools
  createWindow()
  app.on('activate', () => {
    console.log('App activated')
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow()
    }
  })
})

app.on('window-all-closed', () => {
  console.log('All windows closed')
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

// IPC通信 - 获取窗口信息
ipcMain.handle('get-window-size', () => {
  if (mainWindow) {
    return mainWindow.getSize()
  }
  return [1200, 800] // 默认尺寸
})

ipcMain.handle('get-window-bounds', () => {
  if (mainWindow) {
    return mainWindow.getBounds()
  }
  return { x: 0, y: 0, width: 1200, height: 800 }
})

ipcMain.handle('set-window-size', (event, width, height) => {
  if (mainWindow) {
    mainWindow.setSize(width, height)
  }
})

// 窗口控制IPC
ipcMain.handle('minimize-window', () => {
  if (mainWindow) {
    mainWindow.minimize()
  }
})

ipcMain.handle('maximize-window', () => {
  if (mainWindow) {
    if (mainWindow.isMaximized()) {
      mainWindow.unmaximize()
    } else {
      mainWindow.maximize()
    }
  }
})

ipcMain.handle('close-window', () => {
  if (mainWindow) {
    mainWindow.close()
  }
})
