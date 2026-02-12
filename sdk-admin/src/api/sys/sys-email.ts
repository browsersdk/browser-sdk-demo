import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { SysEmail } from "@/api/sys/sys-email.d";

//Api 

/** 获取邮件管理列表 */
export const getSysEmailPage = (data?: object) => {
  return http.request<Result<PageResult<SysEmail>>>(
    "post",
    "/api/sys/sys-email/page",
    {
      data
    }
  );
};

/** 获取邮件 */
export const getSysEmail = (data?: object) => {
  return http.request<Result<SysEmail>>("post", "/api/sys/sys-email/get", {
    data
  });
};

/** 创建邮件 */
export const createSysEmail = (data?: object) => {
  return http.request<Result<SysEmail>>("post", "/api/sys/sys-email/create", {
    data
  });
};

/** 更新邮件 */
export const updateSysEmail = (data?: object) => {
  return http.request<Result<SysEmail>>("post", "/api/sys/sys-email/update",{
    data
  });
};

/** 删除邮件 */
export const delSysEmail = (data?: object) => {
  return http.request<Result<SysEmail>>("post", "/api/sys/sys-email/del",{
    data
  });
};

//Model

/** 邮件 */
interface SysEmail {
  
  /** 主键编码 */
  id: number;
  /** 邮箱地址 */
  email: string;
  /** 验证码 */
  code: string;
  /** 类型 */
  type: string;
  /** 状态 */
  status: number;
  /** 使用状态 */
  useStatus: number;
  /** 创建时间 */
  createdAt: number;
  /** 最后更新时间 */
  updatedAt: number;
}


interface SysEmailFormItemProps {
  
  /** 主键编码 */
  id: number;
  /** 邮箱地址 */
  email: string;
  /** 验证码 */
  code: string;
  /** 类型 */
  type: string;
  /** 状态 */
  status: number;
  /** 使用状态 */
  useStatus: number;
}
interface SysEmailFormProps {
  formInline: SysEmailFormItemProps;
}

export type { SysEmailFormItemProps, SysEmailFormProps, SysEmail };