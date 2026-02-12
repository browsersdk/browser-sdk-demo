const { contextBridge, ipcRenderer } = require('electron');

// 安全地暴露API给渲染进程
contextBridge.exposeInMainWorld('electronAPI', {
  getAppVersion: () => ipcRenderer.invoke('get-app-version'),
  minimizeWindow: () => ipcRenderer.invoke('minimize-window'),
  maximizeWindow: () => ipcRenderer.invoke('maximize-window'),
  closeWindow: () => ipcRenderer.invoke('close-window'),
  
  // 监听窗口状态变化
  onWindowMaximized: (callback) => ipcRenderer.on('window-maximized', callback),
  onWindowUnmaximized: (callback) => ipcRenderer.on('window-unmaximized', callback),
  
  // 移除监听器
  removeWindowMaximizedListener: (callback) => ipcRenderer.removeListener('window-maximized', callback),
  removeWindowUnmaximizedListener: (callback) => ipcRenderer.removeListener('window-unmaximized', callback)
});