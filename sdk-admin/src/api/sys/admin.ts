import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { Admin } from "@/api/sys/admin.d";

//Api

/** 获取用户管理列表 */
export const getAdminPage = (data?: object) => {
  return http.request<Result<PageResult<Admin>>>(
    "post",
    "/api/sys/admin/page",
    {
      data
    }
  );
};

/** 获取用户 */
export const getAdmin = (data?: object) => {
  return http.request<Result<Admin>>("post", "/api/sys/admin/get", {
    data
  });
};

/** 创建用户 */
export const createAdmin = (data?: object) => {
  return http.request<Result<Admin>>("post", "/api/sys/admin/create", {
    data
  });
};

/** 更新用户 */
export const updateAdmin = (data?: object) => {
  return http.request<Result<Admin>>("post", "/api/sys/admin/update", {
    data
  });
};

/** 删除用户 */
export const delAdmin = (data?: object) => {
  return http.request<Result<Admin>>("post", "/api/sys/admin/del", {
    data
  });
};

//Model

/** 用户 */
interface Admin {
  /** 主键 */
  id: number;
  /** 用户名 */
  username: string;
  /** 手机号 */
  phone: string;
  /** 邮箱 */
  email: string;
  /** 密码 */
  password: string;
  /** 昵称 */
  nickname: string;
  /** 姓名 */
  name: string;
  /** 头像 */
  avatar: string;
  /** 签名 */
  bio: string;
  /** 生日 格式 yyyy-MM-dd */
  birthday: Date;
  /** 性别 1男 2女 3未知 */
  gender: number;
  /** 团队id */
  roleId: number;
  /** 部门id */
  deptId: number;
  /** 备注 */
  remark: string;
  /** 锁定结束时间 */
  lockTime: Date;
  /** 状态 1正常  */
  status: number;
  /** 客户端id */
  clientId: string;
  /** 注册ip */
  regIp: string;
  /**  */
  ipLocation: string;
  /** 更新者 */
  updateBy: number;
  /** 创建时间 */
  createdAt: Date;
  /** 最后更新时间 */
  updatedAt: Date;
}

interface AdminFormItemProps {
  /** 主键 */
  id: number;
  /** 用户名 */
  username: string;
  /** 手机号 */
  phone: string;
  /** 邮箱 */
  email: string;
  /** 密码 */
  password: string;
  /** 昵称 */
  nickname: string;
  /** 姓名 */
  name: string;
  /** 头像 */
  avatar: string;
  /** 签名 */
  bio: string;
  /** 生日 格式 yyyy-MM-dd */
  birthday: Date;
  /** 性别 1男 2女 3未知 */
  gender: number;
  /** 团队id */
  roleId: number;
  /** 部门id */
  deptId: number;
  /** 备注 */
  remark: string;
  /** 锁定结束时间 */
  lockTime: Date;
  /** 状态 1正常  */
  status: number;
  /** 客户端id */
  clientId: string;
  /** 注册ip */
  regIp: string;
  /**  */
  ipLocation: string;
  higherDeptOptions: any;
}
interface AdminFormProps {
  formInline: AdminFormItemProps;
}

export type { AdminFormItemProps, AdminFormProps, Admin };