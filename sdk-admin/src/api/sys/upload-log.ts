import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { UploadLog } from "@/api/sys/upload-log.d";

/** 获取UploadLog管理列表 */
export const getUploadLogPage = (data?: object) => {
  return http.request<Result<PageResult<UploadLog>>>(
    "post",
    "/api/sys/upload-log/page",
    {
      data
    }
  );
};

/** 获取UploadLog */
export const getUploadLog = (data?: object) => {
  return http.request<Result<UploadLog>>("post", "/api/sys/upload-log/get", {
    data
  });
};

/** 创建UploadLog */
export const createUploadLog = (data?: object) => {
  return http.request<Result<UploadLog>>("post", "/api/sys/upload-log/create", {
    data
  });
};

/** 更新UploadLog */
export const updateUploadLog = (data?: object) => {
  return http.request<Result<UploadLog>>("post", "/api/sys/upload-log/update", {
    data
  });
};

/** 删除UploadLog */
export const delUploadLog = (data?: object) => {
  return http.request<Result<UploadLog>>("post", "/api/sys/upload-log/del", {
    data
  });
};

//Model

/** UploadLog */
interface UploadLog {
  /** id */
  id: number;
  /** 资源地址 */
  url: string;
  /** 提供方 */
  provider: string;
  /** 1使用中 -1 删除 */
  status: number;
  /** 同步状态 */
  syncStatus: number;
  /** 创建时间 */
  createdAt: Date;
  /** 更新时间 */
  updatedAt: Date;
}

interface UploadLogFormItemProps {
  /** id */
  id: number;
  /** 资源地址 */
  url: string;
  /** 提供方 */
  provider: string;
  /** 1使用中 -1 删除 */
  status: number;
  /** 同步状态 */
  syncStatus: number;
}
interface UploadLogFormProps {
  formInline: UploadLogFormItemProps;
}

export type { UploadLogFormItemProps, UploadLogFormProps, UploadLog };
