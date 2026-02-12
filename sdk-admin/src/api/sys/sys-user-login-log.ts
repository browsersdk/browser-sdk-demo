import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { SysUserLoginLog } from "@/api/sys/user-login-log.d";


/** 获取SysUserLoginLog管理列表 */
export const getSysUserLoginLogPage = (data?: object) => {
  return http.request<Result<PageResult<SysUserLoginLog>>>(
    "post",
    "/api/sys/user-login-log/page",
    {
      data
    }
  );
};

/** 获取SysUserLoginLog */
export const getSysUserLoginLog = (data?: object) => {
  return http.request<Result<SysUserLoginLog>>("post",
    "/api/sys/user-login-log/get",
    {
      data
    }
  );
};

/** 创建SysUserLoginLog */
export const createSysUserLoginLog = (data?: object) => {
  return http.request<Result<SysUserLoginLog>>("post", "/api/sys/user-login-log/create", {
    data
  });
};

/** 更新SysUserLoginLog */
export const updateSysUserLoginLog = (data?: object) => {
  return http.request<Result<SysUserLoginLog>>("post", "/api/sys/user-login-log/update",{
    data
  });
};

/** 删除SysUserLoginLog */
export const delSysUserLoginLog = (data?: object) => {
  return http.request<Result<SysUserLoginLog>>("post", "/api/sys/user-login-log/del",{
    data
  });
};

//Model

/** SysUserLoginLog */
interface SysUserLoginLog {
  /** id */
  id: number;
  /** 用户id */
  uid: number;
  /** 登录方式 */
  method: number;
  /** 登录ip */
  ip: string;
  /** 地域 */
  region: string;
  /** 客户端 */
  clientId: string;
  /** 操作系统 */
  clientVer: string;
  /** 操作系统 */
  os: string;
  /** 操作系统版本 */
  osVer: string;
  /** 更新时间 */
  updatedAt: Date;
}

interface SysUserLoginLogFormItemProps {
  /** id */
  id: number;
  /** 用户id */
  uid: number;
  /** 登录方式 */
  method: number;
  /** 登录ip */
  ip: string;
  /** 地域 */
  region: string;
  /** 客户端 */
  clientId: string;
  /** 操作系统 */
  clientVer: string;
  /** 操作系统 */
  os: string;
  /** 操作系统版本 */
  osVer: string;
}
interface SysUserLoginLogFormProps {
  formInline: SysUserLoginLogFormItemProps;
}

export type { SysUserLoginLogFormItemProps, SysUserLoginLogFormProps, SysUserLoginLog };
