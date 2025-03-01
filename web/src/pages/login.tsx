import { Card, CardBody, Image, Tab, Tabs } from "@nextui-org/react";

import cover1 from "../assets/imgs/cover1.jpg";

import DefaultLayout from "@/layouts/default.tsx";
import VCLoginPanel from "@/components/login-panel/vclogin.tsx";
import PswLoginPanel from "@/components/login-panel/pswlogin";

export default function Login() {
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
              <text className="text-6xl py-4">重邮展会通</text>
              <br />
              <Tabs aria-label="Options">
                <Tab key="vc" className={"w-10/12"} title="验证码登录">
                  <VCLoginPanel />
                </Tab>

                <Tab key="psw" className={"w-5/6"} title="密码登录">
                  <PswLoginPanel />
                </Tab>
              </Tabs>
            </div>
          </CardBody>
        </Card>
      </div>
    </DefaultLayout>
  );
}
