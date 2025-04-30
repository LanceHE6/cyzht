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
  Modal,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Input,
  Checkbox,
  useDisclosure,
  Badge,
} from "@heroui/react";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

import {
  SelectionMenuIcon,
  ExploreIcon,
  AddIcon,
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
import Activity from "@/pages/activity.tsx";
import { getErrorMessage } from "@/utils/error-helper.ts";
import { useActivityStore } from "@/store/use-activity-store.ts";

export default function HomePage() {
  const navigate = useNavigate();
  const [selectedTab, setSelectedTab] = useState("explore");
  const [isAddActivityModalOpen, setIsAddActivityModalOpen] = useState(false);
  const [user, setUser] = useState<any>(LocalStorage.getUser());
  const [isMenuVisible, setIsMenuVisible] = useState(true);
  const [isMobile, setIsMobile] = useState(false);

  const { unreadCounts } = useActivityStore();

  // 设置模态框
  const {
    isOpen: isSettingsOpen,
    onOpen: onSettingsOpen,
    onClose: onSettingsClose,
  } = useDisclosure();
  const [settingsTab, setSettingsTab] = useState("general");
  const {
    isOpen: isChangePasswordOpen,
    onOpen: onChangePasswordOpen,
    onClose: onChangePasswordClose,
  } = useDisclosure();
  const [passwordData, setPasswordData] = useState({
    oldPassword: "",
    newPassword: "",
    confirmPassword: "",
  });

  // 获取已加入展会列表
  const fetchJoinedActivities = async () => {
    const data = await getJoinedActivities();

    console.log("已加入展会列表", data);

    const joinedTabs = data.map((activityUser: any) => ({
      key: activityUser.activity.id,
      title: activityUser.activity.name,
      icon: activityUser.activity.icon ? (
        <Avatar isBordered size="sm" src={activityUser.activity.icon} />
      ) : (
        <DefaultActivityIcon />
      ),
      component: (
        <Activity
          aid={activityUser.activity.id}
          setIsMenuVisible={setIsMenuVisible}
        />
      ),
    }));

    setBaseMenuTabs((prevBaseTabs) => {
      const existingKeys = new Set(prevBaseTabs.map((tab) => tab.key));
      const newTabs = joinedTabs.filter((tab) => !existingKeys.has(tab.key));

      return [...prevBaseTabs, ...newTabs];
    });
  };

  useEffect(() => {
    if (window.innerWidth < 480) {
      setIsMobile(true);
    }
    setNavigateCallback(navigate);
    if (!user) {
      logout();
    }
    setUser(user);
    fetchJoinedActivities();
  }, []);

  const [extraTabs, setExtraTabs] = useState<any[]>([]);

  const handleAddNewTab = (newTab: any) => {
    setBaseMenuTabs((prevBaseTabs) => {
      const tabExists = [...prevBaseTabs, ...extraTabs].some(
        (tab) => tab.key === newTab.key,
      );

      if (!tabExists) {
        setExtraTabs((prevExtraTabs) => [...prevExtraTabs, newTab]);
      }

      return prevBaseTabs;
    });

    setSelectedTab(newTab.key);
  };

  const [baseMenuTabs, setBaseMenuTabs] = useState([
    {
      key: "explore",
      title: "发现展会",
      icon: <ExploreIcon />,
      component: (
        <ExploreActivity
          setIsMenuVisible={setIsMenuVisible}
          onAddNewTab={handleAddNewTab}
        />
      ),
    },
  ]);

  const menuTabs = [...baseMenuTabs, ...extraTabs];

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

  const handleRefreshData = async () => {
    console.log("刷新数据");
    setTimeout(async () => {
      fetchJoinedActivities();
    }, 1000);
  };

  const logout = () => {
    LocalStorage.removeToken();
    LocalStorage.removeUser();
    WebSocketClient.close();
    navigate("/login");
  };

  // 修改密码请求
  const handlePasswordChange = async () => {
    if (!passwordData.oldPassword) {
      Toast.danger("修改密码失败", "请输入旧密码");

      return;
    }
    if (!passwordData.newPassword) {
      Toast.danger("修改密码失败", "请输入新密码");

      return;
    }
    if (passwordData.newPassword !== passwordData.confirmPassword) {
      Toast.danger("修改密码失败", "新密码与确认密码不一致");

      return;
    }

    try {
      const response = await axiosInstanceWithAuth.put(
        "/api/v1/user/update/psw",
        {
          old_password: passwordData.oldPassword,
          new_password: passwordData.newPassword,
        },
      );

      if (response.data.code !== 0) {
        Toast.danger("修改密码失败", response.data.msg);

        return;
      }

      Toast.success("修改密码成功", "请重新登录");
      onChangePasswordClose();
      onSettingsClose();
      // 添加延时
      setTimeout(() => {
        logout();
      }, 2000);
    } catch (error) {
      Toast.danger("修改密码失败", getErrorMessage(error));
    }
  };

  const renderSelectedComponent = () => {
    const selectedTabItem = menuTabs.find((tab) => tab.key === selectedTab);

    return selectedTabItem ? selectedTabItem.component : null;
  };

  return (
    <DefaultLayout>
      <div className="flex items-center justify-center h-full w-full bg-[#e3fdfd] py-0">
        {isMobile ? (
          <Card
            isBlurred
            className="card w-full h-full xs:h-5/6 min-w-96 grid sm:grid-cols-24 grid-cols-8"
          >
            {isMenuVisible && (
              <Card
                className="menus flex items-center col-span-1 py-2 bg-[#cbf1f5]"
                radius="none"
              >
                <Spacer y={4} />
                {user && <UserProfilePopover user={user} />}
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
                    {menuTabs.map((item) => {
                      return (
                        <Tab
                          key={item.key}
                          className={"h-full w-full"}
                          title={
                            <Tooltip
                              showArrow
                              color="default"
                              content={item.title}
                              offset={15}
                              placement={"right"}
                            >
                              <Badge
                                color={"danger"}
                                content={unreadCounts[item.key]}
                                isInvisible={!(unreadCounts[item.key] > 0)}
                              >
                                <div className={"py-1"}>{item.icon}</div>
                              </Badge>
                            </Tooltip>
                          }
                        />
                      );
                    })}
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
                      <DropdownItem key="setting" onPress={onSettingsOpen}>
                        设置
                      </DropdownItem>
                      <DropdownItem key="about">关于</DropdownItem>
                      <DropdownItem
                        key="logout"
                        color="danger"
                        onPress={logout}
                      >
                        退出登录
                      </DropdownItem>
                    </DropdownMenu>
                  </Dropdown>
                </div>
              </Card>
            )}
            <Card
              className={`sm:col-span-23 ${isMenuVisible ? "col-span-7" : "col-span-8"} items-center min-w-40 border-b-blue-500 border-0 bg-[#FFFFFF]`}
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
        ) : (
          <Card
            isBlurred
            className="card w-full h-full xs:h-5/6 min-w-96 grid sm:grid-cols-24 grid-cols-8"
          >
            <Card
              className="menus flex items-center col-span-1 py-2 bg-[#cbf1f5]"
              radius="none"
            >
              <Spacer y={4} />
              {user && <UserProfilePopover user={user} />}
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
                          showArrow
                          color="default"
                          content={item.title}
                          offset={15}
                          placement={"right"}
                        >
                          <Badge
                            color={"danger"}
                            content={unreadCounts[item.key]}
                            isInvisible={!(unreadCounts[item.key] > 0)}
                          >
                            <div className={"py-1"}>{item.icon}</div>
                          </Badge>
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
                    <DropdownItem key="setting" onPress={onSettingsOpen}>
                      设置
                    </DropdownItem>
                    <DropdownItem
                      key="about"
                      href={"https://github.com/LanceHE6/cyzht"}
                    >
                      关于
                    </DropdownItem>
                    <DropdownItem key="logout" color="danger" onPress={logout}>
                      退出登录
                    </DropdownItem>
                  </DropdownMenu>
                </Dropdown>
              </div>
            </Card>

            <Card
              className="sm:col-span-23 col-span-7 items-center min-w-40 border-b-blue-500 border-0 bg-[#FFFFFF]"
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
        )}

        {/* 设置模态框 */}
        <Modal isOpen={isSettingsOpen} size="2xl" onClose={onSettingsClose}>
          <ModalContent>
            <ModalHeader className="flex flex-col gap-1">设置</ModalHeader>
            <ModalBody>
              <Tabs
                selectedKey={settingsTab}
                onSelectionChange={(key) => setSettingsTab(key as string)}
              >
                <Tab
                  key="general"
                  title={
                    <div className="flex items-center gap-2">
                      <span>通用</span>
                    </div>
                  }
                >
                  <div className="p-4">
                    <h3 className="text-lg font-medium mb-4">通用设置</h3>
                    <div className="space-y-4">
                      <Checkbox defaultSelected>启用黑暗模式</Checkbox>
                      <Checkbox defaultSelected>显示通知</Checkbox>
                    </div>
                  </div>
                </Tab>
                <Tab
                  key="security"
                  title={
                    <div className="flex items-center gap-2">
                      <span>账号安全</span>
                    </div>
                  }
                >
                  <div className="p-4">
                    <h3 className="text-lg font-medium mb-4">账号安全</h3>
                    <div className="space-y-4">
                      <div className="flex justify-between items-center">
                        <span>修改密码</span>
                        <Button size="sm" onPress={onChangePasswordOpen}>
                          修改
                        </Button>
                      </div>
                    </div>
                  </div>
                </Tab>
              </Tabs>
            </ModalBody>
            <ModalFooter>
              <Button color="primary" onPress={onSettingsClose}>
                关闭
              </Button>
            </ModalFooter>
          </ModalContent>
        </Modal>

        {/* 修改密码模态框 */}
        <Modal isOpen={isChangePasswordOpen} onClose={onChangePasswordClose}>
          <ModalContent>
            <ModalHeader className="flex flex-col gap-1">修改密码</ModalHeader>
            <ModalBody>
              <div className="space-y-4">
                <Input
                  isRequired
                  label="旧密码"
                  type="password"
                  value={passwordData.oldPassword}
                  onChange={(e) =>
                    setPasswordData({
                      ...passwordData,
                      oldPassword: e.target.value,
                    })
                  }
                />
                <Input
                  isRequired
                  errorMessage="密码长度不能小于6"
                  isInvalid={passwordData.newPassword.length < 6}
                  label="新密码"
                  minLength={6}
                  type="password"
                  value={passwordData.newPassword}
                  onChange={(e) =>
                    setPasswordData({
                      ...passwordData,
                      newPassword: e.target.value,
                    })
                  }
                />
                <Input
                  isRequired
                  errorMessage="密码不一致"
                  isInvalid={
                    passwordData.newPassword !== passwordData.confirmPassword
                  }
                  label="确认密码"
                  minLength={6}
                  type="password"
                  value={passwordData.confirmPassword}
                  onChange={(e) =>
                    setPasswordData({
                      ...passwordData,
                      confirmPassword: e.target.value,
                    })
                  }
                />
              </div>
            </ModalBody>
            <ModalFooter>
              <Button variant="light" onPress={onChangePasswordClose}>
                取消
              </Button>
              <Button
                color="primary"
                type="submit"
                onPress={handlePasswordChange}
              >
                确认修改
              </Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
      </div>
    </DefaultLayout>
  );
}
