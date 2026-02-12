import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { SysSms } from "@/api/sys/sys-sms.d";

//Api 

/** 获取短信管理列表 */
export const getSysSmsPage = (data?: object) => {
  return http.request<Result<PageResult<SysSms>>>(
    "post",
    "/api/sys/sys-sms/page",
    {
      data
    }
  );
};

/** 获取短信 */
export const getSysSms = (data?: object) => {
  return http.request<Result<SysSms>>("post", "/api/sys/sys-sms/get", {
    data
  });
};

/** 创建短信 */
export const createSysSms = (data?: object) => {
  return http.request<Result<SysSms>>("post", "/api/sys/sys-sms/create", {
    data
  });
};

/** 更新短信 */
export const updateSysSms = (data?: object) => {
  return http.request<Result<SysSms>>("post", "/api/sys/sys-sms/update",{
    data
  });
};

/** 删除短信 */
export const delSysSms = (data?: object) => {
  return http.request<Result<SysSms>>("post", "/api/sys/sys-sms/del",{
    data
  });
};

//Model

/** 短信 */
interface SysSms {
  
  /** 主键编码 */
  id: number;
  /** 手机号 */
  phone: string;
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


interface SysSmsFormItemProps {
  
  /** 主键编码 */
  id: number;
  /** 手机号 */
  phone: string;
  /** 验证码 */
  code: string;
  /** 类型 */
  type: string;
  /** 状态 */
  status: number;
  /** 使用状态 */
  useStatus: number;
}
interface SysSmsFormProps {
  formInline: SysSmsFormItemProps;
}

export type { SysSmsFormItemProps, SysSmsFormProps, SysSms };