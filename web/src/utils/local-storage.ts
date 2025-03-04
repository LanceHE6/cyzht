// base64加密存储token

import { getUserInfo } from "@/utils/user-utils.ts";

export class LocalStorageManager {
  constructor() {}

  private readonly IsEncrypted: boolean = false; // 是否加密存储
  private readonly TOKEN_KEY: string = this.IsEncrypted
    ? window.btoa("token")
    : "token";
  private readonly USER_KEY: string = this.IsEncrypted
    ? window.btoa("user")
    : "user";

  // 将token字段加密存储
  public setToken(token: string) {
    if (!this.IsEncrypted) {
      localStorage.setItem(this.TOKEN_KEY, token);
    } else {
      // base64加密存储
      localStorage.setItem(this.TOKEN_KEY, window.btoa(token));
    }
  }

  // 获取base64解密后的token
  public getToken() {
    if (!localStorage.getItem(this.TOKEN_KEY)) {
      return "";
    }
    if (!this.IsEncrypted) {
      return localStorage.getItem(this.TOKEN_KEY);
    } else {
      return window.atob(<string>localStorage.getItem(this.TOKEN_KEY));
    }
  }

  public removeToken() {
    localStorage.removeItem(this.TOKEN_KEY);
  }
  // 设置 user
  public setUser(user: object) {
    const jsonUser = JSON.stringify(user);

    if (!this.IsEncrypted) {
      localStorage.setItem(this.USER_KEY, jsonUser);
    } else {
      localStorage.setItem(this.USER_KEY, window.btoa(jsonUser));
    }
  }

  // 获取 user
  public getUser(): object | null {
    const jsonUser = localStorage.getItem(this.USER_KEY);

    if (jsonUser === null) {
      return null;
    }
    if (!this.IsEncrypted) {
      return JSON.parse(jsonUser);
    } else {
      return JSON.parse(window.atob(jsonUser));
    }
  }
  public async updateUser() {
    const user = await getUserInfo();

    console.log("user:", user);
    if (window.btoa(user) !== JSON.stringify(this.getUser()) && user !== null) {
      this.setUser(user);
    }

    return this.getUser();
  }

  // 删除 user
  public removeUser() {
    localStorage.removeItem(this.USER_KEY);
  }
}
