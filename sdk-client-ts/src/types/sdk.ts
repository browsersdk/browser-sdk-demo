export interface IResponse {
  /** 状态码 */
  code: number
  /** 错误或提示信息 */
  msg: string
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
