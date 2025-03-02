// base64加密存储token

import { getUserInfo } from "@/utils/user-utils.ts";

export class LocalStorageManager {
  constructor() {}

  private readonly TOKEN_KEY: string = window.btoa("token");
  private readonly USER_KEY: string = "user";

  // 将token字段加密存储
  public setToken(token: string) {
    localStorage.setItem(this.TOKEN_KEY, window.btoa(token));
  }

  // 获取base64解密后的token
  public getToken() {
    if (!localStorage.getItem(this.TOKEN_KEY)) {
      return "";
    }

    return window.atob(<string>localStorage.getItem(this.TOKEN_KEY));
  }

  public removeToken() {
    localStorage.removeItem(this.TOKEN_KEY);
  }
  // 设置 user
  public setUser(user: object) {
    const jsonUser = JSON.stringify(user);

    localStorage.setItem(this.USER_KEY, window.btoa(jsonUser));
  }

  // 获取 user
  public getUser(): object {
    const jsonUser = localStorage.getItem(this.USER_KEY);
    const user = jsonUser ? window.atob(jsonUser) : "";

    return JSON.parse(user);
  }
  public async updateUser() {
    const user = await getUserInfo();

    console.log("user:", user);
    if (
      window.btoa(user) !==
        window.atob(<string>localStorage.getItem(this.USER_KEY)) &&
      user !== null
    ) {
      localStorage.setItem(this.USER_KEY, window.btoa(user));
    }

    return window.atob(<string>localStorage.getItem(this.USER_KEY));
  }

  // 删除 user
  public removeUser() {
    localStorage.removeItem(this.USER_KEY);
  }
}
