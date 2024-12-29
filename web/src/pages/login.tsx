import {
  Button,
  Card,
  CardBody,
  CircularProgress,
  Image,
  Input,
  Tab,
  Tabs,
} from "@nextui-org/react";
import React, { useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useForm, UseFormReturn } from "react-hook-form";

import { axiosInstance } from "../utils/axiosInstance.ts";
import { setToken, setUser } from "../utils/localStorage";
import cover1 from "../assets/imgs/cover1.jpg";

import DefaultLayout from "@/layouts/default.tsx";

export default function Login() {
  const navigate = useNavigate();
  const [sendCodeBtn, setSendCodeBtn] = useState<React.ReactNode>("发送验证码");
  const [loginOrRegisterBtn, setLoginOrRegisterBtn] =
    useState<React.ReactNode>("登录/注册");
  const [loginBtn, setLoginBtn] = useState<React.ReactNode>("登录");
  const [account, setAccount] = useState("");
  const [code, setCode] = useState("");
  const [password, setPassword] = useState("");

  // 表单校验
  const sendCodeCheck: UseFormReturn<{ account: string }> = useForm();
  const verifyCodeCheck: UseFormReturn<{ code: string }> = useForm();
  const pswLoginCheck: UseFormReturn<{ account: string; password: string }> =
    useForm();

  // 发送验证码请求
  const sendCode = async () => {
    const data = {
      account: account,
    };

    try {
      const response = await axiosInstance.post(
        "/api/v1/user/register&login_send_code",
        data,
      );

      if (response.status !== 200) {
        alert("请求出错");
      } else if (response.data.code === 0) {
        alert("验证码发送成功");
        startCountdown();
      } else {
        alert("验证码发送失败: " + response.data.msg);
      }
    } catch (error) {
      alert(error);
    }
  };

  // 验证验证码请求
  const verifyCode = async () => {
    setLoginOrRegisterBtn(() => (
      <CircularProgress aria-label="Loading..." color="default" />
    ));
    setTimeout(async () => {
      const data = {
        account: account,
        code: code,
      };

      try {
        const response = await axiosInstance.post(
          "/api/v1/user/register&login_verify_code",
          data,
        );

        if (response.status !== 200) {
          alert("请求出错");
        } else if (response.data.code === 0 || response.data.code === 1) {
          alert("登录成功");
          setToken(response.data.data.token);
          setUser(response.data.data.user);
          navigate("/");
        } else {
          alert("验证失败: " + response.data.msg);
        }
      } catch (error) {
        alert(error);
      }
      setLoginOrRegisterBtn(() => "登录/注册");
    }, 500);
  };

  // 密码登录请求
  const pswLogin = async () => {
    setLoginBtn(() => (
      <CircularProgress aria-label="Loading..." color="default" />
    ));
    setTimeout(async () => {
      const data = {
        account: account,
        password: password,
      };

      try {
        const response = await axiosInstance.post("/api/v1/user/login", data);

        if (response.status !== 200) {
          alert("请求出错");

          return;
        }
        if (response.data.code === 0) {
          alert("登录成功");
          setToken(response.data.data.token);
          setUser(response.data.data.user);
          navigate("/");
        } else {
          alert("登录失败: " + response.data.msg);
        }
      } catch (error) {
        alert(error);
      }
      setLoginBtn(() => "登录");
    }, 500);
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
    <DefaultLayout>
      <div className="container flex items-center justify-center w-screen h-screen bg-teal-400 ">
        <Card isBlurred className="card w-1/2 h-4/6 min-w-96">
          <CardBody className="grid grid-cols-2 gap-4 overflow-hidden py-0 px-0">
            <Image
              isBlurred
              isZoomed
              alt="Card background"
              className="w-full h-auto min-h-full"
              radius={"none"}
              src={cover1}
            />
            <div className="form flex flex-col items-center justify-center w-full">
              <text className="text-6xl py-4">重邮展慧通</text>
              <br />
              <Tabs aria-label="Options">
                <Tab key="vc" className={"w-5/6"} title="验证码登录">
                  <Card className="h-80">
                    <CardBody className="flex items-center justify-center">
                      <Input
                        isClearable
                        isRequired
                        className="px-3"
                        errorMessage={
                          sendCodeCheck.formState.errors.account &&
                          "请输入合法邮箱"
                        }
                        isInvalid={
                          sendCodeCheck.formState.errors.account && true
                        }
                        label="邮箱"
                        labelPlacement="outside"
                        size="lg"
                        type="email"
                        value={account}
                        onValueChange={(newValue: string) =>
                          setAccount(newValue)
                        }
                        {...sendCodeCheck.register("account", {
                          required: true,
                          pattern:
                            /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/,
                        })}
                      />
                      <br />
                      <Input
                        isRequired
                        className="px-3"
                        description="未注册的邮箱将会直接注册账号"
                        endContent={
                          <Button
                            isDisabled={sendCodeBtn !== "发送验证码"}
                            size="sm"
                            onClick={sendCodeCheck.handleSubmit(sendCode)}
                          >
                            {sendCodeBtn}
                          </Button>
                        }
                        errorMessage={
                          verifyCodeCheck.formState.errors.code &&
                          "请输入验证码"
                        }
                        isInvalid={
                          verifyCodeCheck.formState.errors.code && true
                        }
                        label="验证码"
                        labelPlacement="outside"
                        size="lg"
                        type="text"
                        value={code}
                        onValueChange={(newValue: string) => setCode(newValue)}
                        {...verifyCodeCheck.register("code", {
                          required: true,
                        })}
                      />
                      <br />
                      <Button
                        className="px-3 w-1/2"
                        color="secondary"
                        size="lg"
                        // onClick={sendCodeCheck.handleSubmit(
                        //   verifyCodeCheck.handleSubmit(verifyCode),
                        // )}
                      >
                        {loginOrRegisterBtn}
                      </Button>
                    </CardBody>
                  </Card>
                </Tab>
                <Tab key="psw" className={"w-5/6"} title="密码登录">
                  <Card className="h-80">
                    <CardBody className="flex items-center justify-center">
                      <Input
                        isClearable
                        isRequired
                        className="px-3"
                        errorMessage={
                          pswLoginCheck.formState.errors.account &&
                          "请输入合法邮箱"
                        }
                        isInvalid={
                          pswLoginCheck.formState.errors.account && true
                        }
                        label="邮箱"
                        labelPlacement="outside"
                        size="lg"
                        type="email"
                        value={account}
                        onValueChange={(newValue: string) =>
                          setAccount(newValue)
                        }
                        {...pswLoginCheck.register("account", {
                          required: true,
                          pattern:
                            /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/,
                        })}
                      />
                      <br />
                      <Input
                        isRequired
                        className="px-3"
                        errorMessage={
                          pswLoginCheck.formState.errors.password &&
                          "请输入密码"
                        }
                        isInvalid={
                          pswLoginCheck.formState.errors.password && true
                        }
                        label="密码"
                        labelPlacement="outside"
                        size="lg"
                        type={"password"}
                        value={password}
                        onValueChange={(newValue: string) =>
                          setPassword(newValue)
                        }
                        {...pswLoginCheck.register("password", {
                          required: true,
                        })}
                      />
                      <br />
                      <Button
                        className="px-3 w-1/2"
                        color="secondary"
                        size="lg"
                        onClick={pswLoginCheck.handleSubmit(pswLogin)}
                      >
                        {loginBtn}
                      </Button>
                    </CardBody>
                  </Card>
                </Tab>
              </Tabs>
            </div>
          </CardBody>
        </Card>
      </div>
    </DefaultLayout>
  );
}
