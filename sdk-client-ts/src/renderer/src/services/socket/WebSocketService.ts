import { useBrowserStore } from '@/stores/browser'
import type { SocketOptions } from './type'
import { OperateEvent } from './type'

type MessageHandler = (data: unknown) => void

export class WebSocketService {
  private static instance: WebSocketService
  private socket: WebSocket | null = null
  private port: number = 0
  private url: string = ''
  private options: Required<SocketOptions>
  private reconnectAttempts = 0
  private heartbeatTimer: number | null = null
  private readonly handlers = new Map<string, MessageHandler>()
  private browserStore

  private constructor() {
    // 默认配置
    this.options = {
      heartbeatInterval: 30000,
      reconnectInterval: 5000,
      maxReconnectAttempts: 5
    }
    this.browserStore = useBrowserStore()
  }

  public static getInstance(): WebSocketService {
    if (!WebSocketService.instance) {
      WebSocketService.instance = new WebSocketService()
    }
    return WebSocketService.instance
  }

  /**
   * 连接 WebSocket 服务器
   * @param url 服务器地址
   * @param options 配置项
   */
  public connect(port: number, options?: SocketOptions): void {
    this.port = port
    this.url = `http://localhost:${port}`
    this.options = { ...this.options, ...options }

    if (this.socket) {
      console.warn('WebSocket 已经连接，正在关闭旧连接...')
      this.disconnect()
    }

    this.socket = new WebSocket(this.url)

    this.socket.onopen = (event) => {
      console.log('WebSocket 连接已建立')
      this.reconnectAttempts = 0
      this.startHeartbeat()
      // 触发 open 事件的处理函数
      this.emit('open', event)
    }

    this.socket.onmessage = this.onMessage

    this.socket.onerror = (error) => {
      console.error('WebSocket 错误:', error)
      this.emit('error', error)
    }

    this.socket.onclose = (event) => {
      console.log('WebSocket 连接已关闭:', event.code, event.reason)
      this.stopHeartbeat()
      this.socket = null
      // 自动重连逻辑
      this.attemptReconnect()
      this.emit('close', event)
    }
  }

  private onMessage = async (evt: MessageEvent): Promise<void> => {
    try {
      const { type, data } = JSON.parse(evt.data)
      console.log('收到消息:', type, data)

      switch (type) {
        case OperateEvent.OPEN_SUCCESS:
          // 启动成功：使用 store 统一方法处理
          this.browserStore.handleOpenSuccess(data.envId)

          // 同步更新服务端状态为运行中（status: 3）
          await this.browserStore.syncBrowserStatus(data, 3)
          break

        case OperateEvent.OPEN_FAILED:
          // 启动失败：使用 store 统一方法处理
          this.browserStore.handleOpenFailed(data.envId)

          // 同步更新服务端状态为停止（status: 1）
          await this.browserStore.syncBrowserStatus(data, 1)
          break

        case OperateEvent.CLOSE_SUCCESS:
          // 停止成功：使用 store 统一方法处理
          this.browserStore.handleCloseSuccess(data.envId)

          // 同步更新服务端状态为停止（status: 1）
          await this.browserStore.syncBrowserStatus(data, 1)
          break

        case OperateEvent.CLOSE_FAILED:
          // 停止失败：使用 store 统一方法处理
          this.browserStore.handleCloseFailed(data.envId)

          // 停止失败时，环境仍处于运行状态，无需更新服务端状态
          // 如需重新添加到已启动状态，取消以下注释：
          // this.browserStore.startedDict.value.add(data.envId)
          break

        default:
          console.log('未处理的消息类型:', type, data)
      }
    } catch (error) {
      console.error('处理 WebSocket 消息失败:', error, evt.data)
    }
  }

  /**
   * 尝试重连
   */
  private attemptReconnect(): void {
    if (
      this.options.maxReconnectAttempts !== 0 &&
      this.reconnectAttempts < this.options.maxReconnectAttempts
    ) {
      setTimeout(() => {
        this.reconnectAttempts++
        console.log(`正在尝试第 ${this.reconnectAttempts} 次重连...`)
        this.connect(this.port, this.options)
      }, this.options.reconnectInterval)
    } else {
      console.error('已达到最大重连次数')
    }
  }

  /**
   * 开始心跳检测
   */
  private startHeartbeat(): void {
    if (this.heartbeatTimer) return
    this.heartbeatTimer = window.setInterval(() => {
      if (this.socket?.readyState === WebSocket.OPEN) {
        this.socket.send('ping') // 根据后端协议，可以发送更复杂的心跳包
      }
    }, this.options.heartbeatInterval)
  }

  /**
   * 停止心跳检测
   */
  private stopHeartbeat(): void {
    if (this.heartbeatTimer) {
      clearInterval(this.heartbeatTimer)
      this.heartbeatTimer = null
    }
  }

  /**
   * 发送消息
   * @param data 要发送的数据
   */
  public send(data: unknown): void {
    if (this.socket?.readyState === WebSocket.OPEN) {
      const message =
        typeof data === 'object' && data !== null ? JSON.stringify(data) : String(data)
      this.socket.send(message)
    } else {
      console.error('WebSocket 未连接，无法发送消息')
    }
  }

  /**
   * 注册消息处理函数
   * @param type 消息类型
   * @param handler 处理函数
   */
  public on(type: string, handler: MessageHandler): void {
    this.handlers.set(type, handler)
  }

  /**
   * 移除消息处理函数
   * @param type 消息类型
   */
  public off(type: string): void {
    this.handlers.delete(type)
  }

  /**
   * 触发内部事件（如 open, close）
   * @param event 事件名
   * @param data 数据
   */
  private emit(event: string, data: unknown): void {
    // 这里可以集成事件总线（如 mitt）来向外派发事件
    // 例如：eventBus.emit(event, data);
    console.log(event, data)
  }

  /**
   * 断开连接
   */
  public disconnect(): void {
    if (this.socket) {
      this.socket.close()
    }
  }
}
