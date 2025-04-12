import {
  Avatar,
  Button,
  Card,
  Divider,
  ScrollShadow,
  Image,
  Input,
  Spinner,
} from "@heroui/react";
import React, { useState, useEffect, useRef, useCallback } from "react";
import { Textarea } from "@heroui/input";
import { useNavigate } from "react-router-dom";

import {
  axiosInstanceWithAuth,
  setNavigateCallback,
} from "@/utils/axios-instance.ts";
import { LocalStorage, Toast, WebSocketClient } from "@/utils/utils.ts";
import { formatDate } from "@/utils/datetime.ts";
import { AddIcon, SearchIcon } from "@/components/icons.tsx";
import { getErrorMessage } from "@/utils/error-helper.ts";

interface user {
  id: string;
  nickname: string;
  avatar: string;
}

interface activity {
  id: string;
  name: string;
  introduce: string;
  start_at: string;
  end_at: string;
  location: string;
}

interface attachment {
  file_id: string;
  file_url: string;
  thumb_url?: string;
  file_size: number;
  mime_type: string;
  width?: number;
  height?: number;
}

interface Message {
  id: string;
  created_at: string;
  updated_at: string;
  activity: activity;
  exhibitor: {};
  from_user: user;
  to_user: user | null;
  msg_type: number; // 1-文本 2-图片 3-复合
  content: string;
  status: number;
  attachments: attachment[];
}

export interface ChatProps {
  aid: string;
}

const Chat: React.FC<ChatProps> = (props: ChatProps) => {
  const navigate = useNavigate();
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputValue, setInputValue] = useState<string>("");
  const [currentUser] = useState<any>(LocalStorage.getUser());
  const [isUploading, setIsUploading] = useState(false);
  const [localImages, setLocalImages] = useState<
    { file: File; preview: string }[]
  >([]);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const messageContainerRef = useRef<HTMLDivElement>(null);
  const { aid } = props;
  const aidRef = useRef(aid);

  aidRef.current = aid;

  // 初始化
  useEffect(() => {
    setNavigateCallback(navigate);
    if (!currentUser) {
      Toast.danger("请先登录", "请先登录");
      navigate("/login");
    }
    fetchMsg();
  }, [aid, navigate, currentUser]);

  // 获取历史消息
  const fetchMsg = async () => {
    try {
      const response = await axiosInstanceWithAuth.get(
        `/api/v1/activity/${aid}/msg`,
        {
          params: { page: -1, page_size: 10 },
        },
      );

      const parsedMessages: Message[] = response.data.data.rows.map(
        (msg: any) => ({
          ...msg,
          from_user: {
            id: msg.from_user.id,
            nickname: msg.from_user.nickname,
            avatar: msg.from_user.avatar,
          },
          to_user: msg.to_user
            ? {
                id: msg.to_user.id,
                nickname: msg.to_user.nickname,
                avatar: msg.to_user.avatar,
              }
            : null,
        }),
      );

      setMessages(parsedMessages);
    } catch (error) {
      Toast.danger("获取消息失败", getErrorMessage(error));
    }
  };

  // 滚动到底部
  const scrollToBottom = useCallback(() => {
    if (messageContainerRef.current?.lastElementChild) {
      messageContainerRef.current.lastElementChild.scrollIntoView({
        behavior: "smooth",
      });
    }
  }, []);

  // 上传图片到文件服务
  const uploadImage = async (file: File): Promise<attachment> => {
    setIsUploading(true);
    try {
      const formData = new FormData();

      formData.append("file", file);
      formData.append("file_name", file.name);
      formData.append("owner_type", "message");
      formData.append("need_thumb", "true");

      const response = await axiosInstanceWithAuth.post(
        "/api/v1/upload/image",
        formData,
        {
          headers: { "Content-Type": "multipart/form-data" },
        },
      );

      return {
        file_id: response.data.data.file_id,
        file_url: response.data.data.file_url,
        thumb_url: response.data.data.thumb_url,
        file_size: file.size,
        mime_type: file.type,
        width: response.data.data.width,
        height: response.data.data.height,
      };
    } catch (error) {
      Toast.danger("图片上传失败", getErrorMessage(error));
      throw error;
    } finally {
      setIsUploading(false);
    }
  };

  // 处理图片选择
  const handleImageSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;

    if (!files || files.length === 0) return;

    // 转换为数组并过滤非图片文件
    const imageFiles = Array.from(files).filter((file) =>
      file.type.startsWith("image/"),
    );

    // 生成预览并保存到本地状态
    const newImages = imageFiles.map((file) => ({
      file,
      preview: URL.createObjectURL(file),
    }));

    setLocalImages((prev) => [...prev, ...newImages]);

    // 清空input，允许重复选择同一文件
    if (fileInputRef.current) fileInputRef.current.value = "";
  };

  // 移除本地图片
  const removeLocalImage = (index: number) => {
    setLocalImages((prev) => {
      const newImages = [...prev];

      URL.revokeObjectURL(newImages[index].preview); // 释放内存
      newImages.splice(index, 1);

      return newImages;
    });
  };

  // 发送消息（处理所有类型）
  const sendMessage = async () => {
    // 验证是否有内容可发送
    if (inputValue.trim() === "" && localImages.length === 0) {
      Toast.warning("发送失败", "消息内容不能为空");

      return;
    }

    try {
      setIsUploading(true);

      // 1. 先上传所有图片
      let attachments: attachment[] = [];

      if (localImages.length > 0) {
        const uploadPromises = localImages.map((img) => uploadImage(img.file));

        attachments = await Promise.all(uploadPromises);
      }

      // 2. 确定消息类型
      let msgType: number;

      if (inputValue.trim() === "" && attachments.length > 0) {
        msgType = 2; // 纯图片消息
      } else if (inputValue.trim() !== "" && attachments.length === 0) {
        msgType = 1; // 纯文本消息
      } else {
        msgType = 3; // 复合消息
      }

      // 3. 发送消息
      await axiosInstanceWithAuth.post(`/api/v1/activity/${aid}/send`, {
        msg_type: msgType,
        content: inputValue,
        attachments: attachments,
      });

      // 4. 清空输入状态
      setInputValue("");
      setLocalImages([]);
    } catch (error) {
      Toast.warning("发送消息失败", getErrorMessage(error));
    } finally {
      setIsUploading(false);
    }
  };

  // 快捷键发送
  const handleHotKeysPress = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.ctrlKey && event.key === "Enter") {
      sendMessage();
      scrollToBottom();
    }
  };

  // WebSocket消息处理
  const handleWebSocketMessage = useCallback(
    (message: Message) => {
      if (
        message.activity.id === aidRef.current &&
        message.exhibitor === null
      ) {
        setMessages((prev) => [...prev, message]);
        scrollToBottom();
      }
    },
    [scrollToBottom],
  );

  useEffect(() => {
    WebSocketClient.onMessage(handleWebSocketMessage);

    return () => {
      WebSocketClient.offMessage(handleWebSocketMessage);
      // 清理预览URL内存
      localImages.forEach((img) => URL.revokeObjectURL(img.preview));
    };
  }, [handleWebSocketMessage, localImages]);

  // 渲染消息内容
  const renderMessageContent = (message: Message) => {
    switch (message.msg_type) {
      case 1: // 文本
        return <p className="whitespace-pre-wrap">{message.content}</p>;

      case 2: // 图片
        return (
          <div className="space-y-2">
            {message.attachments.map((att, idx) => (
              <Image
                key={idx}
                alt="图片消息"
                className="max-w-full rounded-lg"
                src={att.thumb_url ? att.thumb_url : att.file_url}
                style={{
                  maxWidth: "100%",
                  maxHeight: "300px",
                  cursor: "pointer",
                }}
                onClick={() => window.open(att.file_url, "_blank")}
              />
            ))}
          </div>
        );

      case 3: // 复合消息
        return (
          <div className="space-y-2">
            {message.attachments.map((att, idx) => (
              <Image
                key={idx}
                alt="图片消息"
                className="max-w-full rounded-lg"
                src={att.thumb_url || att.file_url}
                style={{
                  maxWidth: "100%",
                  maxHeight: "300px",
                  cursor: "pointer",
                }}
                onClick={() => window.open(att.file_url, "_blank")}
              />
            ))}
            {message.content && (
              <p className="whitespace-pre-wrap">{message.content}</p>
            )}
          </div>
        );

      default:
        return <p className="text-gray-500">[不支持的消息类型]</p>;
    }
  };

  return (
    <Card
      className="col-span-9 h-full bg-[#FFF6FF] flex flex-col"
      radius="none"
    >
      {/* 聊天室标题 */}
      <div className="sticky top-0 z-10 flex justify-between items-center p-4 px-10 bg-white border-b border-gray-200">
        <div className="text-xl font-bold">展会大厅</div>
      </div>
      <Divider className="" />

      {/* 聊天消息区域 */}
      <Card className="flex flex-col flex-grow overflow-y-auto" radius="none">
        <ScrollShadow
          ref={messageContainerRef}
          hideScrollBar
          className="space-y-6 p-4 px-10"
        >
          {messages.map((message) => (
            <div
              key={message.id}
              className={`flex ${message.from_user.id === currentUser.id ? "justify-end" : "justify-start"}`}
            >
              {message.from_user.id !== currentUser.id && (
                <Avatar
                  showFallback
                  alt={`${message.from_user.nickname}'s avatar`}
                  className="w-8 h-8 rounded-full mr-2"
                  name={message.from_user.nickname}
                  src={message.from_user.avatar}
                />
              )}

              <Card
                className={`border-none w-fit ${message.from_user.id === currentUser.id ? "bg-blue-100" : ""}`}
                style={{ maxWidth: "50%" }}
              >
                <div className="message p-2 w-auto">
                  <div
                    className={`flex ${message.from_user.id === currentUser.id ? "justify-end" : "justify-start"} px-1 items-center`}
                  >
                    {message.from_user.id !== currentUser.id ? (
                      <>
                        <div className="text-sm font-bold mr-2">
                          {message.from_user.nickname}
                        </div>
                        <div className="text-gray-500 text-xs">
                          {formatDate(message.created_at)}
                        </div>
                      </>
                    ) : (
                      <>
                        <div className="text-gray-500 text-xs mr-2">
                          {formatDate(message.created_at)}
                        </div>
                        <div className="text-sm font-bold">
                          {message.from_user.nickname}
                        </div>
                      </>
                    )}
                  </div>
                  <div className="w-fit max-w-full px-2 py-2">
                    {renderMessageContent(message)}
                  </div>
                </div>
              </Card>

              {message.from_user.id === currentUser.id && (
                <Avatar
                  showFallback
                  alt={`${message.from_user.nickname}'s avatar`}
                  className="w-8 h-8 rounded-full ml-2"
                  name={message.from_user.nickname}
                  src={message.from_user.avatar}
                />
              )}
            </div>
          ))}
        </ScrollShadow>
      </Card>

      {/* 输入区域 */}
      <div className="sticky bottom-0 z-10 w-full bg-white border-t border-gray-200 flex flex-col">
        {/* 图片预览区域 */}
        {localImages.length > 0 && (
          <div className="flex p-2 space-x-2 overflow-x-auto">
            {localImages.map((img, index) => (
              <div key={index} className="relative">
                <Image
                  alt="图片预览"
                  className="w-20 h-20 object-cover rounded-lg"
                  src={img.preview}
                />
                <Button
                  isIconOnly
                  className="absolute top-0 right-0 transform translate-x-1/2 -translate-y-1/2"
                  size="sm"
                  variant="flat"
                  onPress={() => removeLocalImage(index)}
                >
                  <SearchIcon size={14} />
                </Button>
              </div>
            ))}
          </div>
        )}

        {/* 文本输入区域 */}
        <div className="flex">
          <Textarea
            className="flex-grow"
            endContent={
              <div className="flex p-2 space-x-2">
                <Button
                  isIconOnly
                  isLoading={isUploading}
                  size="md"
                  onPress={() => fileInputRef.current?.click()}
                >
                  {isUploading ? <Spinner size="sm" /> : <AddIcon />}
                </Button>
                <Input
                  ref={fileInputRef}
                  multiple
                  accept="image/*"
                  className="hidden"
                  type="file"
                  onChange={handleImageSelect}
                />
              </div>
            }
            maxRows={4}
            minRows={1}
            placeholder="输入消息..."
            size="lg"
            type="text"
            value={inputValue}
            onKeyDown={handleHotKeysPress}
            onValueChange={setInputValue}
          />
          <Button
            className="ml-2"
            color="primary"
            isDisabled={inputValue.trim() === "" && localImages.length === 0}
            isLoading={isUploading}
            size="lg"
            onPress={sendMessage}
          >
            发送
          </Button>
        </div>
      </div>
    </Card>
  );
};

export default Chat;
