export interface IResponse {
  /** 状态码 */
  code: number
  /** 错误或提示信息 */
  msg: string
}

export interface LoginRequest {
  /** 状态码 */
  code: string
  /** 设备 ID（客户端生成，用于标识设备） */
  devId: string
  /** 邀请码 */
  inviteCode: string
  /** 设备或客户端名称 */
  name: string
  /** 操作系统名称（如 'Windows', 'Android'） */
  os: string
  /** 操作系统版本（如 '10', '14.4'） */
  osVer: string
  /** 登录密码（若使用密码登录） */
  password: string
  /** 请求来源标识（如 web/app） */
  src: string
  /** 事务或跟踪 ID，可用于调试 */
  tid: string
  /** 登录用户名（账号） */
  username: string
  /** 设备唯一标识 UUID */
  uuid: string
  /** 客户端版本号 */
  ver: string
}

export interface AuthTokens {
  accessToken: string
  refreshToken: string
  expires: string
  refreshExpire: string
  username: string
  need: number
}

export interface LoginResponse {
  /** 状态码 */
  code: number
  /** 登录成功后返回的认证令牌信息 */
  data: AuthTokens
  /** 错误或提示信息 */
  msg: string
  /** 请求唯一 ID，便于排查问题 */
  reqId: string
}

/**
 * UserSig响应数据中的 data 字段类型
 */
export interface IUserSigData {
  /** 用户签名过期时间戳（Unix 时间，秒） */
  expireTime: number
  /** 用户签名，用于身份验证 */
  userSig: string
}

export interface UserInfo {
  /** 用户 ID（数据库主键） */
  id: number
  /** 登录账号/用户名 */
  username: string
  /** 手机号（可能为空字符串） */
  phone: string
  /** 邮箱地址（可能为空字符串） */
  email: string
  /** 昵称 */
  nickname: string
  /** 真实姓名或展示名 */
  name: string
  /** 账号状态（数值表示，具体含义由后端定义） */
  status: number
  /** 账号创建时间，ISO 字符串 */
  createdAt: string
  /** 最后更新时间，ISO 字符串 */
  updatedAt: string
}

export interface UserInfoResponse {
  /** 请求唯一 ID */
  reqId: string
  /** 返回状态码 */
  code: number
  /** 信息或错误消息 */
  msg: string
  /** 用户信息数据 */
  data: UserInfo
}

// Browser相关接口定义
export interface Geographic {
  /** 定位精度（如米），以字符串形式返回 */
  accuracy: string
  /** 是否启用地理位置（0/1） */
  enable: number
  /** 纬度，字符串格式 */
  latitude: string
  /** 经度，字符串格式 */
  longitude: string
  /** 是否使用 IP 定位（0/1） */
  useip: number
}

/**
 * 字体配置接口
 */
interface FontConfig {
  /** 是否启用字体指纹，默认0 */
  enable: number
  /** 字体列表，默认['Arial', 'Helvetica', 'Times New Roman'] */
  list: string[]
}

/**
 * 浏览器指纹信息接口
 * 包含各类浏览器、设备、系统相关的指纹特征
 */
interface FingerInfo {
  /** 音频上下文配置，默认0 */
  audioContext?: number
  /** 电池信息，默认0 */
  battery?: number
  /** 蓝牙配置，默认0 */
  bluetooth?: number
  /** Canvas指纹配置，默认4 */
  canvas?: number
  /** 客户端矩形信息，默认0 */
  clientRects?: number
  /** CPU核心数，默认4 */
  cpu?: number
  /** 设备名称 */
  deviceName?: string
  /** 禁止跟踪配置，默认0 */
  doNotTrack?: number
  /** 屏幕DPI，默认1920x1080 */
  dpi?: string
  /** 启用垃圾回收，默认0 */
  enableGc?: number
  /** 启用通知，默认1 */
  enablenotice?: number
  /** 启用打开操作，默认1 */
  enableopen?: number
  /** 启用打开数量限制，默认0 */
  enableOpenNumber?: number
  /** 启用图片，默认1 */
  enablepic?: number
  /** 启用端口扫描，默认0 */
  enableScanPort?: number
  /** 启用声音，默认1 */
  enablesound?: number
  /** 启用视频，默认1 */
  enablevideo?: number
  /** 字体配置信息 */
  font?: FontConfig
  /** 字体指纹配置，默认0 */
  fontFinger?: number
  /** 垃圾回收时间，默认0 */
  gcTime?: number
  /** 地理位置信息配置 */
  geographic?: Geographic
  /** 硬件信息配置，默认0 */
  hardware?: number
  /** 内核名称 */
  kernel: string
  /** 内核版本 */
  kernelVersion: string
  /** 浏览器语言列表，默认['zh-CN', 'en-US'] */
  language?: string[]
  /** MAC地址 */
  mac?: string
  /** 媒体设备配置，默认0 */
  mediaDevice?: number
  /** 内存大小(GB)，默认8 */
  mem?: number
  /** 图片尺寸，默认1920x1080 */
  picSize?: string
  /** 扫描端口列表，默认空 */
  scanPort?: string
  /** 语音合成配置，默认1 */
  speechVoices?: number
  /** 操作系统信息 */
  system: string
  /** 用户代理字符串 */
  ua?: string
  /** 用户代理版本，默认134 */
  uaVersion?: string
  /** UI语言列表，默认空数组 */
  uilanguage?: string[]
  /** WebGL配置，默认0 */
  webGl?: number
  /** WebGL信息配置，默认0 */
  webGlInfo?: number
  /** WebGL渲染器，默认空字符串 */
  webGlRenderer?: string
  /** WebGL供应商，默认空字符串 */
  webGlVendor?: string
  /** WebRTC配置，默认5 */
  webRTC?: number
  /** WebRTC IP地址，默认127.0.0.1 */
  webRTCIP?: string
  /** 窗口尺寸，默认1920x1080 */
  widowssize?: string // 注意：原字段拼写可能是笔误，正确应为windowSize
  /** 时区/区域配置 */
  zone?: string
}
/**
 * 浏览器环境配置数据接口
 * 包含环境信息和浏览器指纹信息
 */
interface BrowserData {
  /** 环境ID */
  envId: string
  /** 环境名称 */
  envName: string
  /** 桥接代理，默认空字符串 */
  bridgeProxy?: string
  /** Cookie信息，默认空字符串 */
  cookie?: string
  /** 客户ID */
  customerId: string
  /** IP渠道 */
  ipChannel: string
  /** 公网IP */
  publicIp: string
  /** 代理配置 */
  proxy: string
  /** 地区信息，默认空字符串 */
  region?: string
  /** 备注信息 */
  remark: string
  /** 序列号 */
  serial: string
  /** 浏览器指纹信息 */
  finger: FingerInfo
}

export interface Browser {
  /** 是否支持/启用 AudioContext（1/0） */
  audioContext: number
  /** 是否支持蓝牙（1/0） */
  bluetooth: number
  /** canvas 支持情况（1/0） */
  canvas: number
  /** CPU 信息（数值或等级） */
  cpu: number
  /** 客户标识 ID（如果有） */
  customerId: string
  /** 设备名称（如 'Chrome' 或设备型号） */
  deviceName: string
  /** doNotTrack 标志（1/0 或其它约定值） */
  doNotTrack: number
  /** 屏幕或像素密度描述 */
  dpi: string
  /** 是否启用 Cookie（1/0） */
  enableCookie: number
  /** 是否启用端口扫描（1/0） */
  enableScanPort: number
  /** 是否允许通知（1/0） */
  enablenotice: number
  /** 是否允许打开新窗口/外部链接（1/0） */
  enableopen: number
  /** 是否允许图片加载/处理（1/0） */
  enablepic: number
  /** 是否启用声音（1/0） */
  enablesound: number
  /** 是否启用视频（1/0） */
  enablevideo: number
  /** 环境 ID（如测试/生产） */
  envId: string
  /** 环境名称 */
  envName: string
  /** 字体列表 */
  fontList: string[]
  /** 地理位置信息 */
  geographic: Geographic
  /** 硬件信息等级或标识 */
  hardware: number
  /** 是否忽略 Cookie 错误（1/0） */
  ignoreCookieErr: number
  /** IP 通道或来源 */
  ipChannel: string
  /** 内核类型（如 'Blink', 'WebKit'） */
  kernel: string
  /** 内核版本 */
  kernelVersion: string
  /** 语言列表（优先顺序） */
  language: string[]
  /** MAC 地址（可能被掩码或为空） */
  mac: string
  /** 媒体设备数量或标识 */
  mediaDevice: number
  /** 内存信息（MB 或其它约定） */
  mem: number
  /** 图片大小或分辨率描述 */
  picsize: string
  /** 代理信息（如代理地址） */
  proxy: string
  /** 公网 IP 地址 */
  publicIp: string
  /** 备注说明 */
  remark: string
  /** 扫描端口列表 */
  scanPort: number[]
  /** 序列号或设备标识 */
  serial: string
  /** 语音合成或语音选项数量 */
  speechVoices: number
  /** 操作系统信息 */
  system: string
  /** UA 版本字符串 */
  uaVersion: string
  /** 浏览器 userAgent 字符串 */
  ua: string
  /** WebGL 支持情况（1/0） */
  webGl: number
  /** 是否支持 WebRTC（1/0） */
  webRTC: number
  /** WebRTC IP（候选地址） */
  webRTCIP: string
  /** 时区或区域描述 */
  zone: string
  /** 浏览器记录 ID */
  id: number
  /** 浏览器或设备名称 */
  name: string
  /** 关联的用户 ID */
  userId: number
  /** 附加数据，通常为 JSON 字符串 */
  data: string
  /** 创建时间，ISO 字符串 */
  createdAt: string
  /** 更新时间，ISO 字符串 */
  updatedAt: string
  /** 1: 停止, 3: 启动 */
  status: number
}

export interface BrowserDto {
  /** 浏览器记录 ID（用于更新时） */
  id?: number
  /** 浏览器或设备名称 */
  envName: string
  /** 环境 ID */
  envId?: string
  /** 关联用户 ID */
  userId: number
  /** 额外数据（通常为 JSON） */
  data?: BrowserData
  /** 创建时间，ISO 字符串 */
  createdAt?: string
  /** 更新时间，ISO 字符串 */
  updatedAt?: string
  /** 1: 停止, 3: 启动 */
  status?: number
}

export interface BrowserGetPageReq {
  /** 页码，从 1 开始 */
  page: number
  /** 每页大小 */
  size: number
  /** 按名称过滤 */
  envName?: string
  /** 按环境 ID 过滤 */
  envId?: string
  /** 按用户 ID 过滤 */
  userId?: number
}

export interface PageResponse<T> {
  /** 当前页的数据列表 */
  list: T[]
  /** 总条目数 */
  total: number
  /** 当前页码 */
  page: number
  /** 每页大小 */
  size: number
}

export interface BaseResponse<T> {
  /** 返回状态码 */
  code: number
  /** 返回的数据体 */
  data: T
  /** 返回消息或错误描述 */
  msg: string
  /** 请求 ID，用于跟踪 */
  reqId: string
}

export interface ReqId {
  /** 单个 ID（常用于删除或查询） */
  id: number
}

export interface ReqIds {
  /** ID 列表（批量操作） */
  ids: number[]
}
