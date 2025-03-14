package email_template

import (
	"strings"
)

func GetVerifyEmailHTML(email string, verifyCode string) string {
	template := `
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<div class="content-div">
    <div style="border-style: solid; border-width: thin; border-color:#dadce0; border-radius: 8px; padding: 40px 20px;" align="center" class="mdv2rw">
        <img src="https://jsd.onmicrosoft.cn/gh/LanceHE6/PicGo@main/imgs/logo2.png" width="74" height="74" aria-hidden="true" style="margin-bottom: 16px;">
        <div style="">
            <div style="font-size: 24px;">请验证您的邮箱 </div>
        </div>
        <div style="font-family: Roboto-Regular,Helvetica,Arial,sans-serif; font-size: 14px; color: rgba(0,0,0,0.87); line-height: 20px;padding-top: 20px; text-align: left;">NetChat 收到了使用 <a style="font-weight: bold;">${email}</a> 注册账号的请求。<br><br>请使用此验证码完成注册：<br>
            <div style="text-align: center; font-size: 36px; margin-top: 20px; line-height: 44px;">${code}</div><br>此验证码将在 5 分钟后失效。<br><br>如果不是您本人操作，请忽略这封电子邮件。
            <br><br><br>系统邮件 请勿回复
        </div>
    </div>
</div>

<style>
    .content-div {
        position: relative;
        font-size: 14px;
        height: auto;
        padding: 15px 15px 10px 15px;
        z-index: 1;
        zoom: 1;
        line-height: 1.7;
        width: 550px;
        min-width: 500px;
        margin: 10px auto;
    }
</style>
`
	template = strings.Replace(template, "${email}", email, -1)
	template = strings.Replace(template, "${code}", verifyCode, -1)
	return template
}
