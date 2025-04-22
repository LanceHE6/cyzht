import {
  Avatar,
  Button,
  Card,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
  Spacer,
  Tab,
  Tabs,
  Tooltip,
} from "@heroui/react";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

import {
  SelectionMenuIcon,
  ExploreIcon,
  AddIcon,
  MessageIcon,
  DefaultActivityIcon,
} from "@/components/icons";
import DefaultLayout from "@/layouts/default.tsx";
import UserProfilePopover from "@/components/user-profile-popover.tsx";
import { WebSocketClient, LocalStorage, Toast } from "@/utils/utils.ts";
import {
  axiosInstanceWithAuth,
  setNavigateCallback,
} from "@/utils/axios-instance.ts";
import { AddActivity } from "@/components/add-activity.tsx";
import { ExploreActivity } from "@/pages/explore-activity.tsx";
import { Activity } from "@/pages/activity.tsx";
import { getErrorMessage } from "@/utils/error-helper.ts";

export default function HomePage() {
  const navigate = useNavigate();
  const [selectedTab, setSelectedTab] = useState("explore");
  const [isAddActivityModalOpen, setIsAddActivityModalOpen] = useState(false);
  const [user, setUser] = useState<any>(LocalStorage.getUser());

  // 获取已加入展会列表
  const fetchJoinedActivities = async () => {
    const data = await getJoinedActivities();

    const joinedTabs = data.map((activityUser: any) => ({
      key: activityUser.activity.id,
      title: activityUser.activity.name,
      icon: activityUser.activity.icon ? (
        <Avatar
          isBordered
          className="cursor-pointer"
          size="sm"
          src={activityUser.activity.icon}
        />
      ) : (
        <DefaultActivityIcon />
      ),
      component: <Activity aid={activityUser.activity.id} />,
    }));

    setBaseMenuTabs((prevBaseTabs) => {
      // 确保不重复添加,防止useEffect多次触发
      const existingKeys = new Set(prevBaseTabs.map((tab) => tab.key));
      const newTabs = joinedTabs.filter((tab) => !existingKeys.has(tab.key));

      return [...prevBaseTabs, ...newTabs];
    });
  };

  useEffect(() => {
    setNavigateCallback(navigate);
    if (user === null) {
      navigate("/login");
    }
    setUser(user);
    fetchJoinedActivities();
  }, []);
  // 额外tabs用于发现展会页面点击跳转
  const [extraTabs, setExtraTabs] = useState<any[]>([]);
  // 处理添加新tab
  const handleAddNewTab = (newTab: any) => {
    // 使用函数式更新来确保总是获取最新的 baseMenuTabs
    setBaseMenuTabs((prevBaseTabs) => {
      const tabExists = [...prevBaseTabs, ...extraTabs].some(
        (tab) => tab.key === newTab.key,
      );

      if (!tabExists) {
        setExtraTabs((prevExtraTabs) => [...prevExtraTabs, newTab]);
      }

      return prevBaseTabs; // 保持 baseMenuTabs 不变
    });

    // 切换到新tab
    setSelectedTab(newTab.key);
  };

  // 分离基础tabs和额外tabs
  const [baseMenuTabs, setBaseMenuTabs] = useState([
    {
      key: "message",
      title: "私信",
      icon: <MessageIcon />,
      component: <MessageList />,
    },
    {
      key: "explore",
      title: "发现展会",
      icon: <ExploreIcon />,
      component: <ExploreActivity onAddNewTab={handleAddNewTab} />,
    },
  ]);

  // 合并后的tabs
  const menuTabs = [...baseMenuTabs, ...extraTabs];
  // 获取已加入展会列表请求函数
  const getJoinedActivities = async (): Promise<any[]> => {
    try {
      const response = await axiosInstanceWithAuth.get(
        "/api/v1/activity/joined",
      );

      if (response.data.code !== 0) {
        Toast.danger("获取已加入展会列表失败", response.data.msg);

        return [];
      }

      if (response.data.data.rows === null) {
        return [];
      }

      return response.data.data.rows;
    } catch (error) {
      Toast.danger("获取已加入展会列表失败", getErrorMessage(error));

      return [];
    }
  };

  // 用于刷新数据的回调函数
  const handleRefreshData = async () => {
    console.log("刷新数据");
    // 添加一段延迟
    setTimeout(async () => {
      fetchJoinedActivities();
    }, 1000);
  };
  // 登出
  const logout = () => {
    LocalStorage.removeToken();
    LocalStorage.removeUser();
    WebSocketClient.close(); // 关闭ws连接
    navigate("/login");
  };

  // 渲染当前选中的 Tab 的组件
  const renderSelectedComponent = () => {
    const selectedTabItem = menuTabs.find((tab) => tab.key === selectedTab);

    if (selectedTabItem) {
      return selectedTabItem.component;
    }

    return null;
  };

  return (
    <DefaultLayout>
      <div className="flex items-center justify-center h-full w-full bg-white py-0">
        <Card
          isBlurred
          className="card w-full h-5/6 min-w-96 grid grid-cols-24"
        >
          <Card
            className="menus flex items-center col-span-1 py-2 bg-[#ECECEC]"
            radius="none"
          >
            {/*<text className="text-medium justify-self-center">重邮展会通</text>*/}
            <Spacer y={4} />

            <UserProfilePopover user={user} />
            <Spacer y={4} />
            <div className="w-full max-w-[260px] py-2 flex-grow">
              <Tabs
                isVertical
                className={"menuTabs px-0 w-full"}
                defaultSelectedKey="explore"
                selectedKey={selectedTab}
                variant="light"
                onSelectionChange={(key) => setSelectedTab(key as string)}
              >
                {menuTabs.map((item) => (
                  <Tab
                    key={item.key}
                    className={"h-full w-5/6"}
                    title={
                      <Tooltip
                        key={item.key}
                        showArrow
                        color="default"
                        content={item.title}
                        offset={15}
                        placement={"right"}
                      >
                        <div>{item.icon}</div>
                      </Tooltip>
                    }
                  />
                ))}
              </Tabs>
              <Spacer y={4} />
              <Button
                isIconOnly
                className={"px-2 border-0 w-full"}
                variant={"light"}
                onPress={() => {
                  setIsAddActivityModalOpen(true);
                }}
              >
                <Tooltip
                  showArrow
                  color="default"
                  content={"添加展会"}
                  offset={15}
                  placement={"right"}
                >
                  <div>
                    <AddIcon />
                  </div>
                </Tooltip>
              </Button>
            </div>

            <div className="w-full max-w-[260px] px-2 py-2 flex items-end justify-end">
              <Dropdown backdrop="blur">
                <DropdownTrigger>
                  <Button
                    isIconOnly
                    className="px-2 border-0 w-64"
                    size="lg"
                    variant="bordered"
                  >
                    <SelectionMenuIcon />
                  </Button>
                </DropdownTrigger>
                <DropdownMenu aria-label="Static Actions">
                  <DropdownItem key="setting">设置</DropdownItem>
                  <DropdownItem key="about">关于</DropdownItem>
                  <DropdownItem key="logout" color="danger" onPress={logout}>
                    退出登录
                  </DropdownItem>
                </DropdownMenu>
              </Dropdown>
            </div>
          </Card>
          <Card
            className="col-span-23 items-center min-w-40 border-b-blue-500 border-0 bg-[#FFFFFF]"
            radius="none"
          >
            {renderSelectedComponent()}
            <AddActivity
              isOpen={isAddActivityModalOpen}
              onClose={() => setIsAddActivityModalOpen(false)}
              onRefreshData={handleRefreshData}
            />
          </Card>
        </Card>
      </div>
    </DefaultLayout>
  );
}

// 示例子组件
const MessageList = () => (
  <div>
    <h2>消息列表</h2>
    {/* 消息列表内容 */}
  </div>
);
