import {
  Avatar,
  Badge,
  Card,
  CardBody,
  CardHeader,
  Chip,
  Popover,
  PopoverContent,
  PopoverTrigger,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@heroui/react";
import React, { useEffect, useState } from "react";

import { axiosInstanceWithAuth } from "@/utils/axios-instance.ts";
import { Toast } from "@/utils/utils.ts";

export interface MemberProps {
  aid: string;
}
const Member: React.FC<MemberProps> = (props: MemberProps) => {
  let { aid } = props;
  const [users, setUsers] = useState([]);
  // 获取成员列表数据
  const fetchMembers = async (aid: string) => {
    const response = await axiosInstanceWithAuth.get(
      `/api/v1/activity/${aid}/joined`,
    );

    if (response.data.code === 0) {
      console.log("fetchMembers", response.data.data.rows);

      return response.data.data.rows.map((row: any) => ({
        id: row.user.id,
        user: row.user,
        nickname: row.user.nickname,
        avatar: row.user.avatar,
        role: row.activity.creator.id === row.user.id ? "创建者" : "成员",
      }));
    } else {
      Toast.danger("获取成员列表失败", response.data.msg);

      return [];
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      const data = await fetchMembers(aid);

      setUsers(data);
    };

    fetchData();
  }, [aid]);

  return (
    <Card
      className="col-span-2 items-center border-b-blue-500 border-0 bg-[#FFF8FF]"
      radius="none"
    >
      <Table
        isHeaderSticky
        removeWrapper
        aria-label="用户列表"
        className="flex flex-col h-full overflow-y-auto"
      >
        <TableHeader>
          <TableColumn>成员列表</TableColumn>
        </TableHeader>
        <TableBody className="flex-1">
          {users.map((item) => (
            <TableRow key={item.id}>
              <TableCell>
                <div className="flex gap-2 items-center">
                  <UserInfoPopover user={item.user} />
                  <div className="flex flex-col">
                    <span className="text-small">{item.nickname}</span>
                  </div>
                  <div className="justify-end">
                    <Chip
                      color={item.role === "创建者" ? "danger" : "primary"}
                      size="sm"
                      variant="flat"
                    >
                      {item.role}
                    </Chip>
                  </div>
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Card>
  );
};

interface UserInfoProps {
  user: any;
}
const UserInfoPopover: React.FC<UserInfoProps> = (props: UserInfoProps) => {
  let { user } = props;

  return (
    <Popover placement="left">
      <PopoverTrigger>
        <div className="flex items-center justify-center">
          <Badge
            color={user.online_status === 1 ? "success" : "default"}
            content=""
            placement="bottom-right"
          >
            <Avatar
              alt={user.nickname}
              className="flex-shrink-0"
              size="sm"
              src={user.avatar}
            />
          </Badge>
        </div>
      </PopoverTrigger>
      <PopoverContent>
        <Card
          className="max-w-[300px] border-none bg-transparent"
          shadow="none"
        >
          <CardHeader className="justify-between">
            <div className="flex gap-3">
              <Badge
                color={user.online_status === 1 ? "success" : "default"}
                content=""
                placement="bottom-right"
              >
                <Avatar
                  alt={user.nickname}
                  className="flex-shrink-0"
                  size="sm"
                  src={user.avatar}
                />
              </Badge>
              <div className="flex flex-col items-start justify-center">
                <h4 className="text-small font-semibold leading-none text-default-600">
                  {user.nickname}
                </h4>
                <h5 className="text-small tracking-tight text-default-500">
                  {user.account}
                </h5>
              </div>
            </div>
          </CardHeader>
          <CardBody className="px-0 py-0">
            <Table
              hideHeader
              removeWrapper
              aria-label="Example static collection table"
            >
              <TableHeader>
                <TableColumn>KEY</TableColumn>
                <TableColumn>VALUE</TableColumn>
              </TableHeader>
              <TableBody>
                <TableRow key="1">
                  <TableCell className="text-gray-400">UID</TableCell>
                  <TableCell>{user.id}</TableCell>
                </TableRow>
                <TableRow key="2">
                  <TableCell className="text-gray-400">性别</TableCell>
                  <TableCell>
                    {user.sex === 1 ? "男" : user.sex === 2 ? "女" : "保密"}
                  </TableCell>
                </TableRow>
                {/*<TableRow key="2">*/}
                {/*  <TableCell className="text-gray-400">个性签名</TableCell>*/}
                {/*  <TableCell>{signature}</TableCell>*/}
                {/*</TableRow>*/}
                <TableRow key="3">
                  <TableCell className="text-gray-400">注册时间</TableCell>
                  <TableCell>
                    {new Date(user.created_at).toLocaleDateString()}
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </CardBody>
        </Card>
      </PopoverContent>
    </Popover>
  );
};

export default Member;
