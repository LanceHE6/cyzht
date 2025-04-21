import {
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  Input,
  Textarea,
  Button,
  Snippet,
  User,
} from "@heroui/react";
import { useState } from "react";

import { formatDate } from "@/utils/datetime.ts";

// 添加参展商模态框
export const AddExhibitorModal: React.FC<{
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (name: string, introduce: string) => void;
}> = ({ isOpen, onClose, onSubmit }) => {
  const [exhibitorName, setExhibitorName] = useState("");
  const [exhibitorIntroduce, setExhibitorIntroduce] = useState("");

  const handleSubmit = () => {
    onSubmit(exhibitorName, exhibitorIntroduce);
    setExhibitorName("");
    setExhibitorIntroduce("");
    onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalContent>
        <ModalHeader className="flex flex-col gap-1">添加参展商</ModalHeader>
        <ModalBody>
          <Input
            aria-label={"exhibitorName"}
            label="参展商名称"
            placeholder="请输入参展商名称"
            value={exhibitorName}
            onChange={(e) => setExhibitorName(e.target.value)}
          />
          <Textarea
            aria-label={"exhibitorIntroduce"}
            label="参展商介绍"
            placeholder="请输入参展商介绍"
            value={exhibitorIntroduce}
            onChange={(e) => setExhibitorIntroduce(e.target.value)}
          />
        </ModalBody>
        <ModalFooter>
          <Button color="danger" variant="light" onPress={onClose}>
            取消
          </Button>
          <Button color="primary" onPress={handleSubmit}>
            提交
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

// 新增确认模态框组件
export const ConfirmModal: React.FC<{
  isOpen: boolean;
  onClose: () => void;
  onConfirm: () => void;
  title: string;
  message: string;
}> = ({ isOpen, onClose, onConfirm, title, message }) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalContent>
        <ModalHeader>{title}</ModalHeader>
        <ModalBody>{message}</ModalBody>
        <ModalFooter>
          <Button color="danger" onPress={onConfirm}>
            确认
          </Button>
          <Button color="default" onPress={onClose}>
            取消
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

// 展会详情模态框
export const ActivityInfoModal: React.FC<{
  isOpen: boolean;
  onClose: () => void;
  activity: {
    id?: string;
    name?: string;
    creator?: {
      id: string;
      account: string;
      nickname: string;
      avatar: string;
    };
    introduce?: string;
    start_at?: string;
    end_at?: string;
    location?: string;
  };
}> = ({ isOpen, onClose, activity }) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalContent>
        <ModalHeader className="flex flex-col gap-1">展会详情</ModalHeader>
        <ModalBody>
          <div className="space-y-4">
            <div className="grid grid-cols-2">
              <p className="text-large">ID:</p>
              <Snippet
                color="default"
                tooltipProps={{ content: "复制展会ID" }}
                variant="flat"
              >
                {activity?.id || "暂无数据"}
              </Snippet>
            </div>
            <div className="grid grid-cols-2">
              <p className="text-large">展会名称:</p>
              <Snippet
                color="default"
                tooltipProps={{ content: "复制展会名称" }}
                variant="flat"
              >
                {activity?.name || "暂无数据"}
              </Snippet>
            </div>
            <div className="grid grid-cols-2">
              <p className="text-large">展会时间:</p>
              <p>
                {formatDate(activity?.start_at) || "暂无数据"} -{" "}
                {formatDate(activity?.end_at) || "暂无数据"}
              </p>
            </div>
            <div className="grid grid-cols-2">
              <p className="text-large">举办地点:</p>
              <p>{activity?.location || "暂无数据"}</p>
            </div>
            <div className="grid grid-cols-2">
              <p className="text-large">创建者:</p>
              <User
                avatarProps={{
                  src: activity?.creator?.avatar || "",
                }}
                description={activity?.creator?.account}
                name={activity?.creator?.nickname || "暂无数据"}
              />
            </div>
            <div>
              <p className="font-medium">介绍:</p>
              <Textarea
                isReadOnly
                defaultValue={activity?.introduce || "暂无介绍"}
                minRows={1}
              />
            </div>
          </div>
        </ModalBody>
        <ModalFooter>
          <Button color="primary" onPress={onClose}>
            关闭
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
