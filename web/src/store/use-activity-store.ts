import { create } from "zustand";

interface UnreadState {
  unreadCounts: Record<string, number>;
  setUnreadCount: (key: string, count: number) => void;
  getUnreadCount: (key: string) => number;
}
// 创建一个zustand store用于存储未读消息数
export const useActivityStore = create<UnreadState>((set, get) => ({
  unreadCounts: {},

  setUnreadCount: (key: string, count: number) =>
    set((state) => ({
      unreadCounts: {
        ...state.unreadCounts,
        [key]: count,
      },
    })),

  getUnreadCount: (key: string) => {
    return get().unreadCounts[key] || 0;
  },
}));
