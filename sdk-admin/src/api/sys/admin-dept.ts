import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { AdminDept } from "@/api/sys/admin-dept.d";

/** 获取部门管理列表 */
export const getAdminDeptPage = (data?: object) => {
  return http.request<Result<PageResult<AdminDept>>>(
    "post",
    "/api/sys/admin-dept/page",
    {
      data
    }
  );
};

export const getAdminDeptTree = (data?: object) => {
  return http.request<Result<AdminDeptFormItemProps>>(
    "post",
    "/api/sys/admin-dept/all",
    {
      data
    }
  );
};

/** 获取部门 */
export const getAdminDept = (data?: object) => {
  return http.request<Result<AdminDept>>("post", "/api/sys/admin-dept/get", {
    data
  });
};

/** 创建部门 */
export const createAdminDept = (data?: object) => {
  return http.request<Result<AdminDept>>("post", "/api/sys/admin-dept/create", {
    data
  });
};

/** 更新部门 */
export const updateAdminDept = (data?: object) => {
  return http.request<Result<AdminDept>>("post", "/api/sys/admin-dept/update", {
    data
  });
};

/** 删除部门 */
export const delAdminDept = (data?: object) => {
  return http.request<Result<AdminDept>>("post", "/api/sys/admin-dept/del", {
    data
  });
};

//Model

/** 部门 */
interface AdminDept {
  /** 主键 */
  id: number;
  /** 父id */
  parentId: number;
  /** 部门路径 */
  deptPath: string;
  /** 部门名 */
  name: string;
  /** 类型 1分公司 2部门 */
  type: number;
  /** 部门领导 */
  principal: string;
  /** 手机号 */
  phone: string;
  /** 邮箱 */
  email: string;
  /** 排序 */
  sort: number;
  /** 状态 1正常 2关闭 */
  status: number;
  /** 备注 */
  remark: string;
  /** 创建者 */
  createBy: number;
  /** 更新者 */
  updateBy: number;
  /** 创建时间 */
  createdAt: Date;
  /** 最后更新时间 */
  updatedAt: Date;
  /** 删除时间 */
  deletedAt: Date;
}

interface AdminDeptFormItemProps {
  /** 主键 */
  id: number;
  /** 父id */
  parentId: number;
  /** 部门路径 */
  deptPath: string;
  /** 部门名 */
  name: string;
  /** 类型 1分公司 2部门 */
  type: number;
  /** 部门领导 */
  principal: string;
  /** 手机号 */
  phone: string;
  /** 邮箱 */
  email: string;
  /** 排序 */
  sort: number;
  /** 状态 1正常 2关闭 */
  status: number;
  /** 备注 */
  remark: string;
}
interface AdminDeptFormProps {
  formInline: AdminDeptFormItemProps;
}

export type { AdminDeptFormItemProps, AdminDeptFormProps, AdminDept };
