import { Card } from "@heroui/react";
import React from "react";

import Chat from "@/components/activity/chat.tsx";
import Member from "@/components/activity/member.tsx";
import Exhibitor from "@/components/activity/exhibitor.tsx";

export interface ActivityProps {
  aid: string;
}

export const Activity: React.FC<ActivityProps> = (props: ActivityProps) => {
  let { aid } = props;

  console.log(aid);

  return (
    <div className="grid grid-cols-15 w-full h-full">
      <Card
        className="col-span-4 items-center border-b-blue-500 border-0 bg-[#F6FFFF]"
        radius="none"
      >
        <Exhibitor aid={aid} />
      </Card>
      <Card
        className="col-span-9 h-full bg-[#FFF6FF] flex flex-col"
        radius="none"
      >
        <Chat aid={aid} eid={"0"} />
      </Card>
      <Card
        className="col-span-2 items-center border-b-blue-500 border-0 bg-[#FFF8FF]"
        radius="none"
      >
        <Member aid={aid} />
      </Card>
    </div>
  );
};
