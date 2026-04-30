export interface IResponse {
  /** 状态码 */
  code: number
  /** 错误或提示信息 */
  msg: string
}

export interface IBrowserInfoResponse {
  /** 状态码 */
  code: number;
  /** 响应类型标识 */
  type: string;
  /** 数据主体 */
  data: {
    /** 环境列表 */
    envs: Array<{
      /** 环境 ID（数值型） */
      envId: string;
    }>;
  };
}

export interface IOpenCookie {
  domain: string
  expirationDate: number
  hostOnly: boolean
  httpOnly: boolean
  name: string
  path: string
  sameSite: string
  secure: boolean
  session: boolean
  storeId?: string
  value: string
}
export interface IOpenEnv {
  envId: string
  args?: string[]
  urls?: string[]
  cookies?: IOpenCookie[]
}
export interface IAppBindParams {
  port: number
  usersin: string
}
export interface IOpenParams {
  envs: IOpenEnv[]
}
export interface ICloseParams {
  envs: string[]
}
