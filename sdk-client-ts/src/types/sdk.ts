export interface IResponse {
  /** 状态码 */
  code: number
  /** 错误或提示信息 */
  msg: string
}
export interface IOpenEnv {
  envId: string
  args?: string[]
  urls?: string[]
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
