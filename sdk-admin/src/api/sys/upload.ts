import { http } from "@/utils/http";
import { Result } from "@/api/common";

export const uploadFile = (data?: object) => {
  return http.request<Result<string>>(
    "post",
    "/api/sys/cfg/upload",
    {
      data
    },
    {
      headers: {
        "Content-Type": "multipart/form-data"
      }
    }
  );
};
