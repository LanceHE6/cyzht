import { useNavigate } from "react-router-dom";
import axios from "axios";

import { LocalStorage, Toast, WebSocketClient } from "@/utils/utils.ts";

export const BASE_URL = "http://localhost:8080";

// 创建一个 Axios 实例
export const axiosInstance = axios.create({
  baseURL: BASE_URL, // 设置基本URL
  timeout: 5000, // 设置超时时间
});
// 创建一个有鉴权头的Axios实例
export const axiosInstanceWithAuth = axios.create({
  baseURL: BASE_URL,
  timeout: 5000,
  headers: {
    Authorization: `Bearer ` + LocalStorage.getToken(),
  },
});

// 添加响应拦截器
axiosInstanceWithAuth.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 登录过期
    if (error.response.status === 401) {
      LocalStorage.removeToken();
      LocalStorage.removeUser();
      WebSocketClient.close();
      const navigate = useNavigate();

      Toast.danger("登录过期，请重新登录", error.response.data.msg);

      navigate("/login");
    }

    return Promise.reject(error);
  },
);
