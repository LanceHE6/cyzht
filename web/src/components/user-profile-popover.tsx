import React, { useEffect, useState } from "react";
import {
  Avatar,
  AvatarIcon,
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  Popover,
  PopoverContent,
  PopoverTrigger,
  Radio,
  RadioGroup,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
  Textarea,
  useDisclosure,
} from "@nextui-org/react";
import { toast } from "react-toastify";

import { axiosInstanceWithAuth } from "../utils/axios-instance.ts";

export interface UserInfoProps {
  user: any;
}
const UserProfilePopover: React.FC<UserInfoProps> = (props: any) => {
  const { user } = props;

  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const [nickname, setNickname] = useState("unknown");
  const [account, setAccount] = useState("unknown");
  const [sex, setSex] = useState("未知");
  const [signature, setSignature] = useState("无");
  const [avatar, setAvatar] = useState(
    <Avatar icon={<AvatarIcon />} size="lg" />,
  );

  useEffect(() => {
    if (user === null) {
      return;
    }
    setNickname(() => user.nickname);
    setAccount(() => user.account);
    setSex(() => user.sex);
    setSignature(() => user.signature);

    if (user.avatar !== "") {
      setAvatar(() => (
        <Avatar
          isBordered
          as="button"
          className="cursor-pointer"
          size="lg"
          src={user.avatar}
        />
      ));
    }
  }, [user]);

  const updateProfile = async () => {
    console.log(user);
    const data = {
      nickname: nickname === undefined ? user.nickname : nickname,
      sex: sex === undefined ? user.sex : sex,
      signature: signature === undefined ? user.signature : signature,
    };

    // 处理参数
    if (data.sex === "男") {
      data.sex = 1;
    } else if (data.sex === "女") {
      data.sex = 0;
    } else {
      data.sex = -1;
    }
    const response = await axiosInstanceWithAuth.put(
      "/api/v1/user/update/profile",
      data,
    );

    if (response.status !== 200) {
    }
    toast.success("更新成功");
    onOpenChange();
    console.log(response.data);
  };

  return (
    <div>
      <Popover placement="right">
        <PopoverTrigger>
          <div className="flex items-center justify-center">{avatar}</div>
        </PopoverTrigger>
        <PopoverContent>
          <Card
            className="max-w-[300px] border-none bg-transparent"
            shadow="none"
          >
            <CardHeader className="justify-between">
              <div className="flex gap-3">
                {avatar}
                <div className="flex flex-col items-start justify-center">
                  <h4 className="text-small font-semibold leading-none text-default-600">
                    {nickname}
                  </h4>
                  <h5 className="text-small tracking-tight text-default-500">
                    {account}
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
                    <TableCell className="text-gray-400">性别</TableCell>
                    <TableCell>{sex}</TableCell>
                  </TableRow>
                  <TableRow key="2">
                    <TableCell className="text-gray-400">个性签名</TableCell>
                    <TableCell>{signature}</TableCell>
                  </TableRow>
                  <TableRow key="3">
                    <TableCell className="text-gray-400">注册时间</TableCell>
                    <TableCell>{user.created_at}</TableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </CardBody>
            <CardFooter className="gap-3 flex items-center justify-center">
              <Button color="primary" radius="md" size="md" onPress={onOpen}>
                编辑资料
              </Button>
            </CardFooter>
          </Card>
        </PopoverContent>
      </Popover>

      <Modal isOpen={isOpen} placement="top-center" onOpenChange={onOpenChange}>
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader className="flex flex-col gap-1">
                编辑资料
              </ModalHeader>
              <ModalBody>
                <Input
                  label="昵称"
                  labelPlacement="outside"
                  value={nickname}
                  variant="bordered"
                  onValueChange={setNickname}
                />
                <RadioGroup
                  color="primary"
                  defaultValue={user.sex}
                  label="性别"
                  orientation="horizontal"
                  value={sex}
                  onValueChange={setSex}
                >
                  <Radio value="男">男</Radio>
                  <Radio value="女">女</Radio>
                  <Radio value="保密">保密</Radio>
                </RadioGroup>
                <Textarea
                  classNames={{
                    input: "resize-y min-h-[40px]",
                  }}
                  label="个性签名"
                  labelPlacement="outside"
                  value={signature}
                  variant="bordered"
                  onValueChange={setSignature}
                />
              </ModalBody>
              <ModalFooter>
                <Button color="danger" variant="flat" onPress={onClose}>
                  取消
                </Button>
                <Button color="primary" onPress={updateProfile}>
                  提交
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </div>
  );
};

export default UserProfilePopover;
