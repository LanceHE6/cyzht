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
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
  Tooltip,
  useDisclosure,
} from "@heroui/react";
import { useState } from "react";

import { formatDate } from "@/utils/datetime.ts";
import { Notice } from "@/components/activity/exhibitor.tsx";
import { Toast } from "@/utils/utils.ts";
import { DeleteIcon, EditIcon, EyeIcon } from "@/components/icons.tsx";
import { axiosInstanceWithAuth } from "@/utils/axios-instance.ts";

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

interface NoticeModalProps {
  isOpen: boolean;
  onClose: () => void;
  notices: Notice[];
  userRole: number;
  onRefresh: () => void; // 刷新回调
  aid: string;
}
const columns = [
  { name: "标题", uid: "title" },
  { name: "发布人", uid: "creator" },
  { name: "发布时间", uid: "created_at" },
  { name: "操作", uid: "actions" },
];

export const NoticeModal: React.FC<NoticeModalProps> = ({
  isOpen,
  onClose,
  notices,
  userRole,
  onRefresh,
  aid,
}) => {
  // 内部状态管理
  const {
    isOpen: isFormOpen,
    onOpen: onFormOpen,
    onClose: onFormClose,
  } = useDisclosure();

  const {
    isOpen: isViewOpen,
    onOpen: onViewOpen,
    onClose: onViewClose,
  } = useDisclosure();

  const {
    isOpen: isConfirmOpen,
    onOpen: onConfirmOpen,
    onClose: onConfirmClose,
  } = useDisclosure();

  const [selectedNotice, setSelectedNotice] = useState<Notice | null>(null);
  const [actionType, setActionType] = useState<"edit" | "delete" | null>(null);

  // 查看公告详情
  const handleView = (notice: Notice) => {
    setSelectedNotice(notice);
    onViewOpen();
  };

  // 编辑公告
  const handleEdit = (notice: Notice) => {
    setSelectedNotice(notice);
    setActionType("edit");
    onFormOpen();
  };

  // 删除公告
  const handleDelete = (notice: Notice) => {
    setSelectedNotice(notice);
    setActionType("delete");
    onConfirmOpen();
  };

  // 提交表单 (新建或编辑)
  const handleSubmit = async (title: string, content: string) => {
    try {
      if (actionType === "edit" && selectedNotice) {
        await axiosInstanceWithAuth.post(
          `/api/v1/activity/${selectedNotice.id}/notice/update`,
          {
            title,
            content,
          },
        );
        Toast.success("公告更新成功", null);
      } else {
        await axiosInstanceWithAuth.post(
          `/api/v1/activity/${aid}/notice/publish`,
          {
            title,
            content,
          },
        );
        Toast.success("公告发布成功", null);
      }
      onRefresh(); // 刷新列表
      onFormClose();
    } catch (error) {
      Toast.danger("操作失败", "请稍后再试");
    }
  };

  // 确认删除
  const handleConfirmDelete = async () => {
    if (!selectedNotice) return;

    try {
      await axiosInstanceWithAuth.delete(
        `/api/v1/activity/notice/${selectedNotice.id}/del`,
      );
      Toast.success("公告删除成功", null);
      onRefresh(); // 刷新列表
      onConfirmClose();
    } catch (error) {
      Toast.danger("删除失败", "请稍后再试");
    }
  };

  const renderCell = (notice: Notice, columnKey: keyof Notice | "actions") => {
    switch (columnKey) {
      case "title":
        return (
          <span
            className="font-medium cursor-pointer"
            role="button"
            onClick={() => handleView(notice)}
          >
            {notice.title}
          </span>
        );
      case "created_at":
        return (
          <span className="text-sm">
            {new Date(notice.created_at).toLocaleString()}
          </span>
        );
      case "creator":
        return (
          <User
            avatarProps={{ src: notice.creator.avatar }}
            name={notice.creator.nickname}
          />
        );
      case "actions":
        return (
          <div className="flex items-center gap-2 justify-center">
            <Tooltip content="查看详情">
              <span
                className="text-lg text-default-400 cursor-pointer active:opacity-50"
                role="button"
                onClick={() => handleView(notice)}
              >
                <EyeIcon size={20} />
              </span>
            </Tooltip>

            {userRole >= 2 && (
              <>
                <Tooltip content="编辑公告">
                  <span
                    className="text-lg text-default-400 cursor-pointer active:opacity-50"
                    role="button"
                    onClick={() => handleEdit(notice)}
                  >
                    <EditIcon size={20} />
                  </span>
                </Tooltip>
                <Tooltip color="danger" content="删除公告">
                  <span
                    className="text-lg text-danger cursor-pointer active:opacity-50"
                    role="button"
                    onClick={() => handleDelete(notice)}
                  >
                    <DeleteIcon size={20} />
                  </span>
                </Tooltip>
              </>
            )}
          </div>
        );
      default:
        return notice[columnKey];
    }
  };

  return (
    <>
      <Modal isOpen={isOpen} size="3xl" onClose={onClose}>
        <ModalContent>
          <ModalHeader className="flex justify-between items-center">
            <span>公告</span>
            {userRole >= 2 && (
              <div className="px-3">
                <Button
                  color="primary"
                  size="md"
                  onPress={() => {
                    setSelectedNotice(null);
                    setActionType(null);
                    onFormOpen();
                  }}
                >
                  发布
                </Button>
              </div>
            )}
          </ModalHeader>
          <ModalBody>
            <Table removeWrapper aria-label="公告列表">
              <TableHeader columns={columns}>
                {(column) => (
                  <TableColumn
                    key={column.uid}
                    align={column.uid === "actions" ? "center" : "start"}
                  >
                    {column.name}
                  </TableColumn>
                )}
              </TableHeader>
              <TableBody emptyContent="暂无公告" items={notices}>
                {(item) => (
                  <TableRow key={item.id}>
                    {(columnKey) => (
                      <TableCell>
                        {renderCell(
                          item,
                          columnKey as keyof Notice | "actions",
                        )}
                      </TableCell>
                    )}
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </ModalBody>
          <ModalFooter>
            <div className="px-3">
              <Button onPress={onClose}>关闭</Button>
            </div>
          </ModalFooter>
        </ModalContent>
      </Modal>

      {/* 公告表单模态框 (新建/编辑) */}
      <NoticeFormModal
        isOpen={isFormOpen}
        notice={selectedNotice}
        onClose={onFormClose}
        onSubmit={handleSubmit}
      />

      <ViewNoticeModal
        isOpen={isViewOpen}
        notice={selectedNotice}
        onClose={onViewClose}
      />

      {/* 删除确认模态框 */}
      <ConfirmModal
        isOpen={isConfirmOpen}
        message="您确定要删除这条公告吗？此操作不可恢复。"
        title="确认删除公告"
        onClose={onConfirmClose}
        onConfirm={handleConfirmDelete}
      />
    </>
  );
};

interface NoticeFormModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (title: string, content: string) => void;
  notice?: Notice | null;
}

export const NoticeFormModal: React.FC<NoticeFormModalProps> = ({
  isOpen,
  onClose,
  onSubmit,
  notice,
}) => {
  const [title, setTitle] = useState(notice?.title || "");
  const [content, setContent] = useState(notice?.content || "");

  const handleSubmit = () => {
    if (!title || !content) {
      Toast.danger("请输入标题和内容", null);

      return;
    }
    onSubmit(title, content);
    onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalContent>
        <ModalHeader>{notice ? "编辑公告" : "发布公告"}</ModalHeader>
        <ModalBody className="space-y-4">
          <Input
            isRequired
            label="标题"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <Textarea
            isRequired
            label="内容"
            minRows={5}
            value={content}
            onChange={(e) => setContent(e.target.value)}
          />
        </ModalBody>
        <ModalFooter>
          <Button variant="light" onPress={onClose}>
            取消
          </Button>
          <Button color="primary" onPress={handleSubmit}>
            发布
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export const ViewNoticeModal: React.FC<{
  isOpen: boolean;
  onClose: () => void;
  notice: Notice | null;
}> = ({ isOpen, onClose, notice }) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalContent>
        <ModalHeader>公告详情</ModalHeader>
        <ModalBody>
          <>
            <Input
              disabled
              label="标题"
              labelPlacement={"outside"}
              value={notice?.title}
            />
            <Textarea
              readOnly
              label="内容"
              labelPlacement={"outside"}
              minRows={5}
              value={notice?.content}
            />
          </>
        </ModalBody>
        <ModalFooter>
          <Button variant="solid" onPress={onClose}>
            关闭
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
