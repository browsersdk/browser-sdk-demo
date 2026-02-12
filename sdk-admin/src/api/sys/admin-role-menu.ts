import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { AdminRoleMenu } from "@/api/sys/admin-role-menu.d";

/** 获取角色菜单管理列表 */
export const getAdminRoleMenuPage = (data?: object) => {
  return http.request<Result<PageResult<AdminRoleMenu>>>(
    "post",
    "/api/sys/admin-role-menu/page",
    {
      data
    }
  );
};

/** 获取角色菜单 */
export const getAdminRoleMenu = (data?: object) => {
  return http.request<Result<AdminRoleMenu>>(
    "post",
    "/api/sys/admin-role-menu/get",
    {
      data
    }
  );
};

/** 创建角色菜单 */
export const createAdminRoleMenu = (data?: object) => {
  return http.request<Result<AdminRoleMenu>>(
    "post",
    "/api/sys/admin-role-menu/create",
    {
      data
    }
  );
};

/** 更新角色菜单 */
export const updateAdminRoleMenu = (data?: object) => {
  return http.request<Result<AdminRoleMenu>>(
    "post",
    "/api/sys/admin-role-menu/update",
    {
      data
    }
  );
};

/** 删除角色菜单 */
export const delAdminRoleMenu = (data?: object) => {
  return http.request<Result<AdminRoleMenu>>(
    "post",
    "/api/sys/admin-role-menu/del",
    {
      data
    }
  );
};

//Model

/** 角色菜单 */
interface AdminRoleMenu {
  /** 主键 */
  roleId: number;
  /** 主键 */
  menuId: number;
}

interface AdminRoleMenuFormItemProps {
  /** 主键 */
  roleId: number;
  /** 主键 */
  menuId: number;
}
interface AdminRoleMenuFormProps {
  formInline: AdminRoleMenuFormItemProps;
}

export type {AdminRoleMenuFormItemProps, AdminRoleMenuFormProps, AdminRoleMenu};
