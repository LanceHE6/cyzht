import { Card } from "@heroui/react";
import React, { useState } from "react";

import Chat from "@/components/activity/chat.tsx";
import Member from "@/components/activity/member.tsx";
import Exhibitor from "@/components/activity/exhibitor.tsx";

export interface ActivityProps {
  aid: string;
}

export const Activity: React.FC<ActivityProps> = (props: ActivityProps) => {
  let { aid } = props;
  const [selectedExhibitorId, setSelectedExhibitorId] = useState<string | null>(
    null,
  );

  return (
    <div className="grid grid-cols-15 w-full h-full">
      <Card
        className="col-span-4 items-center border-b-blue-500 border-0 bg-[#F6FFFF]"
        radius="none"
      >
        <Exhibitor
          aid={aid}
          onExhibitorSelect={setSelectedExhibitorId} // 传递回调函数,将选中 Exhibitor 的 ID 传递给父组件
        />
      </Card>
      <Card
        className="col-span-9 h-full bg-[#FFF6FF] flex flex-col"
        radius="none"
      >
        <Chat aid={aid} eid={selectedExhibitorId} />
      </Card>
      <Card
        className="col-span-2 items-center border-b-blue-500 border-0 bg-[#FFF8FF]"
        radius="none"
      >
        <Member aid={aid} eid={selectedExhibitorId} />
      </Card>
    </div>
  );
};
