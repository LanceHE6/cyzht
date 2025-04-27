import { Card } from "@heroui/react";
import { useState } from "react";

import Chat from "@/components/activity/chat.tsx";
import Member from "@/components/activity/member.tsx";
import Exhibitor from "@/components/activity/exhibitor.tsx";

export const Activity = ({
  aid,
  setIsMenuVisible,
}: {
  aid: string;
  setIsMenuVisible: React.Dispatch<React.SetStateAction<boolean>>;
}) => {
  const [selectedExhibitorId, setSelectedExhibitorId] = useState<string | null>(
    null,
  );
  // 新增移动端视图状态，默认显示展商列表
  const [mobileView, setMobileView] = useState<
    "exhibitors" | "chat" | "member"
  >("exhibitors");

  // 处理展商选择
  const handleExhibitorSelect = (id: string | null) => {
    console.log("Selected Exhibitor ID:", id);
    setSelectedExhibitorId(id);
    // 移动端选择展商后切换到聊天界面
    // xs断点为480
    if (window.innerWidth < 480) {
      setMobileView("chat");
      // 并隐藏侧边栏
      setIsMenuVisible(false);
    }
  };

  // 返回展商列表
  const handleBackToExhibitors = () => {
    setMobileView("exhibitors");
    // 设置状态栏可见
    setIsMenuVisible(true);
  };
  // 返回聊天界面
  const handleBackToChat = () => {
    setMobileView("chat");
    // 设置状态栏不可见
    setIsMenuVisible(false);
  };

  return (
    <div className="grid xs:grid-cols-15 grid-cols-7 w-full h-full">
      {/* 桌面端视图 - 始终显示所有组件 */}
      {/* 展商列表 */}
      <Card
        className="xs:col-span-4 col-span-7 items-center border-b-blue-500 border-0 bg-[#F6FFFF] xs:block hidden"
        radius="none"
      >
        <Exhibitor aid={aid} onExhibitorSelect={handleExhibitorSelect} />
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
        />
      </Card>

      {/* 成员列表 */}
      <Card
        className="xs:col-span-2 col-span-7 items-center border-b-blue-500 border-0 bg-[#FFF8FF] xs:block hidden"
        radius="none"
      >
        <Member aid={aid} eid={selectedExhibitorId} onBack={handleBackToChat} />
      </Card>

      {/* 移动端视图 - 根据状态条件渲染 */}
      {/* 展商列表视图 */}
      {mobileView === "exhibitors" && (
        <Card
          className="col-span-7 items-center border-b-blue-500 border-0 bg-[#F6FFFF] xs:hidden block"
          radius="none"
        >
          <Exhibitor aid={aid} onExhibitorSelect={handleExhibitorSelect} />
        </Card>
      )}

      {/* 聊天界面视图 */}
      {mobileView === "chat" && (
        <Card
          className="col-span-7 h-full bg-[#FFF6FF] flex-col xs:hidden block"
          radius="none"
        >
          {/* 传递 isMenuVisible 状态到 Chat 组件 */}
          <Chat
            aid={aid}
            eid={selectedExhibitorId}
            setMobileView={setMobileView}
            onBack={handleBackToExhibitors}
          />
        </Card>
      )}

      {/* 成员界面视图 */}
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
};
