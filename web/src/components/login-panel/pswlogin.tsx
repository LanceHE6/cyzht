import {
  Card,
  Form,
  CardBody,
  Input,
  Button,
  CircularProgress,
} from "@nextui-org/react";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";

import { setToken, setUser } from "@/utils/localStorage";
import { axiosInstance } from "@/utils/axiosInstance.ts";

export default function PswLoginPanel() {
  const navigate = useNavigate();

  const [loginBtn, setLoginBtn] = useState<React.ReactNode>("登录");
  const [account, setAccount] = useState("");
  const [password, setPassword] = useState("");

  const [formErrors] = useState({});

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
        setTimeout(() => {
          navigate("/");
        }, 1000);
      } else {
        toast.error(`登录失败: ${response.data.msg}`);
      }
      setLoginBtn(() => "登录");
    }, 1000);
  };

  return (
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
            color="primary"
            size="lg"
            type={"submit"}
          >
            {loginBtn}
          </Button>
        </Form>
      </CardBody>
    </Card>
  );
}
