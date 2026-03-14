export interface SocketOptions {
    /** 心跳间隔（毫秒） */
    heartbeatInterval?: number;
    /** 重连间隔（毫秒） */
    reconnectInterval?: number;
    /** 最大重连次数 */
    maxReconnectAttempts?: number;
}

/** socket 操作事件 */
export enum OperateEvent {
    /** 打开成功 */
    OPEN_SUCCESS = 'browser-open-success',
    /** 打开失败 */
    OPEN_FAILED = 'browser-open-failed',
    /** 关闭成功 */
    CLOSE_SUCCESS = 'browser-close-success',
    /** 关闭失败 */
    CLOSE_FAILED = 'browser-close-failed',
}