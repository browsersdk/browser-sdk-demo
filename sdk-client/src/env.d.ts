/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
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
}