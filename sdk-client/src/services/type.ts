import type { AuthTokens } from '@/utils/tokenManager';

export interface LoginRequest {
    code: string;
    devId: string;
    inviteCode: string;
    name: string;
    os: string;
    osVer: string;
    password: string;
    src: string;
    tid: string;
    username: string;
    uuid: string;
    ver: string;
}

export interface LoginResponse {
    code: number;
    data: AuthTokens;
    msg: string;
    reqId: string;
}

export interface UserInfo {
    id: number;
    username: string;
    phone: string;
    email: string;
    nickname: string;
    name: string;
    status: number;
    createdAt: string;
    updatedAt: string;
}

export interface UserInfoResponse {
    reqId: string;
    code: number;
    msg: string;
    data: UserInfo;
}

// Browser相关接口定义
export interface Geographic {
    accuracy: string;
    enable: number;
    latitude: string;
    longitude: string;
    useip: number;
}

export interface Browser {
    audioContext: number;
    bluetooth: number;
    canvas: number;
    cpu: number;
    customerId: string;
    deviceName: string;
    doNotTrack: number;
    dpi: string;
    enableCookie: number;
    enableScanPort: number;
    enablenotice: number;
    enableopen: number;
    enablepic: number;
    enablesound: number;
    enablevideo: number;
    envId: number;
    envName: string;
    fontList: string[];
    geographic: Geographic;
    hardware: number;
    ignoreCookieErr: number;
    ipChannel: string;
    kernel: string;
    kernelVersion: string;
    language: string[];
    mac: string;
    mediaDevice: number;
    mem: number;
    picsize: string;
    proxy: string;
    publicIp: string;
    remark: string;
    scanPort: number[];
    serial: string;
    speechVoices: number;
    system: string;
    uaVersion: string;
    userAgent: string;
    webGl: number;
    webRTC: number;
    webRTCIP: string;
    zone: string;
    id: number;
    name: string;
    userId: number;
    data: string;
    createdAt: string;
    updatedAt: string;
    /** 1: 停止, 3: 启动 */
    status: number;
}

export interface BrowserDto {
    id?: number;
    name: string;
    envId?: number;
    userId: number;
    data?: string;
    /** 1: 停止, 3: 启动 */
    status?: number;
}

export interface BrowserGetPageReq {
    page: number;
    size: number;
    name?: string;
    envId?: number;
    userId?: number;
}

export interface PageResponse<T> {
    list: T[];
    total: number;
    page: number;
    size: number;
}

export interface BaseResponse<T> {
    code: number;
    data: T;
    msg: string;
    reqId: string;
}

export interface ReqId {
    id: number;
}

export interface ReqIds {
    ids: number[];
}