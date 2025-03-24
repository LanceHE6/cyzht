// chat.tsx
import { Avatar, Card, Divider, ScrollShadow } from "@heroui/react";
import React, { useState, useEffect, useRef } from "react";
import { Textarea } from "@heroui/input";
import { useNavigate } from "react-router-dom";

import {
  axiosInstanceWithAuth,
  setNavigateCallback,
} from "@/utils/axios-instance.ts";
import { LocalStorage, Toast, WebSocketClient } from "@/utils/utils.ts";

interface user {
  id: string;
  nickname: string;
  avatar: string;
}

interface Message {
  id: string;
  created_at: string;
  updated_at: string;
  activity: {};
  exhibitor: {};
  from_user: user;
  to_user: user | null;
  msg_type: number;
  text_msg: string;
  file_url: string;
  file_size: number;
}

export interface ChatProps {
  aid: string;
}

// 聊天页面子组件
const Chat: React.FC<ChatProps> = (props: ChatProps) => {
  // 为请求设置导航回调
  const navigate = useNavigate();

  useEffect(() => {
    setNavigateCallback(navigate);
  }, [navigate]);

  const [messages, setMessages] = useState<Message[]>([]);
  const [inputValue, setInputValue] = useState<string>("");

  const [currentUser] = useState<any>(LocalStorage.getUser());

  if (!currentUser) {
    Toast.danger("请先登录", "请先登录");
    navigate("/login");
  }

  let { aid } = props;

  const fetchMsg = async () => {
    try {
      const response = await axiosInstanceWithAuth.get(
        `/api/v1/activity/${aid}/msg`,
      );

      if (response.status !== 200) {
        throw new Error(response.data.data.msg);
      }
      // 解析和映射数据
      const parsedMessages: Message[] = response.data.data.rows.map(
        (msg: any) => ({
          id: msg.id,
          created_at: msg.created_at,
          updated_at: msg.updated_at,
          activity: msg.activity,
          exhibitor: msg.exhibitor,
          from_user: {
            id: msg.from_user.id,
            nickname: msg.from_user.nickname,
            avatar: msg.from_user.avatar,
          },
          to_user:
            msg.to_user === null
              ? null
              : {
                  id: msg.to_user.id,
                  nickname: msg.to_user.nickname,
                  avatar: msg.to_user.avatar,
                },
          msg_type: msg.msg_type,
          text_msg: msg.text_msg,
          file_url: msg.file_url,
          file_size: msg.file_size,
        }),
      );

      setMessages(parsedMessages);
      console.log("parsed", parsedMessages);
    } catch (error) {
      Toast.danger("获取消息失败", error.message);
    }
  };

  useEffect(() => {
    // 模拟从后端获取消息数据
    fetchMsg();
  }, []);

  const sendMessage = () => {
    const newMessage = {
      msg_type: 1,
      text_msg: inputValue,
    };

    axiosInstanceWithAuth
      .post(`/api/v1/activity/${aid}/send`, newMessage)
      .then((response) => {
        if (response.data.code !== 0) {
          Toast.warning("发送消息失败", response.data.msg);
        }
      });

    setInputValue("");
  };

  const handleKeyPress = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      sendMessage();
    }
  };

  // 滚动到最新消息
  const messageContainerRef = useRef<HTMLDivElement>(null);

  // 监听 messages 变化，滚动到最新消息
  useEffect(() => {
    if (messageContainerRef.current) {
      const lastMessage = messageContainerRef.current.lastElementChild;

      if (lastMessage) {
        lastMessage.scrollIntoView({ behavior: "smooth" });
      }
    }
  }, [messages]);

  // WebSocket 消息处理回调
  const handleWebSocketMessage = (message: Message) => {
    console.log("WebSocket 消息处理回调:", message);
    setMessages((prevMessages) => [...prevMessages, message]);
  };

  useEffect(() => {
    // 注册 WebSocket 消息处理回调
    WebSocketClient.onMessage(handleWebSocketMessage);

    // 清理函数
    return () => {
      WebSocketClient.offMessage(handleWebSocketMessage);
    };
  }, []);

  return (
    <Card className="col-span-8 h-full bg-[#FFF6FF] flex flex-col">
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
                className={`border-none max-w-md ${message.from_user.id === currentUser.id ? "bg-blue-100" : ""}`}
              >
                <div className="message p-2">
                  <div className="flex items-center">
                    {/*适配自己消息,使名字始终为靠近头像的一侧*/}
                    {message.from_user.id !== currentUser.id ? (
                      <>
                        <div className="text-sm font-bold mr-2">
                          {message.from_user.nickname}
                        </div>
                        <div className="text-gray-500 text-xs">
                          {message.created_at}
                        </div>
                      </>
                    ) : (
                      <>
                        <div className="text-gray-500 text-xs mr-2">
                          {message.created_at}
                        </div>
                        <div className="text-sm font-bold">
                          {message.from_user.nickname}
                        </div>
                      </>
                    )}
                  </div>
                  <Textarea
                    disabled
                    minRows={1}
                    value={message.text_msg}
                    variant="flat"
                  />
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

      {/* 输入框 */}
      <div className="sticky bottom-0 z-10 w-full bg-white border-t border-gray-200">
        <Textarea
          className=""
          maxRows={4}
          minRows={1}
          placeholder="输入消息..."
          size="lg"
          type="text"
          value={inputValue}
          onKeyPress={handleKeyPress}
          onValueChange={setInputValue}
        />
      </div>
    </Card>
  );
};

export default Chat;
