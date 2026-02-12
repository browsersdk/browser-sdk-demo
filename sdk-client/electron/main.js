const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');

let mainWindow;

function createWindow() {
  console.log('Creating main window...');
  
  mainWindow = new BrowserWindow({
    width: 1200,
    height: 800,
    minWidth: 800,
    minHeight: 600,
    webPreferences: {
      nodeIntegration: false,
      contextIsolation: true,
      preload: path.join(__dirname, 'preload.js')
    },
    icon: path.join(__dirname, '../public/favicon.ico'),
    backgroundColor: '#0a0a0a', // 改为深色背景匹配应用主题
    autoHideMenuBar: true, // 自动隐藏菜单栏
    resizable: true, // 允许调整窗口大小
    scrollBounce: false, // 禁用滚动弹性效果
    frame: true, // 保持窗口框架
    show: false, // 先隐藏窗口，等加载完成后再显示
  });

  // 等待页面加载完成后再显示窗口
  mainWindow.once('ready-to-show', () => {
    mainWindow.show();
    if (process.env.NODE_ENV === 'development') {
      mainWindow.webContents.openDevTools();
    }
  });

  // 开发环境下加载Vite开发服务器
  if (process.env.NODE_ENV === 'development') {
    console.log('Loading development server from http://localhost:5173');
    mainWindow.loadURL('http://localhost:5173');
  } else {
    // 生产环境下加载打包后的文件
    console.log('Loading production build');
    mainWindow.loadFile(path.join(__dirname, '../dist/index.html'));
  }

  // 监听窗口大小变化
  mainWindow.on('resize', () => {
    const [width, height] = mainWindow.getSize();
    console.log(`Window resized to ${width}x${height}`);
    
    // 通知渲染进程窗口大小变化
    mainWindow.webContents.send('window-resize', { width, height });
  });

  // 监听页面加载事件
  mainWindow.webContents.on('did-finish-load', () => {
    console.log('Page loaded successfully');
    
    // 获取当前窗口尺寸并发送给渲染进程
    const [width, height] = mainWindow.getSize();
    mainWindow.webContents.send('window-ready', { width, height });
    
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
    `);
  });
  
  mainWindow.webContents.on('did-fail-load', (event, errorCode, errorDescription, validatedURL) => {
    console.error('Page failed to load:', errorCode, errorDescription, validatedURL);
  });

  mainWindow.on('closed', () => {
    mainWindow = null;
  });
}

app.whenReady().then(() => {
  console.log('App is ready, creating window...');
  createWindow();

  app.on('activate', () => {
    console.log('App activated');
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow();
    }
  });
});

app.on('window-all-closed', () => {
  console.log('All windows closed');
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

// IPC通信 - 获取窗口信息
ipcMain.handle('get-window-size', () => {
  if (mainWindow) {
    return mainWindow.getSize();
  }
  return [1200, 800]; // 默认尺寸
});

ipcMain.handle('get-window-bounds', () => {
  if (mainWindow) {
    return mainWindow.getBounds();
  }
  return { x: 0, y: 0, width: 1200, height: 800 };
});

ipcMain.handle('set-window-size', (event, width, height) => {
  if (mainWindow) {
    mainWindow.setSize(width, height);
  }
});

// 窗口控制IPC
ipcMain.handle('minimize-window', () => {
  if (mainWindow) {
    mainWindow.minimize();
  }
});

ipcMain.handle('maximize-window', () => {
  if (mainWindow) {
    if (mainWindow.isMaximized()) {
      mainWindow.unmaximize();
    } else {
      mainWindow.maximize();
    }
  }
});

ipcMain.handle('close-window', () => {
  if (mainWindow) {
    mainWindow.close();
  }
});