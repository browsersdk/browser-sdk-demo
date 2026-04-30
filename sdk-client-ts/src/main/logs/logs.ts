import { app } from 'electron'
import path from 'path'
import log from 'electron-log'

export default class Logger {
  /** 生产环境 */
  #logger: log.MainLogger
  static #instance
  constructor() {
    this.#logger = log.create({ logId: 'app' })
    // 'error' | 'warn' | 'info' | 'verbose' | 'debug' | 'silly';
    this.#logger.transports.console.level = app.isPackaged ? 'info' : 'debug'
    this.#logger.transports.console.format = '%c{h}:{i}:{s}.{ms}{scope} [{level}]%c {text}'
    this.#logger.transports.file.level = 'debug'
    this.#logger.transports.file.format = '[{y}-{m}-{d} {h}:{i}:{s}] [{level}] {text}'
    this.#logger.transports.file.resolvePathFn = () => this.#setLogDirByEnv()

    this.#logger.transports.console.useStyles = true
    // 开启颜色（通过 styles 配置级别对应的颜色）
    this.#logger.transports.console.colorMap = {
      silly: 'grey',
      debug: 'gray',
      verbose: 'blue',
      info: 'green',
      warn: 'yellow',
      error: 'red'
    }
    // 日志文件大小限制（超过则分割），默认 10485760 字节（10MB）
    this.#logger.transports.file.maxSize = 5 * 1024 * 1024 // 5MB
    // this.#logger.transports.file.maxFiles = 30 // 保留30天
  }
  static getInstance(): log.MainLogger {
    if (!Logger.#instance) {
      const wrapper = new Logger()
      Logger.#instance = wrapper.getLog()
    }
    return Logger.#instance
  }
  getLog(): log.MainLogger {
    return this.#logger
  }
  /**
   * 区分环境设置日志目录
   * - 开发环境：项目根目录/logs
   * - 生产环境：userData/logs
   */
  #setLogDirByEnv(): string {
    if (!app.isPackaged) {
      // 开发环境：项目根目录（向上找，适配 main 进程文件在 src/main 下的情况）
      // __dirname 是当前文件所在目录，根据你的项目结构调整 ../ 的层数
      // return path.join(app.getAppPath(), 'logs')
      return path.join(app.getAppPath(), 'logs', this.#getDailyFileName())
    } else {
      // 生产环境：系统用户数据目录（有写入权限）
      // return path.join(process.resourcesPath, 'logs', this.#getDailyFileName())
      return path.join(app.getPath('userData'), 'logs', this.#getDailyFileName())
    }
  }
  /**
   * 生成按天命名的 fileName
   * @returns 文件名
   */
  #getDailyFileName = (): string => {
    const date = new Date()
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')
    return `app-${year}${month}${day}.log`
  }
}
