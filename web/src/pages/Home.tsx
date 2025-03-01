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
} from "@nextui-org/react";
import { useNavigate } from "react-router-dom";

import {
  MessageIcon,
  FriendsIcon,
  SelectionMenuIcon,
} from "@/components/icons";
import { removeToken, removeUser } from "@/utils/localStorage.ts";
import DefaultLayout from "@/layouts/default.tsx";

export default function Home() {
  const navigate = useNavigate();
  // 用户信息处理
  // 登出
  const logout = () => {
    removeUser();
    removeToken();
    navigate("/login");
  };

  return (
    <DefaultLayout>
      <div className="container flex items-center justify-center h-screen bg-amber-50 ">
        <Card isBlurred className="card w-3/4 h-5/6 min-w-96 grid grid-cols-12">
          <Card
            className="menus flex items-center col-span-1 py-2 bg-[#ECECEC]"
            radius="none"
          >
            <text className="text-medium justify-self-center">重邮展会通</text>
            <Spacer y={4} />
            {/*<UserProfilePopover avatar={avatar} user={formattedUser} />*/}
            <Spacer y={4} />
            <div className="w-full max-w-[260px] px-2 py-2 flex-grow">
              <Tabs isVertical defaultSelectedKey="message" variant="light">
                <Tab key="message" className="h-full">
                  <MessageIcon />
                </Tab>
                <Tab key="friends" className="h-full">
                  <FriendsIcon />
                </Tab>
              </Tabs>
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
                  <DropdownItem key="logout" color="danger" onClick={logout}>
                    退出登录
                  </DropdownItem>
                </DropdownMenu>
              </Dropdown>
            </div>
          </Card>
          <div className="message-list col-span-3 min-w-40 border-b-blue-500 border-0 bg-[#FFFFFF]">
            消息列表
          </div>
          <div className="menus col-span-8 min-w-52 border-r-emerald-500 border-0 bg-[#F2F2F2]">
            聊天框
          </div>
        </Card>
      </div>
    </DefaultLayout>
  );
}
