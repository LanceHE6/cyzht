import {
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
import { axiosInstanceWithAuth } from "@/utils/axios-instance.ts";
import { AddActivity } from "@/components/add-activity.tsx";

export default function HomePage() {
  const navigate = useNavigate();
  const [user, setUser] = useState<any>(LocalStorage.getUser());

  const defaultMenuTabs = [
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
      component: <ExploreActivities />,
    },
  ];
  // 左侧菜单
  const [menuTabs, setMenuTabs] = useState(defaultMenuTabs);
  const [selectedTab, setSelectedTab] = useState("explore");
  const [isAddActivityModalOpen, setIsAddActivityModalOpen] = useState(false);

  useEffect(() => {
    if (user === null) {
      navigate("/login");
    }
    setUser(user);
    getJoinedActivities();
  }, []);

  // 登出
  const logout = () => {
    LocalStorage.removeToken();
    LocalStorage.removeUser();
    WebSocketClient.close(); // 关闭ws连接
    navigate("/login");
  };

  const getJoinedActivities = async () => {
    try {
      const response = await axiosInstanceWithAuth.get(
        "/api/v1/activity/joined",
      );

      if (response.data.code !== 0) {
        Toast.danger("获取已加入展会列表失败", response.data.msg);

        return;
      }
      if (response.data.data.rows === null) {
        return;
      }
      // 将列表加入 menuTabs
      const newTabs = response.data.data.rows.map((activityUser: any) => ({
        key: activityUser.id,
        title: activityUser.activity.name,
        icon: <DefaultActivityIcon />,
        component: <JoinedActivity activity={activityUser.activity} />,
      }));

      // 更新 menuTabs 状态
      setMenuTabs([...defaultMenuTabs, ...newTabs]);
    } catch (error) {
      console.error("Error fetching joined activities:", error);
      Toast.danger("获取已加入展会列表失败", "网络错误或其他问题");
    }
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
      <div className="container flex items-center justify-center h-screen bg-white py-0">
        <Card
          isBlurred
          className="card w-full h-5/6 min-w-96 grid grid-cols-16"
        >
          <Card
            className="menus flex items-center col-span-1 py-2 bg-[#ECECEC]"
            radius="none"
          >
            {/*<text className="text-medium justify-self-center">重邮展会通</text>*/}
            <Spacer y={4} />

            <UserProfilePopover user={user} />
            <Spacer y={4} />
            <div className="w-full max-w-[260px] px-1 py-2 flex-grow">
              <Tabs
                isVertical
                className={"menuTabs px-0"}
                defaultSelectedKey="explore"
                selectedKey={selectedTab}
                variant="light"
                onSelectionChange={(key) => setSelectedTab(key as string)}
              >
                {menuTabs.map((item) => (
                  <Tab
                    key={item.key}
                    className={"h-full"}
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
          <div className="message-list col-span-15 min-w-40 border-b-blue-500 border-0 bg-[#FFFFFF]">
            {renderSelectedComponent()}
            <AddActivity
              isOpen={isAddActivityModalOpen}
              onClose={() => setIsAddActivityModalOpen(false)}
            />
          </div>
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

const ExploreActivities = () => (
  <div>
    <h2>发现展会</h2>
    {/* 发现展会内容 */}
  </div>
);

const JoinedActivity = ({ activity }: { activity: any }) => (
  <div>
    <h2>{activity.name}</h2>
    {/* 已加入展会内容 */}
  </div>
);
