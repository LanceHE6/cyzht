import { LocalStorageManager } from "@/utils/local-storage.ts";

const WS_BASE_URL = "ws://127.0.0.1:8080/api/v1/user/ws/online"; // WebSocket 服务器地址

export class WebSocketManager {
  private ws: WebSocket | null = null;
  private heartbeatInterval: NodeJS.Timeout | null = null;
  private reconnectTimeout: NodeJS.Timeout | null = null;
  private readonly reconnectDelay: number = 5000; // 重连延迟时间，单位：毫秒
  private readonly heartbeatIntervalTime: number = 5000; // 心跳间隔时间，单位：毫秒
  private localStorage: LocalStorageManager | null = null;

  // 构造函数
  constructor(localStorage: LocalStorageManager) {
    this.localStorage = localStorage;
    this.connect();
  }

  // 连接 WebSocket
  private connect() {
    // 延迟初始化
    const token = this.localStorage?.getToken();

    if (!token) {
      console.warn("No token found, WebSocket connection aborted.");

      return;
    }

    this.ws = new WebSocket(`${WS_BASE_URL}?token=${token}`);

    this.ws.onopen = () => {
      console.log("WebSocket connection opened.");
      this.startHeartbeat();
      if (this.reconnectTimeout) {
        clearTimeout(this.reconnectTimeout);
        this.reconnectTimeout = null;
      }
    };

    this.ws.onmessage = (event) => {
      console.log("WebSocket message received:", event.data);
      // 在这里处理接收到的消息
      this.handleMessage(event.data);
    };

    this.ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    this.ws.onclose = (event) => {
      console.warn("WebSocket connection closed:", event);
      this.stopHeartbeat();
      this.reconnect();
    };
  }

  private handleMessage(message: string) {
    // 根据消息内容执行不同的操作
    if (message === "ping") {
      // 心跳响应
      console.log("Heartbeat received.");
    } else {
      // 其他业务消息
      // ...
    }
  }

  private startHeartbeat() {
    this.heartbeatInterval = setInterval(() => {
      if (this.ws && this.ws.readyState === WebSocket.OPEN) {
        this.ws.send("pong");
        console.log("Heartbeat sent.");
      }
    }, this.heartbeatIntervalTime);
  }

  private stopHeartbeat() {
    if (this.heartbeatInterval) {
      clearInterval(this.heartbeatInterval);
      this.heartbeatInterval = null;
    }
  }

  private reconnect() {
    if (this.reconnectTimeout) return;

    console.log(`WebSocket reconnecting in ${this.reconnectDelay / 1000}s...`);
    this.reconnectTimeout = setTimeout(() => {
      this.connect();
    }, this.reconnectDelay);
  }

  public send(message: string) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(message);
    } else {
      console.error("WebSocket connection not open, message not sent.");
    }
  }

  public close() {
    if (this.ws) {
      console.log("WebSocket closed");
      this.ws.close();
    }
  }
}
