// explore-activity.tsx
import { Input } from "@heroui/input";
import {
  Card,
  CardHeader,
  ScrollShadow,
  Spacer,
  Image,
  CardFooter,
  Button,
  Avatar,
} from "@heroui/react";
import React, { useEffect, useState } from "react";

import { DefaultActivityIcon, SearchIcon } from "@/components/icons.tsx";
import { axiosInstance } from "@/utils/axios-instance.ts";
import { Toast } from "@/utils/utils.ts";
import { formatDate } from "@/utils/datetime.ts";
import Activity from "@/pages/activity.tsx";

interface ActivityCardProps {
  activity: any;
  onAddNewTab: (tab: any) => void;
  setIsMenuVisible: React.Dispatch<React.SetStateAction<boolean>>;
}

// ActivityCard 展会展示卡片
const ActivityCard = ({
  activity,
  onAddNewTab,
  setIsMenuVisible,
}: ActivityCardProps) => {
  const handleViewClick = () => {
    const newTab = {
      key: activity.id,
      title: activity.name,
      icon: activity.icon ? (
        <Avatar
          isBordered
          className="cursor-pointer"
          size="sm"
          src={activity.icon}
        />
      ) : (
        <DefaultActivityIcon />
      ),
      component: (
        <Activity aid={activity.id} setIsMenuVisible={setIsMenuVisible} />
      ),
    };

    onAddNewTab(newTab);
  };

  return (
    <div className="h-full px-4">
      <Card key={activity.id} isFooterBlurred className="py-1 xs:h-5/6 h-5/6">
        <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
          <h4 className="font-bold text-large">{activity.name}</h4>
          <p className="text-tiny uppercase font-bold">
            时间: {formatDate(activity.start_at)} -{" "}
            {formatDate(activity.end_at)}
          </p>
          <Image
            removeWrapper
            alt="Card background"
            className="object-cover rounded-xl h-full"
            src={
              activity.icon
                ? activity.icon
                : "https://heroui.com/images/hero-card.jpeg"
            }
          />
        </CardHeader>
        <Spacer y={2} />

        <CardFooter className="absolute bg-black/40 bottom-0 z-10 border-t-1 border-default-600 dark:border-default-100 h-1/5">
          <div className="flex flex-grow gap-2 items-center">
            <Avatar
              alt="activity icon"
              className="rounded-full w-10 h-11 bg-black"
              src={
                activity.creator.avatar
                  ? activity.creator.avatar
                  : "https://heroui.com/images/hero-card.jpeg"
              }
            />
            <div className="flex flex-col">
              <p className="text-tiny text-white/60">
                {activity.creator.nickname}
              </p>
            </div>
          </div>
          <Button radius="full" size="sm" onPress={handleViewClick}>
            查看
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
};

interface ExploreActivityProps {
  onAddNewTab: (tab: any) => void;
  setIsMenuVisible: React.Dispatch<React.SetStateAction<boolean>>;
}
export const ExploreActivity = ({
  onAddNewTab,
  setIsMenuVisible,
}: ExploreActivityProps) => {
  // 展会数据
  const [notStartedActivities, setNotStartedActivities] = useState<any[]>([]);
  const [inProgressActivities, setInProgressActivities] = useState<any[]>([]);
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
      searchActivity(0, debouncedKeyword).then((data) => {
        setNotStartedActivities(data);
      });
      searchActivity(1, debouncedKeyword).then((data) => {
        setInProgressActivities(data);
      });
      searchActivity(2, debouncedKeyword).then((data) => {
        setCompletedActivities(data);
      });
    } else {
      searchActivity(0).then((data) => {
        setNotStartedActivities(data);
      });
      searchActivity(1).then((data) => {
        setInProgressActivities(data);
      });
      searchActivity(2).then((data) => {
        setCompletedActivities(data);
      });
    }
  }, [debouncedKeyword]);

  const searchActivity = async (status: number = 1, keyword: string = "") => {
    const response = await axiosInstance.get("/api/v1/activity/search", {
      params: {
        page: -1,
        page_size: 10,
        status: status,
        keyword: keyword,
      },
    });

    if (response.data.code === 0) {
      return response.data.data.rows;
    }
    Toast.danger("获取展会失败", response.data.msg);

    return [];
  };

  return (
    <ScrollShadow hideScrollBar className="">
      <div className="w-full h-full flex flex-col items-center justify-start pt-20">
        <div className="xs:w-full w-9/12 max-w-md mb-10">
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

        {/* 未开始活动 */}
        {notStartedActivities.length > 0 && (
          <>
            <Spacer y={20} />
            <div className="w-full px-10">
              <h2 className="font-bold text-[32px] text-left">未开始</h2>
            </div>
            <Spacer y={10} />
            <div className="w-full grid xs:grid-cols-6 gap-4 xs:px-10 px-6">
              {notStartedActivities.map((activity, index) => (
                <ActivityCard
                  key={index}
                  activity={activity}
                  setIsMenuVisible={setIsMenuVisible}
                  onAddNewTab={onAddNewTab}
                />
              ))}
            </div>
          </>
        )}

        {/* 进行中活动 */}
        {inProgressActivities.length > 0 && (
          <>
            <Spacer y={20} />
            <div className="w-full px-10">
              <h2 className="font-bold text-[32px] text-left">进行中</h2>
            </div>
            <Spacer y={10} />
            <div className="w-full grid xs:grid-cols-6 gap-4 xs:px-10 px-6">
              {inProgressActivities.map((activity, index) => (
                <ActivityCard
                  key={index}
                  activity={activity}
                  setIsMenuVisible={setIsMenuVisible}
                  onAddNewTab={onAddNewTab}
                />
              ))}
            </div>
          </>
        )}

        {/* 已结束活动 */}
        {completedActivities.length > 0 && (
          <>
            <Spacer y={20} />
            <div className="w-full px-10">
              <h2 className="font-bold text-[32px] text-left">已结束</h2>
            </div>
            <Spacer y={10} />
            <div className="w-full grid xs:grid-cols-6 xs:px-10 px-6">
              {completedActivities.map((activity, index) => (
                <ActivityCard
                  key={index}
                  activity={activity}
                  setIsMenuVisible={setIsMenuVisible}
                  onAddNewTab={onAddNewTab}
                />
              ))}
            </div>
          </>
        )}

        {/* 当三种状态都没有数据时显示“暂无数据” */}
        {!notStartedActivities.length &&
          !inProgressActivities.length &&
          !completedActivities.length && (
            <div className="flex flex-col items-center justify-center w-full h-full">
              <Spacer y={20} />
              <p className="text-gray-500 text-[32px] font-bold">暂无数据</p>
              <Spacer y={20} />
            </div>
          )}
      </div>
    </ScrollShadow>
  );
};
