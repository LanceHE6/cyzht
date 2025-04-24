import {
  Avatar,
  Button,
  Card,
  Divider,
  ScrollShadow,
  Image,
  Input,
  Spinner,
  Badge,
  Tooltip,
  Chip,
  Modal,
  ModalContent,
  ModalBody,
  ModalFooter,
  ModalHeader,
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
import {
  AttachmentIcon,
  CloseIcon,
  ImageIcon,
  FileIcon,
  ChevronDownIcon,
  CopyIcon,
  WithDrawIcon,
} from "@/components/icons.tsx";
import { getErrorMessage } from "@/utils/error-helper.ts";

// message数据定义接口
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

interface exhibitor {
  id: string;
  name: string;
  introduce: string;
}

interface attachment {
  file_id: string;
  file_url: string;
  thumb_url?: string;
  file_name: string;
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
  exhibitor: exhibitor;
  from_user: user;
  to_user: user | null;
  msg_type: number; // 1-文本 2-图片 3-复合 4-文件
  content: string;
  status: number;
  attachments: attachment[];
}

export interface ChatProps {
  aid: string;
  eid: string | null;
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
  const [newMessagesCount, setNewMessagesCount] = useState(0);
  const [isAtBottom, setIsAtBottom] = useState(true);
  const imageInputRef = useRef<HTMLInputElement>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const messageContainerRef = useRef<HTMLDivElement>(null);
  let { aid, eid } = props;
  const aidRef = useRef(aid);
  const eidRef = useRef(eid);
  const [exhibitorName, setExhibitorName] = useState<string | null>(null);
  const [isJoined, setIsJoined] = useState(false);

  eid = eid === null ? "0" : eid;
  aidRef.current = aid;
  eidRef.current = eid;

  // 初始化
  useEffect(() => {
    setNavigateCallback(navigate);
    if (!currentUser) {
      Toast.danger("请先登录", "请先登录");
      navigate("/login");
    }
    const fetchJoinedData = async () => {
      try {
        const response = await axiosInstanceWithAuth.get(
          `/api/v1/activity/joined`,
        );

        if (response.data.code === 0) {
          response.data.data.rows.map((row: any) => {
            if (row.activity.id === aid) {
              setIsJoined(true);
            }
          });
        }
      } catch (error) {
        console.error("获取已加入展会信息失败", error);
      }
    };

    fetchJoinedData();
    fetchMsg();
  }, [aid, eid, navigate, currentUser]);

  useEffect(() => {
    const fetchExhibitorInfo = async () => {
      try {
        if (eid && eid !== "0") {
          const response = await axiosInstanceWithAuth.get(
            `/api/v1/exhibitor/info/${eid}`,
          );

          if (response.data.code === 0) {
            const exhibitorInfo = response.data.data;

            setExhibitorName(exhibitorInfo.name);
          }
        } else {
          setExhibitorName(null);
        }
      } catch (error) {
        console.error("获取参展商信息失败", error);
      }
    };

    fetchExhibitorInfo();
  }, [eid]);

  // 消息加载相关状态
  const [isLoadingMore, setIsLoadingMore] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const [page, setPage] = useState(1);
  const pageSize = 10; // 每页消息数量

  // 获取历史消息,支持分页加载
  const fetchMsg = async (loadMore = false) => {
    try {
      if (loadMore) {
        setIsLoadingMore(true);
        // 保存当前滚动位置和第一条消息的DOM元素
        const container = messageContainerRef.current;

        if (!container) return;

        const prevScrollHeight = container.scrollHeight;
        const prevScrollTop = container.scrollTop;

        // 添加延迟模拟网络请求
        await new Promise((resolve) => setTimeout(resolve, 800));

        const response = await axiosInstanceWithAuth.get(
          `/api/v1/chat/${aid}/${eid}/msg`,
          {
            params: {
              page: loadMore ? page + 1 : 1,
              page_size: pageSize,
            },
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

        if (loadMore) {
          // 加载更多时，将新消息添加到前面
          setMessages((prev) => [...parsedMessages, ...prev]);
          setPage((prev) => prev + 1);
          // 检查是否还有更多消息
          setHasMore(parsedMessages.length >= pageSize);

          // 在下一个渲染周期调整滚动位置
          setTimeout(() => {
            if (!container) return;
            const newScrollHeight = container.scrollHeight;
            const scrollOffset = newScrollHeight - prevScrollHeight;

            container.scrollTop = prevScrollTop + scrollOffset;
          }, 0);
        } else {
          // 初始加载或刷新时，直接设置消息
          setMessages(parsedMessages);
          setPage(1);
          setHasMore(parsedMessages.length >= pageSize);
          // 初始加载后滚动到底部
          setTimeout(() => scrollToBottom("auto"), 100);
        }
      } else {
        // 非加载更多的情况（初始加载）
        setIsUploading(true);
        await new Promise((resolve) => setTimeout(resolve, 800));

        const response = await axiosInstanceWithAuth.get(
          `/api/v1/chat/${aid}/${eid}/msg`,
          {
            params: {
              page: 1,
              page_size: pageSize,
            },
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
        setPage(1);
        setHasMore(parsedMessages.length >= pageSize);
        setTimeout(() => scrollToBottom("auto"), 100);
      }
    } catch (error) {
      Toast.danger("获取消息失败", getErrorMessage(error));
    } finally {
      setIsLoadingMore(false);
      setIsUploading(false);
    }
  };

  // 检查是否在底部
  const checkIfAtBottom = useCallback(() => {
    if (!messageContainerRef.current) return;
    const { scrollTop, scrollHeight, clientHeight } =
      messageContainerRef.current;
    const threshold = 50; // 距离底部多少像素算作"底部"

    const atBottom = scrollHeight - (scrollTop + clientHeight) < threshold;

    setIsAtBottom(atBottom);

    // 如果用户滚动到底部，清除新消息计数
    if (atBottom) {
      setNewMessagesCount(0);
    }
  }, []);

  // 滚动到底部
  const scrollToBottom = useCallback((behavior: ScrollBehavior = "smooth") => {
    if (messageContainerRef.current?.lastElementChild) {
      messageContainerRef.current.lastElementChild.scrollIntoView({
        behavior,
      });
      setIsAtBottom(true);
      setNewMessagesCount(0);
    }
  }, []);

  // 处理滚动事件
  // 修改handleScroll函数，添加加载更多的逻辑
  const handleScroll = useCallback(() => {
    checkIfAtBottom();

    // 检查是否滚动到顶部且没有正在加载
    if (messageContainerRef.current && !isLoadingMore && hasMore) {
      const { scrollTop } = messageContainerRef.current;
      const threshold = 50; // 距离顶部多少像素触发加载

      if (scrollTop < threshold) {
        fetchMsg(true);
      }
    }
  }, [checkIfAtBottom, isLoadingMore, hasMore]);

  // 上传文件到服务端
  const uploadFile = async (
    file: File,
    isImage: boolean,
  ): Promise<attachment> => {
    setIsUploading(true);
    try {
      const formData = new FormData();

      formData.append("file", file);
      formData.append("file_name", file.name);
      formData.append("owner_type", "message");

      // 只有图片需要生成缩略图
      if (isImage) {
        formData.append("need_thumb", "true");
      }

      const endpoint = isImage ? "/api/v1/upload/image" : "/api/v1/upload/file";
      const response = await axiosInstanceWithAuth.post(endpoint, formData, {
        headers: { "Content-Type": "multipart/form-data" },
      });

      return {
        file_id: response.data.data.file_id,
        file_url: response.data.data.file_url,
        thumb_url: response.data.data.thumb_url,
        file_name: file.name,
        file_size: file.size,
        mime_type: file.type,
        width: response.data.data.width,
        height: response.data.data.height,
      };
    } catch (error) {
      Toast.danger("文件上传失败", getErrorMessage(error));
      throw error;
    } finally {
      setIsUploading(false);
    }
  };

  // 处理图片选择
  const handleImageSelect = async (e: React.ChangeEvent<HTMLInputElement>) => {
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
    if (imageInputRef.current) imageInputRef.current.value = "";
  };

  // 处理文件选择（非图片）
  const handleFileSelect = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;

    if (!files || files.length === 0) return;

    try {
      // 过滤掉图片文件（图片走单独的流程）
      const nonImageFiles = Array.from(files).filter(
        (file) => !file.type.startsWith("image/"),
      );

      if (nonImageFiles.length === 0) {
        Toast.warning("请选择非图片文件", "图片请使用图片上传按钮");

        return;
      }

      // 上传所有文件
      const uploadPromises = nonImageFiles.map((file) =>
        uploadFile(file, false),
      );
      const attachments = await Promise.all(uploadPromises);

      // 发送纯文件消息
      await sendMessage({
        msg_type: 4, // 文件消息
        content: "",
        attachments: attachments,
      });
    } catch (error) {
      console.error("文件上传失败:", error);
    } finally {
      if (fileInputRef.current) fileInputRef.current.value = "";
    }
  };

  // 移除本地图片预览
  const removeLocalImage = (index: number) => {
    setLocalImages((prev) => {
      const newImages = [...prev];

      URL.revokeObjectURL(newImages[index].preview); // 释放内存
      newImages.splice(index, 1);

      return newImages;
    });
  };

  // 发送消息（处理所有类型）
  const sendMessage = async (messageData?: {
    msg_type: number;
    content: string;
    attachments?: attachment[];
  }) => {
    // 验证是否有内容可发送
    if (!messageData && inputValue.trim() === "" && localImages.length === 0) {
      Toast.warning("发送失败", "消息内容不能为空");

      return;
    }

    const data = messageData || {
      msg_type: localImages.length > 0 ? 3 : 1, // 复合消息或纯文本
      content: inputValue,
      attachments: [],
    };

    try {
      setIsUploading(true);

      // 如果有本地图片需要上传
      if (localImages.length > 0 && !messageData) {
        const uploadPromises = localImages.map((img) =>
          uploadFile(img.file, true),
        );

        data.attachments = await Promise.all(uploadPromises);
      }

      // 发送消息
      await axiosInstanceWithAuth.post(`/api/v1/chat/${aid}/${eid}/send`, data);

      // 清空输入状态（如果不是外部调用的消息）
      if (!messageData) {
        setInputValue("");
        setLocalImages([]);
      }

      // 发送消息后总是滚动到底部
      setTimeout(() => scrollToBottom(), 100);
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
    }
  };

  // WebSocket消息处理
  const handleWebSocketMessage = useCallback(
    (message: Message) => {
      if (
        message.activity.id === aidRef.current &&
        (message.exhibitor?.id === eidRef.current || message.exhibitor === null)
      ) {
        setMessages((prev) => [...prev, message]);

        // 如果用户在底部，自动滚动到底部
        if (isAtBottom) {
          setTimeout(() => scrollToBottom(), 100);
        } else {
          // 否则增加新消息计数
          setNewMessagesCount((prev) => prev + 1);
        }
      }
    },
    [isAtBottom, scrollToBottom],
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
                src={att.thumb_url || att.file_url}
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
            {message.content && (
              <p className="whitespace-pre-wrap">{message.content}</p>
            )}
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
          </div>
        );

      case 4: // 文件消息
        return (
          <Card isBlurred className="space-y-2">
            {message.attachments.map((att, idx) => (
              <div
                key={idx}
                className="flex items-center p-2 border rounded-lg"
              >
                <FileIcon className="" size={32} />
                <div className="flex-1 min-w-0 px-2">
                  <p className="truncate font-medium">{att.file_name}</p>
                  <p className="text-xs text-gray-500">
                    {(att.file_size / 1024).toFixed(1)} KB
                  </p>
                </div>
                <Button
                  size="sm"
                  variant="solid"
                  onPress={() => window.open(att.file_url, "_blank")}
                >
                  下载
                </Button>
              </div>
            ))}
          </Card>
        );

      default:
        return <p className="text-gray-500">[不支持的消息类型]</p>;
    }
  };

  const handleJoinActivity = async () => {
    try {
      await axiosInstanceWithAuth.post(`/api/v1/activity/${aid}/join`);
      Toast.success("加入展会成功", null);
      setTimeout(() => {
        window.location.reload();
      }, 1000);
    } catch (error) {
      Toast.warning("加入展会失败", getErrorMessage(error));
    }
  };

  // 复制消息内容到剪贴板
  const copyMessage = (message: Message) => {
    let textToCopy = message.content;

    // 如果是复合消息，添加附件信息
    if (message.msg_type === 3 && message.attachments.length > 0) {
      textToCopy += message.attachments.map((a) => a.file_url).join("\n");
    }

    // 如果是图片或文件消息
    if (message.msg_type === 2 || message.msg_type === 4) {
      textToCopy = message.attachments.map((a) => a.file_url).join("\n");
    }

    navigator.clipboard
      .writeText(textToCopy)
      .then(() => Toast.success("已复制到剪贴板", null))
      .catch(() => Toast.warning("复制失败", "请检查浏览器权限"));
  };

  const [isWithdrawModalOpen, setIsWithdrawModalOpen] = useState(false);
  const [messageToWithdraw, setMessageToWithdraw] = useState<string | null>(
    null,
  );
  // 打开撤回确认对话框
  const openWithdrawConfirm = (messageId: string) => {
    setMessageToWithdraw(messageId);
    setIsWithdrawModalOpen(true);
  };

  // 确认撤回
  const confirmWithdraw = async () => {
    if (!messageToWithdraw) return;

    try {
      const response = await axiosInstanceWithAuth.delete(
        `/api/v1/chat/msg/${messageToWithdraw}/withdraw`,
      );

      if (response.data.code === 0) {
        setMessages((prev) =>
          prev.map((msg) =>
            msg.id === messageToWithdraw ? { ...msg, status: 2 } : msg,
          ),
        );
        Toast.success("消息撤回成功", null);
      } else if (response.data.code === 11) {
        Toast.warning("撤回失败", response.data.msg);
      }
    } catch (error) {
      Toast.danger("撤回消息失败", getErrorMessage(error));
    } finally {
      setIsWithdrawModalOpen(false);
      setMessageToWithdraw(null);
    }
  };

  // 取消撤回
  const cancelWithdraw = () => {
    setIsWithdrawModalOpen(false);
    setMessageToWithdraw(null);
  };

  return (
    <>
      {/* 聊天室标题 */}
      <div className="sticky top-0 z-10 flex justify-between items-center p-4 px-10 bg-white border-b border-gray-200">
        <div className="text-xl font-bold">
          {exhibitorName ? exhibitorName : "展会大厅"}
        </div>
      </div>
      <Divider className="" />

      {/* 聊天消息区域 */}
      <Card className="flex flex-col flex-grow overflow-y-auto" radius="none">
        <ScrollShadow
          ref={messageContainerRef}
          hideScrollBar
          className="space-y-6 p-4 px-10"
          onScroll={handleScroll}
        >
          {isLoadingMore && (
            <div className="flex justify-center py-2">
              <Spinner size="sm" />
            </div>
          )}
          {messages.map((message) => {
            // 已撤回消息的特殊渲染
            if (message.status === 2) {
              return (
                <div key={message.id} className="flex justify-center py-2">
                  <div className="text-gray-400 text-sm">
                    {message.from_user.id === currentUser.id
                      ? "你撤回了一条消息"
                      : `${message.from_user.nickname}撤回了一条消息`}
                  </div>
                </div>
              );
            }

            // 正常消息的渲染
            return (
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
                  className={`border-none w-fit relative group ${message.from_user.id === currentUser.id ? "bg-blue-100" : ""}`}
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
                    <div className="max-w-full px-2 py-2">
                      <div>{renderMessageContent(message)}</div>
                      <div className="py-1">
                        <div
                          className={`absolute -bottom-1 right-0 opacity-0 group-hover:opacity-100 transition-opacity duration-200 flex space-x-1 rounded-lg shadow-sm p-1`}
                        >
                          <Tooltip content="复制" showArrow={true}>
                            <Button
                              isIconOnly
                              className="w-6 h-6 min-w-6"
                              size="sm"
                              variant="light"
                              onPress={() => copyMessage(message)}
                            >
                              <CopyIcon size={14} />
                            </Button>
                          </Tooltip>
                        </div>

                        {message.from_user.id === currentUser.id && (
                          <div
                            className={`absolute -bottom-1 right-6 opacity-0 group-hover:opacity-100 transition-opacity duration-200 flex space-x-1 rounded-lg shadow-sm p-1`}
                          >
                            <Tooltip content="撤回" showArrow={true}>
                              <Button
                                isIconOnly
                                className="w-6 h-6 min-w-6"
                                size="sm"
                                variant="light"
                                onPress={() => openWithdrawConfirm(message.id)}
                              >
                                <WithDrawIcon size={14} />
                              </Button>
                            </Tooltip>
                          </div>
                        )}
                      </div>
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
            );
          })}
        </ScrollShadow>
      </Card>

      {/* 输入区域 */}
      <div className="sticky bottom-0 z-10 w-full bg-white flex flex-col">
        {/* 图片预览区域 */}
        {localImages.length > 0 && (
          <div className="flex p-2 space-x-2 overflow-x-auto">
            {localImages.map((img, index) => (
              <div key={index} className="relative">
                <Badge
                  color="default"
                  content={
                    <Button
                      isIconOnly
                      size="sm"
                      variant="light"
                      onPress={() => removeLocalImage(index)}
                    >
                      <CloseIcon size={14} />
                    </Button>
                  }
                >
                  <Image
                    alt="图片预览"
                    className="w-20 h-20 object-cover rounded-lg"
                    src={img.preview}
                  />
                </Badge>
              </div>
            ))}
          </div>
        )}

        {/* 文本输入区域 */}
        <div className="flex relative">
          {/* 新消息提示标签 - 现在放在输入框的右上方 */}
          {newMessagesCount > 0 && !isAtBottom && (
            <div className="absolute -top-8 right-24 z-20">
              <Chip
                color="primary"
                endContent={<ChevronDownIcon size={14} />}
                size="sm"
                onClick={() => scrollToBottom()}
              >
                {newMessagesCount}条新消息
              </Chip>
            </div>
          )}
          <Textarea
            className="flex-grow"
            endContent={
              <div className="flex items-center space-x-1">
                <Tooltip content="上传图片" showArrow={true}>
                  <Button
                    isIconOnly
                    isLoading={isUploading}
                    size="sm"
                    variant="light"
                    onPress={() => imageInputRef.current?.click()}
                  >
                    {isUploading ? (
                      <Spinner size="sm" />
                    ) : (
                      <ImageIcon size={20} />
                    )}
                  </Button>
                </Tooltip>
                <Input
                  ref={imageInputRef}
                  multiple
                  accept="image/*"
                  className="hidden"
                  type="file"
                  onChange={handleImageSelect}
                />

                <Tooltip content="发送文件" showArrow={true}>
                  <Button
                    isIconOnly
                    isLoading={isUploading}
                    size="sm"
                    variant="light"
                    onPress={() => fileInputRef.current?.click()}
                  >
                    {isUploading ? (
                      <Spinner size="sm" />
                    ) : (
                      <AttachmentIcon size={20} />
                    )}
                  </Button>
                </Tooltip>
                <Input
                  ref={fileInputRef}
                  multiple
                  accept=".pdf,.zip,.rar,.7z,.doc,.docx,.xls,.xlsx,.ppt,.pptx,.txt"
                  className="hidden"
                  type="file"
                  onChange={handleFileSelect}
                />
              </div>
            }
            isDisabled={!isJoined}
            maxRows={4}
            minRows={1}
            placeholder={
              isJoined
                ? "输入消息..."
                : "你还不能在这里聊天，必须先加入展会才能开始互动"
            }
            size="lg"
            type="text"
            value={inputValue}
            onKeyDown={handleHotKeysPress}
            onValueChange={setInputValue}
          />
          {isJoined ? (
            <Button
              className="ml-2"
              color="primary"
              isDisabled={inputValue.trim() === "" && localImages.length === 0}
              isLoading={isUploading}
              size="lg"
              onPress={() => sendMessage()}
            >
              发送
            </Button>
          ) : (
            <Button
              className="ml-2"
              color="success"
              size="lg"
              onPress={() => handleJoinActivity()}
            >
              加入
            </Button>
          )}
        </div>
      </div>
      <Modal isOpen={isWithdrawModalOpen} onClose={cancelWithdraw}>
        <ModalContent>
          <ModalHeader>确认撤回</ModalHeader>
          <ModalBody>
            确定要撤回这条消息吗？撤回后所有用户将看不到这条消息内容。
          </ModalBody>
          <ModalFooter>
            <Button variant="light" onPress={cancelWithdraw}>
              取消
            </Button>
            <Button color="danger" onPress={confirmWithdraw}>
              撤回
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default Chat;
