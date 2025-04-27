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
  useDisclosure,
  Chip,
} from "@heroui/react";
import { useEffect, useState } from "react";

import {
  AddExhibitorModal,
  ActivityInfoModal,
  ConfirmModal,
  NoticeModal,
} from "./exhibitor-modal.tsx"; // 引入抽离的模态框组件

import { axiosInstanceWithAuth } from "@/utils/axios-instance.ts";
import { HallIcon, NoticeIcon, SelectionIcon } from "@/components/icons.tsx";
import { Toast } from "@/utils/utils.ts";
import { LocalStorage } from "@/utils/utils.ts"; // 假设 LocalStorage 从这里导入

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

// 公告接口定义
export interface Notice {
  id: string;
  title: string;
  content: string;
  created_at: string;
  creator: {
    id: string;
    nickname: string;
    avatar: string;
  };
}

const Exhibitor: React.FC<ExhibitorProps> = ({ aid, onExhibitorSelect }) => {
  const [activity, setActivity] = useState<{
    name?: string;
    icon?: string;
    creator?: string | any; // 假设 creator 是用户的 ID
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
  const [isCreator, setIsCreator] = useState(false); // 新增状态来判断是否为创建者
  const [exhibitorCreators, setExhibitorCreators] = useState<{
    [key: string]: boolean;
  }>({}); // 判断是否为参展商创建者

  // 模态框状态
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [isInfoModalOpen, setIsInfoModalOpen] = useState(false); // 新增状态用于控制展会详情模态框

  // 新增状态变量用于管理确认模态框
  const [confirmModal, setConfirmModal] = useState<{
    isOpen: boolean;
    type: "delete" | "withdraw" | null;
    exhibitorId: string | null;
  }>({
    isOpen: false,
    type: null,
    exhibitorId: null,
  });

  //结束展会和退出展会的确认模态框
  const [endActivityModal, setEndActivityModal] = useState(false);
  const [withdrawActivityModal, setWithdrawActivityModal] = useState(false);
  // 新增公告相关状态
  const [notices, setNotices] = useState<Notice[]>([]);

  const {
    isOpen: isNoticeOpen,
    onOpen: onNoticeOpen,
    onClose: onNoticeClose,
  } = useDisclosure();

  // 获取用户角色
  const getUserRole = (): number => {
    if (isCreator) return 3; // 创建者最高权限

    return joinedExhibitors.reduce((max, item) => Math.max(max, item.role), 0);
  };

  // 获取公告列表
  const fetchNotices = async () => {
    try {
      const res = await axiosInstanceWithAuth.get(
        `/api/v1/activity/${aid}/notices`,
      );

      setNotices(res.data.data.rows || []);
    } catch (error) {
      Toast.danger("获取公告失败", "请稍后再试");
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        // 获取当前用户信息
        const currentUser = LocalStorage.getUser();
        const currentUserId = currentUser?.id;

        // 获取展会信息
        const activityRes = await axiosInstanceWithAuth(
          `/api/v1/activity/search`,
          {
            params: { id: aid },
          },
        );

        const activityData = activityRes.data.data.rows[0];

        setActivity(activityData);
        setLoading((prev) => ({ ...prev, activity: false }));

        // 判断当前用户是否是活动创建者
        setIsCreator(currentUserId === activityData.creator.id);

        // 获取所有参展商
        const exhibitorsRes = await axiosInstanceWithAuth(
          `/api/v1/exhibitor/${aid}/list`,
        );

        const exhibitorsData = exhibitorsRes.data.data.rows;

        setAllExhibitors(exhibitorsData);
        setLoading((prev) => ({ ...prev, exhibitors: false }));

        // 获取已加入的参展商
        const joinedRes = await axiosInstanceWithAuth(
          `/api/v1/exhibitor/${aid}/joined`,
        );

        const joinedData = joinedRes.data.data.rows;

        setJoinedExhibitors(joinedData);
        setLoading((prev) => ({ ...prev, joined: false }));

        // 判断当前用户是否是每个参展商的创建者
        const creatorMap: { [key: string]: boolean } = {};

        exhibitorsData.forEach((row: any) => {
          creatorMap[row.id] = currentUserId === row.creator?.id;
        });

        setExhibitorCreators(creatorMap);
      } catch (err) {
        console.error("加载失败", err);
        setLoading({
          activity: false,
          exhibitors: false,
          joined: false,
        });
      }
    };

    fetchData().then(() => {
      console.log("数据加载完成");
    });
    fetchNotices();
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

  // 退出参展商的 onPress 事件
  const handleWithDrawExhibitor = (exhibitorId: string) => {
    setConfirmModal({ isOpen: true, type: "withdraw", exhibitorId });
  };

  // 删除参展商的 onPress 事件
  const handleDeleteExhibitor = (exhibitorId: string) => {
    setConfirmModal({ isOpen: true, type: "delete", exhibitorId });
  };

  const handleExhibitorClick = (exhibitorId: string | null) => {
    if (onExhibitorSelect) {
      onExhibitorSelect(exhibitorId);
    }
  };

  // 添加参展商请求
  const handleSubmitExhibitor = async (name: string, introduce: string) => {
    try {
      const response = await axiosInstanceWithAuth.post(
        `/api/v1/exhibitor/add`,
        {
          aid: aid,
          name,
          introduce,
        },
      );

      if (response.data.code === 0) {
        Toast.success("添加参展商成功", null);
        onClose();
      } else {
        Toast.danger("添加参展商失败", response.data.msg);
      }
    } catch (error) {
      console.error("添加参展商失败", error);
      Toast.danger("添加参展商失败", "请稍后再试");
    }
  };

  // 处理确认删除参展商
  const handleConfirmDelete = async () => {
    if (!confirmModal.exhibitorId) return;

    try {
      await axiosInstanceWithAuth.post(`/api/v1/exhibitor/del`, {
        id: confirmModal.exhibitorId,
      });
      Toast.success("删除参展商成功", null);
      setAllExhibitors((prev) =>
        prev.filter((exhibitor) => exhibitor.id !== confirmModal.exhibitorId),
      );
      setJoinedExhibitors((prev) =>
        prev.filter((item) => item.exhibitor.id !== confirmModal.exhibitorId),
      );
    } catch (error) {
      console.error("删除参展商失败", error);
      Toast.danger("删除参展商失败", "请稍后再试");
    } finally {
      setConfirmModal({ isOpen: false, type: null, exhibitorId: null });
    }
  };

  // 处理确认退出参展商
  const handleConfirmWithdraw = async () => {
    if (!confirmModal.exhibitorId) return;

    try {
      await axiosInstanceWithAuth.post(
        `/api/v1/exhibitor/${confirmModal.exhibitorId}/withdraw`,
      );
      Toast.success("退出参展商成功", null);
      setJoinedExhibitors((prev) =>
        prev.filter((item) => item.exhibitor.id !== confirmModal.exhibitorId),
      );
    } catch (error) {
      console.error("退出参展商失败", error);
      Toast.danger("退出参展商失败", "请稍后再试");
    } finally {
      setConfirmModal({ isOpen: false, type: null, exhibitorId: null });
    }
  };

  // 处理确认结束展会
  const handleConfirmEndActivity = async () => {
    try {
      const response = await axiosInstanceWithAuth.post(
        `/api/v1/activity/del`,
        {
          id: aid,
        },
      );

      if (response.data.code === 0) {
        Toast.success("结束展会成功", null);
        window.location.reload();
        onClose();
      } else {
        Toast.danger("结束展会失败", response.data.msg);
      }
    } catch (error) {
      console.error("结束展会失败", error);
      Toast.danger("结束展会失败", "请稍后再试");
    } finally {
      setEndActivityModal(false); // 关闭确认模态框
    }
  };

  // 处理确认退出展会
  const handleConfirmWithdrawActivity = async () => {
    try {
      await axiosInstanceWithAuth.post(`/api/v1/activity/${aid}/withdraw`);
      Toast.success("退出展会成功", null);
      window.location.reload();
      setJoinedExhibitors((prev) =>
        prev.filter((item) => item.exhibitor.id !== aid),
      );
    } catch (error) {
      console.error("退出展会失败", error);
      Toast.danger("退出展会失败", "请稍后再试");
    } finally {
      setWithdrawActivityModal(false); // 关闭确认模态框
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
              {isCreator ? (
                <>
                  <DropdownItem key="add" onPress={onOpen}>
                    添加参展商
                  </DropdownItem>
                  <DropdownItem
                    key="info"
                    onPress={() => setIsInfoModalOpen(true)}
                  >
                    展会详情
                  </DropdownItem>
                  <DropdownItem
                    key="end"
                    className="text-danger"
                    color="danger"
                    onPress={() => setEndActivityModal(true)}
                  >
                    结束展会
                  </DropdownItem>
                </>
              ) : (
                <>
                  <DropdownItem key="join" onPress={onOpen}>
                    添加参展商
                  </DropdownItem>
                  <DropdownItem
                    key="info"
                    onPress={() => setIsInfoModalOpen(true)}
                  >
                    展会详情
                  </DropdownItem>
                  <DropdownItem
                    key="withdraw"
                    className="text-danger"
                    color="danger"
                    onPress={() => setWithdrawActivityModal(true)}
                  >
                    退出展会
                  </DropdownItem>
                </>
              )}
            </DropdownMenu>
          </Dropdown>
        </CardFooter>
      </Card>
      <Divider className="my-0" />
      <Listbox>
        <ListboxItem
          className="flex items-center"
          startContent={<HallIcon size={32} />}
          textValue="展会大厅"
          onPress={() => handleExhibitorClick(null)}
        >
          <div className="flex items-center w-full">
            <p className="font-medium text-medium truncate">展会大厅</p>
          </div>
        </ListboxItem>
        <ListboxItem
          className="flex items-center"
          endContent={<Chip color="primary">{notices.length}</Chip>}
          startContent={<NoticeIcon size={32} />}
          textValue="公告"
          onPress={onNoticeOpen}
        >
          <div className="flex items-center w-full">
            <p className="font-medium text-medium truncate">公告</p>
          </div>
        </ListboxItem>
      </Listbox>
      {/* 新增公告模态框 */}
      <NoticeModal
        aid={aid}
        isOpen={isNoticeOpen}
        notices={notices}
        userRole={getUserRole()}
        onClose={onNoticeClose}
        onRefresh={fetchNotices}
      />

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
                                onPress={() => {
                                  if (exhibitorCreators[item.id]) {
                                    handleDeleteExhibitor(item.id);
                                  } else {
                                    handleWithDrawExhibitor(item.id);
                                  }
                                }}
                              >
                                {exhibitorCreators[item.id]
                                  ? "删除参展商"
                                  : "退出"}
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
                              加入
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
      {/* 添加参展商模态框 */}
      <AddExhibitorModal
        isOpen={isOpen}
        onClose={onClose}
        onSubmit={handleSubmitExhibitor}
      />
      {/* 展会详情模态框 */}
      <ActivityInfoModal
        activity={activity}
        isOpen={isInfoModalOpen}
        onClose={() => setIsInfoModalOpen(false)}
      />

      {/* 确认删除参展商模态框 */}
      <ConfirmModal
        isOpen={confirmModal.isOpen}
        message={
          confirmModal.type === "delete"
            ? "您确定要删除该参展商吗？"
            : "您确定要退出该参展商吗？"
        }
        title={
          confirmModal.type === "delete" ? "确认删除参展商" : "确认退出参展商"
        }
        onClose={() =>
          setConfirmModal({ isOpen: false, type: null, exhibitorId: null })
        }
        onConfirm={
          confirmModal.type === "delete"
            ? handleConfirmDelete
            : handleConfirmWithdraw
        }
      />

      {/* 确认结束展会模态框 */}
      <ConfirmModal
        isOpen={endActivityModal}
        message="您确定要结束该展会吗？"
        title="确认结束展会"
        onClose={() => setEndActivityModal(false)}
        onConfirm={handleConfirmEndActivity}
      />

      {/* 确认退出展会模态框 */}
      <ConfirmModal
        isOpen={withdrawActivityModal}
        message="您确定要退出该展会吗？"
        title="确认退出展会"
        onClose={() => setWithdrawActivityModal(false)}
        onConfirm={handleConfirmWithdrawActivity}
      />
    </>
  );
};

export default Exhibitor;
