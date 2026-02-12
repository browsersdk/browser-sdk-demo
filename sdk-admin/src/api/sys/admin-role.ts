import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { AdminRole } from "@/api/sys/admin-role.d";

/** 获取角色管理列表 */
export const getAdminRolePage = (data?: object) => {
  return http.request<Result<PageResult<AdminRole>>>(
    "post",
    "/api/sys/admin-role/page",
    {
      data
    }
  );
};

export const getAdminRoleList = (data?: object) => {
  return http.request<Result<Array<AdminRole>>>(
    "post",
    "/api/sys/admin-role/list",
    {
      data
    }
  );
};

/** 获取角色 */
export const getAdminRole = (data?: object) => {
  return http.request<Result<AdminRole>>("post", "/api/sys/admin-role/get", {
    data
  });
};

/** 创建角色 */
export const createAdminRole = (data?: object) => {
  return http.request<Result<AdminRole>>("post", "/api/sys/admin-role/create", {
    data
  });
};

/** 更新角色 */
export const updateAdminRole = (data?: object) => {
  return http.request<Result<AdminRole>>("post", "/api/sys/admin-role/update",{
    data
  });
};

/** 删除角色 */
export const delAdminRole = (data?: object) => {
  return http.request<Result<AdminRole>>("post", "/api/sys/admin-role/del",{
    data
  });
};

//Model

/** 角色 */
interface AdminRole {
  /** 主键 */
  id: number;
  /** 角色名称 */
  name: string;
  /** 角色代码 */
  roleKey: string;
  /** 排序 */
  roleSort: number;
  /** 状态 */
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

interface AdminRoleFormItemProps {
  /** 主键 */
  id: number;
  /** 角色名称 */
  name: string;
  /** 角色代码 */
  roleKey: string;
  /** 排序 */
  roleSort: number;
  /** 状态 */
  status: number;
  /** 备注 */
  remark: string;
}
interface AdminRoleFormProps {
  formInline: AdminRoleFormItemProps;
}

export type { AdminRoleFormItemProps, AdminRoleFormProps, AdminRole };
