package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/data/model"
	"server/pkg/jwt"
	"server/pkg/random"
	"server/pkg/response"
	"server/pkg/rpc/file_server/api/v1/file_server"
	"server/pkg/smtp"
)

// RegisterAndLoginSendCode
//
//	@Description: 注册和登录发送验证码
//	@receiver s userService 服务
//	@param context gin.Context
func (s userService) RegisterAndLoginSendCode() gin.HandlerFunc {
	return func(context *gin.Context) {
		type registerRequest struct {
			Account string `json:"account" form:"account" binding:"required"`
		}
		var data registerRequest
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
			return
		}
		u := s.UserRepo.SelectByAccount(data.Account)
		// 账号已存在, 登录
		if u != nil {
			// 生成验证码
			code := random.CreateRandomStr(6, random.Number)
			fmt.Println(code)
			// 发送验证码
			err := smtp.SMTPService.SendVerifyCodeEmail(data.Account, u.Nickname, code, smtp.LoginEmail)
			if err != nil {
				context.JSON(http.StatusInternalServerError, response.ErrorResponse(10, "发送验证码失败", err))
				return
			}
			// 验证码写入redis 5分钟
			s.VerifyCodeRepo.SetVerifyCode(data.Account, code)
			context.JSON(http.StatusOK, response.SuccessResponse(nil))
		} else {
			// 注册
			// 生成验证码
			code := random.CreateRandomStr(6, random.Number)
			// 发送验证码
			err := smtp.SMTPService.SendVerifyCodeEmail(data.Account, "亲爱的用户", code, smtp.RegisterEmail)
			if err != nil {
				context.JSON(http.StatusInternalServerError, response.ErrorResponse(11, "发送验证码失败", err))
				return
			}
			// 验证码写入redis 5分钟
			s.VerifyCodeRepo.SetVerifyCode(data.Account, code)
			context.JSON(http.StatusOK, response.SuccessResponse(nil))
		}
	}
}

// RegisterAndLoginVerifyCode
//
//	@Description: 验证注册和登录验证码
//	@receiver s userService
//	@return gin.HandlerFunc
func (s userService) RegisterAndLoginVerifyCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type registerRequest struct {
			Account string `json:"account" form:"account" binding:"required"`
			Code    string `json:"code" form:"code" binding:"required"`
		}
		var data registerRequest
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
			return
		}
		// 获取验证码
		code := s.VerifyCodeRepo.GetVerifyCode(data.Account)
		if code != data.Code {
			ctx.JSON(http.StatusOK, response.FailedResponse(10, "验证码错误"))
			return
		}
		u := s.UserRepo.SelectByAccount(data.Account)
		if u == nil {
			// 账号不存在,注册

			// 用临时密码,并发送邮件给用户
			psw := random.CreateRandomStr(8, random.NumberAndLetter)
			name := "user_" + random.CreateRandomStr(6, random.NumberAndLetter)
			newUser := model.UserModel{
				Account:  data.Account,
				Password: psw,
				Nickname: name,
			}
			// 发送邮件
			err := smtp.SMTPService.SendTemporaryPswEmail(data.Account, psw)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-3, "发送临时密码失败", err))
				return
			}
			// 创建用户
			err = s.UserRepo.Insert(newUser)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-4, "创建用户失败", err))
				return
			}
		}
		u = s.UserRepo.SelectByAccount(data.Account)
		// 登录
		sessionID := random.CreateRandomStr(16, random.NumberAndLetter)
		u.SessionID = sessionID
		token, err := jwt.GenerateToken(u.ID, sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "生成token失败", err))
			return
		}
		// 更新session_id
		err = s.UserRepo.UpdateSessionID(u.ID, sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-2, "更新session_id失败", err))
		}

		// 移除验证码
		s.VerifyCodeRepo.DeleteVerifyCode(data.Account)

		var avatar string
		rep, err := s.FileRpcServer.GetAvatarUrl(context.Background(), &file_server.GetAvatarUrlRequest{Id: u.ID}) // 获取头像
		if err != nil {
			fmt.Println("failed to get avatar url:", err.Error())
		} else {
			avatar = rep.FileUrl
			u.Avatar = avatar
		}

		ctx.JSON(http.StatusOK, response.SuccessResponse(map[string]any{
			"token": token,
			"user":  u,
		}))
	}
}

// Login
//
//	@Description: 登录
//	@receiver s userService 服务
//	@param account 账号
//	@param password 密码
//	@return *pkg.Response 返回结果
func (s userService) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// loginRequest
		//
		//	@Description: 登录请求参数结构体
		type loginRequest struct {
			Account  string `json:"account" form:"account" binding:"required"`
			Password string `json:"password" form:"password" binding:"required"`
		}
		var data loginRequest
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
			return
		}

		u := s.UserRepo.SelectByAccountAndPsw(data.Account, data.Password)
		if u == nil {
			ctx.JSON(http.StatusOK, response.FailedResponse(10, "账号或密码错误"))
			return
		}
		sessionID := random.CreateRandomStr(16, random.NumberAndLetter)
		u.SessionID = sessionID
		token, err := jwt.GenerateToken(u.ID, sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, "生成token失败"))
			return
		}
		// 更新session_id
		err = s.UserRepo.UpdateSessionID(u.ID, sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-2, "更新session_id失败"))
			return
		}
		var avatar string
		rep, err := s.FileRpcServer.GetAvatarUrl(context.Background(), &file_server.GetAvatarUrlRequest{Id: u.ID}) // 获取头像
		if err != nil {
			fmt.Println("failed to get avatar url:", err.Error())
		} else {
			avatar = rep.FileUrl
			u.Avatar = avatar
		}
		ctx.JSON(http.StatusOK, response.SuccessResponse(map[string]any{
			"token": token,
			"user":  u,
		}))
		return
	}
}
