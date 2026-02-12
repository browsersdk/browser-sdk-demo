import { http } from "@/utils/http";
import { Result } from "@/api/common";

// type Result = {
//   success: boolean;
//   data: Array<any>;
// };

export const getAsyncRoutes = () => {
  return http.request<Result<Array<any>>>("post", "/api/sys/user/menus");
};
