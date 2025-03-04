import React, { useEffect, useState } from "react";
import {
  Avatar,
  AvatarIcon,
  Badge,
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
  useDisclosure,
} from "@heroui/react";

import { axiosInstanceWithAuth } from "../utils/axios-instance.ts";

import { LocalStorage, Toast } from "@/utils/utils.ts";

export interface UserInfoProps {
  user: any;
}
const UserProfilePopover: React.FC<UserInfoProps> = (props: any) => {
  let { user } = props;

  const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();
  const [uid, setUID] = useState(user.id);
  const [nickname, setNickname] = useState(user.nickname);
  const [sex, setSex] = useState("");
  const [avatar, setAvatar] = useState(
    <Avatar icon={<AvatarIcon />} size="lg" />,
  );

  const [avatarFile, setAvatarFile] = useState<File | null>(null);
  const [avatarPreview, setAvatarPreview] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setNickname(user.nickname);
    setSex("" + user.sex); // 将性别值转换为字符串
    setUID(user.id);
    if (user.avatar !== "") {
      setAvatar(
        <Avatar
          isBordered
          as="button"
          className="cursor-pointer"
          size="lg"
          src={user.avatar}
        />,
      );
    }
  }, [user]);

  // 修改资料接口
  const updateProfile = async () => {
    setIsLoading(true);
    try {
      // 如果有新的头像文件，先上传头像
      if (avatarFile) {
        const err = await uploadAvatar(avatarFile);

        if (err) {
          return; // 如果头像上传失败，停止后续操作
        } else {
          Toast.success("头像更新成功", null);
        }
      }

      // 构造资料更新的数据
      const data = {
        nickname: nickname || user.nickname,
        sex: sex === "1" ? 1 : sex === "2" ? 2 : -1,
      };
      console.log("data: ", data);

      // 调用资料更新接口
      const response = await axiosInstanceWithAuth.put(
        "/api/v1/user/update/profile",
        data,
      );

      if (response.status === 200) {
        Toast.success("资料更新成功", null);
        onOpenChange(); // 关闭模态框
      } else {
        throw new Error("资料更新失败");
      }
    } catch (error) {
      console.error("资料更新出错:", error);
      Toast.danger("资料更新失败", null);
    } finally {
      setIsLoading(false);
      setTimeout(async () => {
        user = await LocalStorage.updateUser();
        console.log("updated user: ", user);
        // 延迟刷新页面
        onClose();
        // window.location.reload();
      }, 2000);
    }
  };

  // 上传头像接口
  const uploadAvatar = async (file: File) => {
    const formData = new FormData();

    formData.append("file", file);

    try {
      const response = await axiosInstanceWithAuth.put(
        "/api/v1/user/update/avatar",
        formData,
      );

      console.log("response: ", response);

      if (response.status === 200) {
        Toast.success("头像上传成功", null);
      } else {
        throw new Error("头像上传失败");
      }
    } catch (error) {
      console.error("头像上传出错:", error);
      Toast.danger("头像上传失败: ", "" + error);

      return error;
    }
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
                <Badge color="success" content="" placement="bottom-right">
                  {avatar}
                </Badge>
                <div className="flex flex-col items-start justify-center">
                  <h4 className="text-small font-semibold leading-none text-default-600">
                    {nickname}
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
                    <TableCell>{uid}</TableCell>
                  </TableRow>
                  <TableRow key="2">
                    <TableCell className="text-gray-400">性别</TableCell>
                    <TableCell>
                      {sex === "1" ? "男" : sex === "2" ? "女" : "保密"}
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
                <div className="flex justify-center mb-4">
                  {avatarPreview ? (
                    <Avatar
                      isBordered
                      as="button"
                      className="cursor-pointer"
                      size="lg"
                      src={avatarPreview}
                    />
                  ) : (
                    avatar
                  )}
                </div>
                <Input
                  accept="image/*"
                  type="file"
                  onChange={(e) => {
                    const file = e.target.files?.[0];

                    if (file) {
                      setAvatarFile(file);
                      setAvatarPreview(URL.createObjectURL(file));
                    }
                  }}
                />

                <Input
                  label="昵称"
                  labelPlacement="outside"
                  value={nickname}
                  variant="bordered"
                  onValueChange={setNickname}
                />
                <RadioGroup
                  color="primary"
                  defaultValue={sex}
                  label="性别"
                  orientation="horizontal"
                  value={sex}
                  onValueChange={setSex}
                >
                  <Radio value="1">男</Radio>
                  <Radio value="2">女</Radio>
                  <Radio value="-1">保密</Radio>
                </RadioGroup>
                {/*<Textarea*/}
                {/*  classNames={{*/}
                {/*    input: "resize-y min-h-[40px]",*/}
                {/*  }}*/}
                {/*  label="个性签名"*/}
                {/*  labelPlacement="outside"*/}
                {/*  value={signature}*/}
                {/*  variant="bordered"*/}
                {/*  onValueChange={setSignature}*/}
                {/*/>*/}
              </ModalBody>
              <ModalFooter>
                <Button color="danger" variant="flat" onPress={onClose}>
                  取消
                </Button>
                <Button
                  color="primary"
                  isLoading={isLoading}
                  onPress={updateProfile}
                >
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
