/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

interface IResponse {
  /** 状态码 */
  code: number;
  /** 错误或提示信息 */
  msg: string;
}


interface sdkAPI {
  appBind: (data) => Promise<IResponse>;
  appTokenUpdate: (data) => Promise<IResponse>;
  browserOpen: (data) => Promise<IResponse>;
  browserClose: (data) => Promise<IResponse>;
  appShutdown: () => Promise<IResponse>;
}

interface Window {
  electronAPI?: {
    getAppVersion: () => Promise<string>;
    minimizeWindow: () => Promise<void>;
    maximizeWindow: () => Promise<void>;
    closeWindow: () => Promise<void>;
    onWindowMaximized: (callback: () => void) => void;
    onWindowUnmaximized: (callback: () => void) => void;
    removeWindowMaximizedListener: (callback: () => void) => void;
    removeWindowUnmaximizedListener: (callback: () => void) => void;
  }
  sdkAPI: sdkAPI;
}