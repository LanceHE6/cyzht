import {
  Button,
  Card,
  CardBody,
  CircularProgress,
  Image,
  Input,
  InputOtp,
  Tab,
  Tabs,
  Form,
} from "@nextui-org/react";
import React, { useRef, useState } from "react";
import { ToastContainer, toast } from "react-toastify";
import { useNavigate } from "react-router-dom";

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

  const [formErrors, setformErrors] = useState({});

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
        navigate("/");
      } else {
        toast.error(`登录失败: ${response.data.msg}`);
      }

      setLoginOrRegisterBtn(() => "登录/注册");
    }, 500);
  };

  // 密码登录请求
  const pswLogin = async (event: { preventDefault: () => void }) => {
    event.preventDefault();
    setLoginBtn(() => (
      <CircularProgress aria-label="Loading..." color="default" />
    ));
    setTimeout(async () => {
      const data = {
        account: account,
        password: password,
      };

      const response = await axiosInstance.post("/api/v1/user/login", data);

      if (response.status !== 200) {
        toast.error(`请求出错: ${response.data.msg}`);
      } else if (response.data.code === 0 || response.data.code === 1) {
        toast.success("登录成功");
        setToken(response.data.data.token);
        setUser(response.data.data.user);
        navigate("/");
      } else {
        toast.error(`登录失败: ${response.data.msg}`);
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
      <div className="container flex items-center justify-center h-screen bg-white ">
        <Card isBlurred className="card w-3/4 h-4/6 min-w-96">
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
                <Tab key="vc" className={"w-10/12"} title="验证码登录">
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
                            description="未注册的邮箱将会直接注册账号"
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
                            color="secondary"
                            size="lg"
                            type={"submit"}
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
                            className={"px-3 w-1/2"}
                            color="secondary"
                            size="sm"
                            onClick={() => {
                              setIsCodeInputVisible(false);
                            }}
                          >
                            返回
                          </Button>
                          <Button
                            className="px-3 w-1/2"
                            color="secondary"
                            size="lg"
                            type={"submit"}
                          >
                            登录/注册
                          </Button>
                        </Form>
                      )}
                    </CardBody>
                  </Card>
                </Tab>

                <Tab key="psw" className={"w-5/6"} title="密码登录">
                  <Card className="h-80">
                    <CardBody className="flex items-center justify-center">
                      <Form
                        className="w-full justify-center items-center space-y-4"
                        validationBehavior="native"
                        validationErrors={formErrors}
                        onSubmit={pswLogin}
                      >
                        <Input
                          isClearable
                          isRequired
                          className="px-3"
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
                        <br />
                        <Input
                          isRequired
                          className="px-3"
                          errorMessage={({ validationDetails }) => {
                            if (validationDetails.valueMissing) {
                              return "请输入密码";
                            }
                            if (validationDetails.tooShort) {
                              return "密码长度必须大于8";
                            }
                          }}
                          label="密码"
                          labelPlacement="outside"
                          size="lg"
                          type={"password"}
                          value={password}
                          onValueChange={setPassword}
                        />
                        <br />
                        <Button
                          className="px-3 w-1/2"
                          color="secondary"
                          size="lg"
                          type={"submit"}
                        >
                          {loginBtn}
                        </Button>
                      </Form>
                    </CardBody>
                  </Card>
                </Tab>
              </Tabs>
            </div>
          </CardBody>
        </Card>
      </div>
      <ToastContainer />
    </DefaultLayout>
  );
}
