import { http } from "@/utils/http";
import { Result, PageResult } from "@/api/common";
//import { Counter } from "@/api/sys/counter.d";

/** 获取计数器管理列表 */
export const getCounterPage = (data?: object) => {
  return http.request<Result<PageResult<Counter>>>(
    "post",
    "/api/sys/counter/page",
    {
      data
    }
  );
};

/** 获取计数器 */
export const getCounter = (data?: object) => {
  return http.request<Result<Counter>>("post", "/api/sys/counter/get", {
    data
  });
};

/** 创建计数器 */
export const createCounter = (data?: object) => {
  return http.request<Result<Counter>>("post", "/api/sys/counter/create", {
    data
  });
};

/** 更新计数器 */
export const updateCounter = (data?: object) => {
  return http.request<Result<Counter>>("post", "/api/sys/counter/update", {
    data
  });
};

/** 删除计数器 */
export const delCounter = (data?: object) => {
  return http.request<Result<Counter>>("post", "/api/sys/counter/del", {
    data
  });
};

//Model

/** 计数器 */
interface Counter {
  /** 关键字 */
  key: string;
  /** 值 */
  seq: number;
}

interface CounterFormItemProps {
  /** 关键字 */
  key: string;
  /** 值 */
  seq: number;
}
interface CounterFormProps {
  formInline: CounterFormItemProps;
}

export type { CounterFormItemProps, CounterFormProps, Counter };
