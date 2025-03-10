import {
  Button,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
} from "@heroui/react";
import { useState } from "react";

import { axiosInstanceWithAuth } from "@/utils/axios-instance.ts";
import { Toast } from "@/utils/utils.ts";

export const AddActivity = ({ isOpen, onClose }: { isOpen: boolean; onClose: () => void }) => {
  // 新增状态管理
  const [selectedOption, setSelectedOption] = useState<string | null>(null);

  // 关闭模态框函数
  const closeModal = () => {
    onClose();
    setSelectedOption(null);
  };

  // 创建展会表单
  const createEventForm = (
    <form onSubmit={(e) => handleCreateEvent(e)}>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="name"
        >
          展会名称
        </label>
        <input
          required
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="name"
          placeholder="展会名称"
          type="text"
        />
      </div>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="description"
        >
          展会描述
        </label>
        <textarea
          required
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="description"
          placeholder="展会描述"
        />
      </div>
      <div className="flex items-center justify-between">
        <Button color="primary" type="submit">
          提交
        </Button>
        <Button color="danger" onPress={closeModal}>
          取消
        </Button>
      </div>
    </form>
  );

  // 加入展会表单
  const joinEventForm = (
    <form onSubmit={(e) => handleJoinEvent(e)}>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="inviteLink"
        >
          服务器ID
        </label>
        <input
          required
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="inviteLink"
          placeholder="输入邀请链接或服务器ID"
          type="text"
        />
      </div>
      <div className="flex items-center justify-between">
        <Button color="primary" type="submit">
          加入
        </Button>
        <Button color="danger" onPress={closeModal}>
          取消
        </Button>
      </div>
    </form>
  );

  // 处理创建展会表单提交
  const handleCreateEvent = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const name = formData.get("name") as string;
    const description = formData.get("description") as string;

    try {
      const response = await axiosInstanceWithAuth.post(
        "/api/v1/activity/create",
        {
          name,
          description,
        },
      );

      if (response.data.code === 0) {
        Toast.success("创建展会成功", null);
        closeModal();
      } else {
        Toast.danger("创建展会失败", response.data.msg);
      }
    } catch (error) {
      console.error("Error creating event:", error);
      Toast.danger("创建展会失败", "网络错误或其他问题");
    }
  };

  // 处理加入展会表单提交
  const handleJoinEvent = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const inviteLink = formData.get("inviteLink") as string;

    try {
      const response = await axiosInstanceWithAuth.post(
        "/api/v1/activity/join",
        {
          inviteLink,
        },
      );

      if (response.data.code === 0) {
        Toast.success("加入展会成功", null);
        closeModal();
      } else {
        Toast.danger("加入展会失败", response.data.msg);
      }
    } catch (error) {
      console.error("Error joining event:", error);
      Toast.danger("加入展会失败", "网络错误或其他问题");
    }
  };

  return (
    <>
      {/* 模态框 */}
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalContent>
          <ModalHeader className="flex flex-col gap-1">添加展会</ModalHeader>
          <ModalBody>
            {!selectedOption && (
              <>
                <Button onPress={() => setSelectedOption("create")}>
                  创建展会
                </Button>
                <Button onPress={() => setSelectedOption("join")}>
                  加入展会
                </Button>
              </>
            )}
            {selectedOption === "create" && createEventForm}
            {selectedOption === "join" && joinEventForm}
          </ModalBody>
          <ModalFooter>{/* 你可以在这里添加其他操作 */}</ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};
