import {
  Card,
  Form,
  CardBody,
  Input,
  Button,
  InputOtp,
  CircularProgress,
} from "@heroui/react";
import React, { useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";

import { setToken, setUser } from "@/utils/localStorage";
import { axiosInstance } from "@/utils/axiosInstance.ts";
import { websocketClient } from "@/utils/websocket.ts";
// 验证码登陆面板
export default function VCLoginPanel() {
  const navigate = useNavigate();
  const [sendCodeBtn, setSendCodeBtn] = useState<React.ReactNode>("发送验证码");
  const [loginOrRegisterBtn, setLoginOrRegisterBtn] =
    useState<React.ReactNode>("登录/注册");
  const [account, setAccount] = useState("");
  const [code, setCode] = useState("");

  const [formErrors] = useState({});

  const [isCodeInputVisible, setIsCodeInputVisible] = useState(false);

  // 发送验证码请求
  const sendCode = async (event: { preventDefault: () => void }) => {
    event.preventDefault();
    const data = {
      account: account,
    };
    const response = await axiosInstance.post(
      "/api/v1/user/register&login_send_code",
      data,
    );

    console.log("data:", response.data);
    if (response.status !== 200) {
      toast.error(`请求出错: ${response.data.msg}`);
    } else if (response.data.code === 0) {
      toast.success("验证码发送成功");
      startCountdown();
      setIsCodeInputVisible(true);
    } else {
      toast.error(`验证码发送失败: ${response.data.msg}`);
    }
  };

  // 验证验证码请求
  const verifyCode = async (event: { preventDefault: () => void }) => {
    event.preventDefault();
    setLoginOrRegisterBtn(() => (
      <CircularProgress aria-label="Loading..." color="default" />
    ));
    setTimeout(async () => {
      const data = {
        account: account,
        code: code,
      };
      const response = await axiosInstance.post(
        "/api/v1/user/register&login_verify_code",
        data,
      );

      if (response.status !== 200) {
        toast.error(`请求出错: ${response.data.msg}`);
      } else if (response.data.code === 0 || response.data.code === 1) {
        toast.success("登录成功");
        setToken(response.data.data.token);
        setUser(response.data.data.user);
        setTimeout(() => {
          // 初始化在线心跳
          websocketClient.close();
          navigate("/");
        }, 1000);
      } else {
        toast.error(`登录失败: ${response.data.msg}`);
      }

      setLoginOrRegisterBtn(() => "登录/注册");
    }, 1000);
  };

  type TimerID = number | NodeJS.Timeout;
  const countdownRef = useRef<TimerID | null>(null);
  // 发送验证码冷却倒计时
  const startCountdown = () => {
    let count = 60;

    setSendCodeBtn(`${count}s`);
    countdownRef.current = setInterval(() => {
      count -= 1;
      if (count > 0) {
        setSendCodeBtn(`${count}s`);
      } else {
        clearInterval(countdownRef.current!);
        setSendCodeBtn("发送验证码");
      }
    }, 1000);
  };

  return (
    <Card className="h-80">
      <CardBody className="flex items-center justify-center">
        {!isCodeInputVisible ? (
          <Form
            className="w-full justify-center items-center space-y-4"
            validationBehavior="native"
            validationErrors={formErrors}
            onSubmit={sendCode}
          >
            <Input
              isClearable
              isRequired
              className="px-3"
              description="首次登录将自动注册账号"
              errorMessage={({ validationDetails }) => {
                if (validationDetails.valueMissing) {
                  return "请输入邮箱";
                }
                if (validationDetails.typeMismatch) {
                  return "请输入合法的邮箱";
                }
              }}
              label="邮箱"
              labelPlacement="outside"
              size="lg"
              type="email"
              value={account}
              onValueChange={setAccount}
            />
            <Button
              className="px-3 w-1/2"
              color="primary"
              size="lg"
              type={"submit"}
              isDisabled={!(sendCodeBtn === "发送验证码")}
            >
              {sendCodeBtn}
            </Button>
          </Form>
        ) : (
          <Form
            className="w-full justify-center items-center space-y-4"
            validationBehavior="native"
            validationErrors={formErrors}
            onSubmit={verifyCode}
          >
            <div className="text-default-500">验证码</div>
            <InputOtp
              isRequired
              className={"px-3 w-10/12"}
              errorMessage={({ validationDetails }) => {
                if (validationDetails.valueMissing) {
                  return "请输入验证码";
                }
              }}
              length={6}
              size="md"
              type="number"
              value={code}
              onValueChange={setCode}
            />
            <Button
              className={"px-3 w-1/4"}
              color="secondary"
              size="sm"
              variant="light"
              onClick={() => {
                setIsCodeInputVisible(false);
              }}
            >
              返回
            </Button>
            <Button
              className="px-3 w-1/2"
              color="primary"
              size="lg"
              type={"submit"}
            >
              {loginOrRegisterBtn}
            </Button>
          </Form>
        )}
      </CardBody>
    </Card>
  );
}
