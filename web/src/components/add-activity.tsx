import {
  Avatar,
  Button,
  DateInput,
  DateValue,
  Form,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  Spacer,
} from "@heroui/react";
import { useEffect, useState } from "react";
import { Textarea } from "@heroui/input";
import { useNavigate } from "react-router-dom";

import {
  axiosInstanceWithAuth,
  setNavigateCallback,
} from "@/utils/axios-instance.ts";
import { Toast } from "@/utils/utils.ts";
import {
  CreateActivityIcon,
  DefaultActivityIcon,
  JoinActivityIcon,
} from "@/components/icons.tsx";
import { getCurrentDateTime } from "@/utils/datetime.ts";

export const AddActivity = ({
  isOpen,
  onClose,
  onRefreshData,
}: {
  isOpen: boolean;
  onClose: () => void;
  onRefreshData: () => void;
}) => {
  // 为请求设置导航回调
  const navigate = useNavigate();

  useEffect(() => {
    setNavigateCallback(navigate);
  }, [navigate]);

  // 新增状态管理
  const [selectedOption, setSelectedOption] = useState<string | null>(null);

  // 关闭模态框函数
  const closeModal = () => {
    onClose();
    setSelectedOption(null);
  };

  // 创建展会表单数据
  const [avatarFile, setAvatarFile] = useState<File | null>(null);
  const [avatarPreview, setAvatarPreview] = useState<string | null>(null);
  const [avatar] = useState(
    <Avatar icon={<DefaultActivityIcon />} size="lg" />,
  );
  const [name, setName] = useState("");
  const [introduce, setIntroduce] = useState("");
  const [startAt, setStartAt] = useState<DateValue | null>(
    getCurrentDateTime(),
  );
  const [endAt, setEndAt] = useState<DateValue | null>(getCurrentDateTime());
  const [location, setLocation] = useState("");
  // 处理创建展会表单提交
  const handleCreateEvent = async () => {
    const formData = new FormData();

    if (avatarFile) {
      formData.append("avatar", avatarFile);
    }
    formData.append("name", name);
    formData.append("introduce", introduce);
    formData.append("start_at", startAt?.toString() ?? "");
    formData.append("end_at", endAt?.toString() ?? "");
    formData.append("location", location);

    try {
      const response = await axiosInstanceWithAuth.post(
        "/api/v1/activity/add",
        formData,
      );

      console.log("response: ", response);

      if (response.data.code === 0) {
        Toast.success("创建展会成功", null);
        onRefreshData();
        closeModal();
      } else {
        Toast.danger("创建展会失败", response.data.msg);
      }
    } catch (error) {
      console.error("Error creating event:", error);
      Toast.danger("创建展会失败", "网络错误或其他问题");
    }
  };
  // 创建展会表单
  const createEventForm = (
    <>
      <ModalHeader className="flex flex-col gap-1">创建展会</ModalHeader>
      <ModalBody>
        <Form>
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
            isRequired
            label="展会名称"
            labelPlacement="outside"
            size="lg"
            type="text"
            value={name}
            onValueChange={setName}
          />
          <Spacer y={4} />
          <Textarea
            label="展会描述"
            labelPlacement="outside"
            size="lg"
            value={introduce}
            onValueChange={setIntroduce}
          />
          <Spacer y={4} />
          <div className="w-full max-w-xl flex flex-row gap-4">
            <div className="w-full flex flex-col gap-1">
              <DateInput
                hideTimeZone
                defaultValue={getCurrentDateTime()}
                label="开始时间"
                minValue={getCurrentDateTime().subtract({ days: 1 })}
                value={startAt}
                onChange={setStartAt}
              />
            </div>
            <div className="w-full flex flex-col gap-2">
              <DateInput
                defaultValue={getCurrentDateTime().add({ days: 1 })}
                label="结束时间"
                // maxValue={today(getLocalTimeZone())}
                value={endAt}
                onChange={setEndAt}
              />
            </div>
          </div>
          <Input
            isRequired
            label="举办地点"
            labelPlacement="outside"
            size="lg"
            type="text"
            value={location}
            onValueChange={setLocation}
          />
        </Form>
      </ModalBody>
      <ModalFooter>
        <Button color="default" onPress={() => setSelectedOption(null)}>
          返回
        </Button>
        <Button color="primary" type="submit" onPress={handleCreateEvent}>
          创建
        </Button>
      </ModalFooter>
    </>
  );

  const [aid, setAid] = useState("");
  // 处理加入展会表单提交
  const handleJoinEvent = async () => {
    const response = await axiosInstanceWithAuth.post(
      `/api/v1/activity/${aid}/join`,
    );

    if (response.data.code === 0) {
      Toast.success("加入展会成功", null);
      onRefreshData();
      closeModal();
    } else {
      Toast.warning("加入展会失败", response.data.msg);
    }
  };
  // 加入展会表单
  const joinEventForm = (
    <>
      <ModalHeader className="flex flex-col gap-1">加入展会</ModalHeader>
      <ModalBody>
        <Input
          isRequired
          label="展会ID"
          labelPlacement="outside"
          size="lg"
          type="number"
          value={aid}
          onValueChange={setAid}
        />
      </ModalBody>
      <ModalFooter>
        <Button color="default" onPress={() => setSelectedOption(null)}>
          返回
        </Button>
        <Button color="primary" onPress={handleJoinEvent}>
          加入
        </Button>
      </ModalFooter>
    </>
  );

  return (
    <>
      {/* 模态框 */}
      <Modal backdrop="transparent" isOpen={isOpen} onClose={onClose}>
        <ModalContent>
          {!selectedOption && (
            <>
              <ModalHeader className="flex flex-col gap-1">
                添加展会
              </ModalHeader>
              <ModalBody className="py-10">
                <Button
                  color="default"
                  size="lg"
                  startContent={<CreateActivityIcon />}
                  variant="shadow"
                  onPress={() => setSelectedOption("create")}
                >
                  创建展会
                </Button>
                <Button
                  color="default"
                  size="lg"
                  startContent={<JoinActivityIcon />}
                  variant="shadow"
                  onPress={() => setSelectedOption("join")}
                >
                  加入展会
                </Button>
              </ModalBody>
            </>
          )}
          {selectedOption === "create" && createEventForm}
          {selectedOption === "join" && joinEventForm}
        </ModalContent>
      </Modal>
    </>
  );
};
