import axios from "axios";

import { LocalStorage, Toast, WebSocketClient } from "@/utils/utils.ts";

export const BASE_URL = "http://localhost:8080";

// 创建一个 Axios 实例
export const axiosInstance = axios.create({
  baseURL: BASE_URL, // 设置基本URL
  timeout: 5000, // 设置超时时间
});

axiosInstance.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error) {
    Toast.danger("请求失败", error.response.data.msg);

    return Promise.reject(error);
  },
);

// 创建一个有鉴权头的Axios实例
export const axiosInstanceWithAuth = axios.create({
  baseURL: BASE_URL,
  timeout: 5000,
});

// 导出一个函数来设置导航回调
export const setNavigateCallback = (
  navigateCallback: (path: string) => void,
) => {
  // 添加请求拦截器
  axiosInstanceWithAuth.interceptors.request.use(
    function (config) {
      const token = LocalStorage.getToken();

      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      } else {
        Toast.warning("请先登录", "");
        if (navigateCallback) {
          navigateCallback("/login");
        }
      }

      return config;
    },
    function (error) {
      return Promise.reject(error);
    },
  );

  // 添加响应拦截器
  axiosInstanceWithAuth.interceptors.response.use(
    function (response) {
      return response;
    },
    function (error) {
      console.log("error", error);
      if (error.response.status === 401) {
        LocalStorage.removeToken();
        LocalStorage.removeUser();
        WebSocketClient.close();
        if (navigateCallback) {
          navigateCallback("/login");
        }
        Toast.danger("登录过期，请重新登录", error.response.data.msg);
      }

      return Promise.reject(error);
    },
  );
};
