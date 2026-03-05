export interface IOpenEnv {
    envId: string
    args?: string[]
    urls?: string[]
}
export interface IOpenParams {
    envs: IOpenEnv[]
}
export interface ICloseParams {
    envs: string[]
}
