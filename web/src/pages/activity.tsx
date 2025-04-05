import { Card } from "@heroui/react";
import React from "react";

import Chat from "@/components/activity/chat.tsx";
import Member from "@/components/activity/member.tsx";

export interface ActivityProps {
  aid: string;
}

export const Activity: React.FC<ActivityProps> = (props: ActivityProps) => {
  let { aid } = props;

  console.log(aid);

  return (
    <div className="grid grid-cols-15 w-full h-full">
      <Card className="col-span-4 items-center border-b-blue-500 border-0 bg-[#F6FFFF]">
        {/* 参展商列表 */}
        <h3>参展商列表</h3>
        {/* 参展商列表内容 */}
      </Card>
      <Chat aid={aid} />
      <Member />
    </div>
  );
};
