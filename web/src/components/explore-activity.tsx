// explore-activity.tsx
import { Input } from "@heroui/input";
import {
  ScrollShadow,
  Spacer,
} from "@heroui/react";
import { useEffect, useState } from "react";

import { SearchIcon } from "@/components/icons.tsx";
import { axiosInstance } from "@/utils/axios-instance.ts";
import { Toast } from "@/utils/utils.ts";
import { ActivityCard } from "@/components/activity-card.tsx"; // 引入新创建的组件

export const ExploreActivity = () => {
  // 展会数据
  const [activities, setActivities] = useState<any[]>([]);
  const [completedActivities, setCompletedActivities] = useState<any[]>([]);
  const [keyword, setKeyword] = useState("");
  const [debouncedKeyword, setDebouncedKeyword] = useState("");

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedKeyword(keyword);
    }, 2000);

    return () => clearTimeout(handler);
  }, [keyword]);

  useEffect(() => {
    if (debouncedKeyword !== "") {
      searchActivity(true, debouncedKeyword).then((data) => {
        setActivities(data);
      });
      searchActivity(false, debouncedKeyword).then((data) => {
        setCompletedActivities(data);
      });
    } else {
      searchActivity(true).then((data) => {
        setActivities(data);
      });
      searchActivity(false).then((data) => {
        setCompletedActivities(data);
      });
    }
  }, [debouncedKeyword]);

  const searchActivity = async (
    isInProgress: boolean = true,
    keyword: string = "",
  ) => {
    const response = await axiosInstance.get("/api/v1/activity/search", {
      params: {
        page: -1,
        page_size: 10,
        is_in_progress: isInProgress,
        keyword: keyword,
      },
    });

    console.log(response);
    if (response.data.code === 0) {
      return response.data.data.rows;
    }
    Toast.danger("获取展会失败", response.data.msg);

    return [];
  };

  return (
    <ScrollShadow hideScrollBar className="">
      <div className="w-full h-full flex flex-col items-center justify-start pt-20">
        <div className="w-full max-w-md mb-10">
          <Input
            isClearable
            aria-label="Search"
            classNames={{
              inputWrapper: "bg-default-100",
              input: "text-sm",
            }}
            labelPlacement="outside"
            placeholder="搜索你感兴趣的展会"
            size="lg"
            startContent={
              <SearchIcon className="text-base text-default-400 pointer-events-none flex-shrink-0" />
            }
            type="search"
            value={keyword}
            onValueChange={(value) => {
              setKeyword(value);
            }}
          />
        </div>
        <Spacer y={20} />
        <div className="w-full px-10">
          <h2 className="font-bold text-[32px] text-left">进行中</h2>
        </div>
        <Spacer y={10} />
        <div className="w-full grid grid-cols-4 gap-4 px-10">
          {activities.map((activity, index) => (
            <ActivityCard key={index} activity={activity} />
          ))}
        </div>
        <Spacer y={36} />
        <div className="w-full px-10">
          <h2 className="font-bold text-[32px] text-left">已结束</h2>
        </div>
        <Spacer y={10} />
        <div className="w-full grid grid-cols-4 gap-4 px-10">
          {completedActivities.map((activity, index) => (
            <ActivityCard key={index} activity={activity} />
          ))}
        </div>
      </div>
    </ScrollShadow>
  );
};
