import {
  Accordion,
  AccordionItem,
  Button,
  Card,
  CardFooter,
  CardHeader,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
  Image,
  Spinner,
  Listbox,
  ListboxItem,
  Divider,
} from "@heroui/react";
import { useEffect, useState } from "react";

import { axiosInstanceWithAuth } from "@/utils/axios-instance.ts";
import { SelectionIcon } from "@/components/icons.tsx";
import { Toast } from "@/utils/utils.ts";

interface ExhibitorProps {
  aid: string;
  onExhibitorSelect?: (exhibitorId: string | null) => void; // 添加回调函数类型
}

interface Exhibitor {
  id: string;
  name: string;
  introduce: string;
  creator: {
    nickname: string;
    avatar: string;
  } | null;
}

interface JoinedExhibitor {
  id: string;
  exhibitor: Exhibitor;
  role: number;
}

const Exhibitor: React.FC<ExhibitorProps> = ({ aid, onExhibitorSelect }) => {
  const [activity, setActivity] = useState<{
    name?: string;
    icon?: string;
  }>({});
  const [allExhibitors, setAllExhibitors] = useState<Exhibitor[]>([]);
  const [joinedExhibitors, setJoinedExhibitors] = useState<JoinedExhibitor[]>(
    [],
  );
  const [loading, setLoading] = useState({
    activity: true,
    exhibitors: true,
    joined: true,
  });

  useEffect(() => {
    const fetchData = async () => {
      try {
        // 获取展会信息
        const activityRes = await axiosInstanceWithAuth(
          `/api/v1/activity/search`,
          {
            params: { id: aid },
          },
        );

        setActivity(activityRes.data.data.rows[0]);
        setLoading((prev) => ({ ...prev, activity: false }));

        // 获取所有参展商
        const exhibitorsRes = await axiosInstanceWithAuth(
          `/api/v1/exhibitor/${aid}/list`,
        );

        setAllExhibitors(exhibitorsRes.data.data.rows);
        setLoading((prev) => ({ ...prev, exhibitors: false }));

        // 获取已加入的参展商
        const joinedRes = await axiosInstanceWithAuth(
          `/api/v1/exhibitor/${aid}/joined`,
        );

        setJoinedExhibitors(joinedRes.data.data.rows);
        setLoading((prev) => ({ ...prev, joined: false }));
      } catch (err) {
        console.error("加载失败", err);
        setLoading({
          activity: false,
          exhibitors: false,
          joined: false,
        });
      }
    };

    fetchData();
  }, [aid]);

  // 判断是否为已加入的参展商
  const isJoined = (exhibitorId: string) => {
    return joinedExhibitors.some((item) => item.exhibitor.id === exhibitorId);
  };

  // 分离已加入和未加入的参展商
  const joinedList = allExhibitors
    .filter((exhibitor: Exhibitor) => isJoined(exhibitor.id)) // 先筛选出已加入的参展商
    .map((exhibitor: Exhibitor) => {
      // 再对筛选结果进行映射
      const joinedInfo = joinedExhibitors.find(
        (item) => item.exhibitor.id === exhibitor.id,
      );

      return {
        ...exhibitor,
        role: joinedInfo?.role || 0,
      };
    });

  const unjoinedList = allExhibitors.filter(
    (exhibitor: Exhibitor) => !isJoined(exhibitor.id),
  );

  const isLoading = loading.activity || loading.exhibitors || loading.joined;

  const handleJoinClick = async (exhibitorId: string) => {
    try {
      await axiosInstanceWithAuth.post(`/api/v1/exhibitor/${exhibitorId}/join`);
      Toast.success("加入参展商成功", null);
      setJoinedExhibitors((prev) => [
        ...prev,
        {
          id: exhibitorId,
          exhibitor: allExhibitors.find(
            (exhibitor) => exhibitor.id === exhibitorId,
          ) || {
            id: "",
            name: "",
            introduce: "",
            creator: null,
          },
          role: 1,
        },
      ]);
    } catch (error) {
      console.error("加入参展商失败", error);
    }
  };

  const handleWithDrawExhibitor = async (exhibitorId: string) => {
    try {
      await axiosInstanceWithAuth.post(
        `/api/v1/exhibitor/${exhibitorId}/withdraw`,
      );
      Toast.success("退出参展商成功", null);
      setJoinedExhibitors((prev) =>
        prev.filter((item) => item.exhibitor.id !== exhibitorId),
      );
    } catch (error) {
      console.error("退出参展商失败", error);
    }
  };

  const handleExhibitorClick = (exhibitorId: string | null) => {
    if (onExhibitorSelect) {
      onExhibitorSelect(exhibitorId);
    }
  };

  return (
    <>
      <Card isFooterBlurred className="w-full h-1/4" radius="none">
        <CardHeader className="absolute z-10 top-1 flex-col items-start" />
        <Image
          removeWrapper
          alt="Activity Icon"
          className="z-0 w-full h-full scale-125 -translate-y-6 object-cover"
          src={activity.icon}
        />
        <CardFooter className="absolute bg-white/30 bottom-0 border-t-1 border-zinc-100/50 z-10 justify-between">
          <Dropdown className="w-full" placement="bottom-end">
            <DropdownTrigger>
              <Button
                className="w-full relative"
                endContent={
                  <div className="absolute right-2">
                    <SelectionIcon size={20} />
                  </div>
                }
                radius="sm"
                variant="shadow"
              >
                <span className="mx-auto">
                  {activity?.name || "加载展会..."}
                </span>
              </Button>
            </DropdownTrigger>
            <DropdownMenu aria-label="Static Actions">
              <DropdownItem key="new">申请参展商</DropdownItem>
              <DropdownItem key="copy">展会详情</DropdownItem>
              <DropdownItem key="delete" className="text-danger" color="danger">
                结束展会/退出展会
              </DropdownItem>
            </DropdownMenu>
          </Dropdown>
        </CardFooter>
      </Card>
      <Divider className="my-0" />
      <Listbox>
        <ListboxItem
          className="flex items-center justify-between"
          textValue="展会大厅"
          onPress={() => handleExhibitorClick(null)}
        >
          <div className="flex items-center w-full">
            <p className="font-medium text-medium truncate">展会大厅</p>
          </div>
        </ListboxItem>
      </Listbox>
      <Divider className="my-0" />
      <Accordion selectionMode="multiple" variant="light">
        <AccordionItem
          key="joined"
          aria-label="Joined exhibitors"
          title={`已加入 (${joinedList.length})`}
        >
          {isLoading ? (
            <div className="flex justify-center p-4">
              <Spinner size="lg" />
            </div>
          ) : (
            <div className="space-y-3">
              <Listbox>
                {joinedList.length > 0 ? (
                  joinedList.map((item) => (
                    <ListboxItem
                      key={item.id}
                      className="flex items-center justify-between"
                      textValue={item.name}
                      onPress={() => handleExhibitorClick(item.id)}
                    >
                      <div className="flex items-center justify-between w-full">
                        <div className="flex-1 min-w-0 max-w-[75%]">
                          {/* 限制文本宽度 */}
                          <p className="font-medium text-medium truncate">
                            {/* 添加truncate */}# {item.name}
                          </p>
                          <p className="text-sm text-gray-500 truncate">
                            {/* 添加truncate */}
                            {item.introduce}
                          </p>
                        </div>
                        <div className="flex-shrink-0 ml-2">
                          {/* 确保Dropdown靠右 */}
                          <Dropdown placement="bottom-end">
                            <DropdownTrigger>
                              <Button
                                isIconOnly
                                endContent={<SelectionIcon size={20} />}
                                radius="sm"
                                variant="light"
                              />
                            </DropdownTrigger>
                            <DropdownMenu aria-label="Static Actions">
                              <DropdownItem key="copy">详情</DropdownItem>
                              <DropdownItem
                                key="delete"
                                className="text-danger"
                                color="danger"
                                onPress={() => handleWithDrawExhibitor(item.id)}
                              >
                                退出
                              </DropdownItem>
                            </DropdownMenu>
                          </Dropdown>
                        </div>
                      </div>
                    </ListboxItem>
                  ))
                ) : (
                  <ListboxItem isDisabled variant="light">
                    <p className="text-center text-gray-500">暂未加入参展商</p>
                  </ListboxItem>
                )}
              </Listbox>
            </div>
          )}
        </AccordionItem>

        <AccordionItem
          key="unjoined"
          aria-label="Unjoined exhibitors"
          title={`未加入 (${unjoinedList.length})`}
        >
          {isLoading ? (
            <div className="flex justify-center">
              <Spinner size="lg" />
            </div>
          ) : (
            <Listbox>
              {unjoinedList.length > 0 ? (
                unjoinedList.map((item) => (
                  <ListboxItem
                    key={item.id}
                    className="flex items-center justify-between"
                    textValue={item.name}
                  >
                    <div className="flex items-center justify-between w-full">
                      <div className="flex-1 min-w-0 max-w-[75%]">
                        {/* 限制文本宽度 */}
                        <p className="font-medium text-medium truncate">
                          {/* 添加truncate */}# {item.name}
                        </p>
                        <p className="text-sm text-gray-500 truncate">
                          {/* 添加truncate */}
                          {item.introduce}
                        </p>
                      </div>
                      <div className="flex-shrink-0 ml-2">
                        {/* 确保Dropdown靠右 */}
                        <Dropdown placement="bottom-end">
                          <DropdownTrigger>
                            <Button
                              isIconOnly
                              endContent={<SelectionIcon size={20} />}
                              radius="sm"
                              variant="light"
                            />
                          </DropdownTrigger>
                          <DropdownMenu aria-label="Static Actions">
                            <DropdownItem
                              key="join"
                              className="text-primary"
                              color="primary"
                              onPress={() => handleJoinClick(item.id)}
                            >
                              申请加入
                            </DropdownItem>
                            <DropdownItem key="info">详情</DropdownItem>
                          </DropdownMenu>
                        </Dropdown>
                      </div>
                    </div>
                  </ListboxItem>
                ))
              ) : (
                <ListboxItem isDisabled>
                  <div className="px-4 py-2 text-center text-gray-500">
                    暂无未加入的参展商
                  </div>
                </ListboxItem>
              )}
            </Listbox>
          )}
        </AccordionItem>
      </Accordion>
    </>
  );
};

export default Exhibitor;
