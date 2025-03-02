import { axiosInstanceWithAuth } from "@/utils/axios-instance.ts";

// 获取用户信息
export const getUserInfo = async () => {
  const resp = await axiosInstanceWithAuth.get("/api/v1/user/info");
  console.log("response:", resp);

  if (resp.status === 200) {
    return resp.data.data;
  }

  return null;
};
