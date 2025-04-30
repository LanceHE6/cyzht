import React, {
  useState,
  useCallback,
  forwardRef,
  useImperativeHandle,
} from "react";
import { Card } from "@heroui/react";

import Chat from "@/components/activity/chat.tsx";
import Member from "@/components/activity/member.tsx";
import Exhibitor from "@/components/activity/exhibitor.tsx";
import { LocalStorage } from "@/utils/utils.ts";
import { useActivityStore } from "@/store/use-activity-store.ts";

export interface ActivityRef {
  getUnreadCount: (key: string) => number;
  setUnreadCount: (key: string, count: number) => void;
}

const Activity = forwardRef<
  ActivityRef,
  {
    aid: string;
    setIsMenuVisible: React.Dispatch<React.SetStateAction<boolean>>;
  }
>(({ aid, setIsMenuVisible }, ref) => {
  const [selectedExhibitorId, setSelectedExhibitorId] = useState<string | null>(
    null,
  );
  const [mobileView, setMobileView] = useState<
    "exhibitors" | "chat" | "member"
  >("exhibitors");

  // 未读计数状态
  const unreadCounts = useActivityStore((state) => state.unreadCounts);
  const setUnreadCount = useActivityStore((state) => state.setUnreadCount);

  // 提供给外部的方法
  const getUnreadCount = useCallback(
    (key: string): number => unreadCounts[key] || 0,
    [unreadCounts],
  );

  // 暴露方法给父组件（保留 ref 接口兼容性）
  useImperativeHandle(ref, () => ({
    getUnreadCount,
    setUnreadCount: (key: string, count: number) => setUnreadCount(key, count),
  }));

  // 处理展商选择
  const handleExhibitorSelect = (id: string | null) => {
    const exhibitorKey = id || aid;

    setSelectedExhibitorId(id);
    setUnreadCount(exhibitorKey, 0);
    if (window.innerWidth < 480) {
      setMobileView("chat");
      setIsMenuVisible(false);
    }
  };

  // 处理新消息回调
  const handleNewMessage = useCallback(
    (message: any) => {
      const currentUser = LocalStorage.getUser();

      if (message.from_user.id === currentUser?.id) return;

      const messageExhibitorKey = message.exhibitor?.id || aid;
      const currentCount =
        useActivityStore.getState().unreadCounts[messageExhibitorKey] || 0;

      // console.log("key", messageExhibitorKey);
      setUnreadCount(messageExhibitorKey, currentCount + 1);
    },
    [aid, selectedExhibitorId],
  );

  const handleBackToExhibitors = () => {
    setMobileView("exhibitors");
    setIsMenuVisible(true);
  };
  const handleBackToChat = () => {
    setMobileView("chat");
    setIsMenuVisible(false);
  };

  return (
    <div className="grid xs:grid-cols-15 grid-cols-7 w-full h-full">
      {/* 展商列表 */}
      <Card
        className="xs:col-span-4 col-span-7 items-center border-b-blue-500 border-0 bg-[#F6FFFF] xs:block hidden"
        radius="none"
      >
        <Exhibitor
          aid={aid}
          unreadCounts={unreadCounts}
          onExhibitorSelect={handleExhibitorSelect}
        />
      </Card>

      {/* 聊天界面 */}
      <Card
        className="xs:col-span-9 col-span-7 h-full bg-[#FFF6FF] xs:flex flex-col hidden"
        radius="none"
      >
        <Chat
          aid={aid}
          eid={selectedExhibitorId}
          setMobileView={setMobileView}
          onBack={handleBackToExhibitors}
          onNewMessage={handleNewMessage}
        />
      </Card>

      {/* 成员列表 */}
      <Card
        className="xs:col-span-2 col-span-7 items-center border-b-blue-500 border-0 bg-[#FFF8FF] xs:block hidden"
        radius="none"
      >
        <Member aid={aid} eid={selectedExhibitorId} onBack={handleBackToChat} />
      </Card>

      {/* 移动端视图 */}
      {mobileView === "exhibitors" && (
        <Card
          className="col-span-7 items-center border-b-blue-500 border-0 bg-[#F6FFFF] xs:hidden block"
          radius="none"
        >
          <Exhibitor
            aid={aid}
            unreadCounts={unreadCounts}
            onExhibitorSelect={handleExhibitorSelect}
          />
        </Card>
      )}

      {mobileView === "chat" && (
        <Card
          className="col-span-7 h-full bg-[#FFF6FF] flex-col xs:hidden block"
          radius="none"
        >
          <Chat
            aid={aid}
            eid={selectedExhibitorId}
            setMobileView={setMobileView}
            onBack={handleBackToExhibitors}
            onNewMessage={handleNewMessage}
          />
        </Card>
      )}

      {mobileView === "member" && (
        <Card
          className="col-span-7 h-full bg-[#FFF6FF] flex-col xs:hidden block"
          radius="none"
        >
          <Member
            aid={aid}
            eid={selectedExhibitorId}
            onBack={handleBackToChat}
          />
        </Card>
      )}
    </div>
  );
});

Activity.displayName = "Activity";
export default Activity;
