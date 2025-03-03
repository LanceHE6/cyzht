import { LocalStorageManager } from "@/utils/local-storage.ts";
import { WebSocketManager } from "@/utils/websocket.ts";
import { ToastProvider } from "@/utils/toast.ts";

export const LocalStorage: LocalStorageManager = new LocalStorageManager();
export const WebSocketClient: WebSocketManager = new WebSocketManager(
  LocalStorage,
);
export const Toast: ToastProvider = new ToastProvider();
