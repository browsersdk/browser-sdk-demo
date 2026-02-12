import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";

//Api

/** 获取浏览器管理列表 */
export const getBrowserPage = (data?: object) => {
  return http.request<Result<PageResult<Browser>>>(
    "post",
    "/api/sys/browser/page",
    {
      data
    }
  );
};

/** 获取浏览器详情 */
export const getBrowser = (data?: object) => {
  return http.request<Result<Browser>>("post", "/api/sys/browser/get", {
    data
  });
};

/** 创建浏览器 */
export const createBrowser = (data?: object) => {
  return http.request<Result<Browser>>("post", "/api/sys/browser/create", {
    data
  });
};

/** 更新浏览器 */
export const updateBrowser = (data?: object) => {
  return http.request<Result<Browser>>("post", "/api/sys/browser/update", {
    data
  });
};

/** 删除浏览器 */
export const delBrowser = (data?: object) => {
  return http.request<Result<Browser>>("post", "/api/sys/browser/del", {
    data
  });
};

//Model

/** 浏览器 */
interface Browser {
  /** 主键 */
  id: number;
  /** 名称 */
  name: string;
  /** 环境ID */
  envId: number;
  /** 用户ID */
  userId: number;
  /** 数据 */
  data: string;
  /** 创建时间 */
  createdAt: string;
  /** 最后更新时间 */
  updatedAt: string;
}

interface BrowserFormItemProps {
  /** 主键 */
  id: number;
  /** 名称 */
  name: string;
  /** 环境ID */
  envId: number;
  /** 用户ID */
  userId: number;
  /** 数据 */
  data: string;
}

interface BrowserFormProps {
  formInline: BrowserFormItemProps;
}

export type { BrowserFormItemProps, BrowserFormProps, Browser };