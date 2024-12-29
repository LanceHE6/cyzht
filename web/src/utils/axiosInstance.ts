import { useNavigate } from "react-router-dom";
import axios from "axios";

import { getToken, removeToken, removeUser } from "./localStorage.js";

const base_url = "http://localhost:8080";

// 创建一个 Axios 实例
export const axiosInstance = axios.create({
  baseURL: base_url, // 设置基本URL
  timeout: 5000, // 设置超时时间
});
// 创建一个有鉴权头的Axios实例
export const axiosInstanceWithAuth = axios.create({
  baseURL: base_url,
  timeout: 5000,
  headers: {
    Authorization: `Bearer ` + getToken(),
  },
});

// 添加响应拦截器
axios.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 登录过期
    if (error.response.status === 401) {
      removeUser();
      removeToken();
      const navigate = useNavigate();

      navigate("/login");
    }
    console.log("请求出错了: " + error.response.status);

    return Promise.reject(error);
  },
);
